package main

import (
	"GoBasic/lang/wire"
	"fmt"
	"os"
)

func main() {
	//e := wire.InitializeEvent()
	//e.Start()

	e, err := wire.InitializeEvent("hello wire")
	if err != nil {
		fmt.Printf("failed to create event: %s\n", err)
		os.Exit(2)
	}
	e.Start()
}
