package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
	"vkexplorer/core"
)

func main() {
	core.InitCore()
	core.StartCore()

	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(path)

	interruptCh := make(chan os.Signal, 1)
	signal.Notify(interruptCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-interruptCh
		fmt.Println("\nReceived interrupt signal. Exiting...")
		core.QuitCore()
		os.Exit(0)
	}()

	for {
		time.Sleep(time.Second)
	}

}
