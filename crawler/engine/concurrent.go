package engine

import (
	"crypto/md5"
	"io"
	"log"
)

type ConcurrentEngine struct {
	Scheduler        Scheduler
	WorkerCount      int
	ItemChan         chan Item
	RequestProcessor Processor
}

type Processor func(Request) (ParseResult, error)

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		e.createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		if isDuplicate(r.Url) {
			//log.Printf("Duplicate request: %s", r.Url)
			continue
		}
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			go func(i Item) {
				e.ItemChan <- i
			}(item)
		}

		for _, r := range result.Requests {
			if isDuplicate(r.Url) {
				//log.Printf("Duplicate request: %s", r.Url)
				continue
			}
			e.Scheduler.Submit(r)
		}
	}
}

var visitedUrls = make(map[string]bool)

func isDuplicate(url string) bool {
	h := md5.New()
	if _, err := io.WriteString(h, url); err != nil {
		log.Printf("isDuplicate md5 err: %v", err)
		return false
	}
	s := h.Sum(nil)

	newUrl := string(s)
	if visitedUrls[newUrl] {
		return true
	}
	visitedUrls[newUrl] = true
	return false
}

func (e *ConcurrentEngine) createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			ready.WorkerReady(in)
			r := <-in
			parserResult, err := e.RequestProcessor(r)
			if err != nil {
				continue
			}
			out <- parserResult
		}
	}()
}
