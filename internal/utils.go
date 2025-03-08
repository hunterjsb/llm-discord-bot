package internal

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// SetupCloseHandler creates a handler that will catch SIGINT and SIGTERM signals
// and gracefully close the application
func SetupCloseHandler(cleanupFunc func() error) {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\nShutting down...")
		err := cleanupFunc()
		if err != nil {
			fmt.Printf("Error during cleanup: %v\n", err)
			os.Exit(1)
		}
		os.Exit(0)
	}()
}