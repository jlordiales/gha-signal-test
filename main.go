package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	log.Println("Starting...")
	for {
		select {
		case s := <-c:
			log.Printf("Caught signal %s\n", s.String())
			cancel()
		case <-ctx.Done():
			log.Fatal(ctx.Err())
		}
	}
}



