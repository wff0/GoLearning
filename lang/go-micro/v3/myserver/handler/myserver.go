package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	myserver "myserver/proto"
)

type Myserver struct{}

func (e *Myserver) SayHello(ctx context.Context, req *myserver.SayRequest, res *myserver.SayResponse) error {
	res.Answer = "你也好" + req.Message
	return nil
}

// Return a new handler
func New() *Myserver {
	return &Myserver{}
}

// Call is a single request handler called via client.Call or the generated client code
func (e *Myserver) Call(ctx context.Context, req *myserver.Request, rsp *myserver.Response) error {
	log.Info("Received Myserver.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Myserver) Stream(ctx context.Context, req *myserver.StreamingRequest, stream myserver.Myserver_StreamStream) error {
	log.Infof("Received Myserver.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&myserver.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Myserver) PingPong(ctx context.Context, stream myserver.Myserver_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&myserver.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
