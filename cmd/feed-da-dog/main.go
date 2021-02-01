package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/omerkaya1/feed-da-dog/internal"
)

// 12-factor-app
func main() {
	// TODO: Parse env variables

	// Initialise a logger
	logger := log.New(os.Stdout, "feed-da-dog", log.Ldate|log.Ltime|log.Llongfile|log.Lmsgprefix)
	// Base context
	ctx, cancel := context.WithCancel(context.Background())

	// Handle interrupts
	quit := make(chan os.Signal, 1)
	signal.Notify(quit)

	// Monitor errors
	errChan := make(chan error)
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case err := <-errChan:
				if err != nil {
					logger.Println(err)
				}
			}
		}
	}()

	// Start the web server
	go func() {
		if err := internal.NewServer(internal.NewDummyDB(), logger, errChan).Start(ctx); err != nil {
			logger.Fatal(err)
		}
	}()

	<-quit
	cancel()
	close(errChan)
}
