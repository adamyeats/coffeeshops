package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/adamyeats/coffeeshops/graphql-server/internal/o11y"
	"github.com/adamyeats/coffeeshops/graphql-server/internal/server"
	"github.com/adamyeats/coffeeshops/graphql-server/pkg/config"
)

// main is the entrypoint for the server
// this is a bit of a mess right now
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cleanup := o11y.NewTracer()
	defer cleanup(ctx)

	srv, err := server.New(ctx, config.FromEnv())
	if err != nil {
		log.Fatal("Failed to create server:", err)
	}

	go func() {
		if err := srv.Serve(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Failed to serve:", err)
		}
	}()

	// wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Server is shutting down...")

	ctx, cancel = context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := srv.Stop(ctx); err != nil {
		log.Fatal("Failed to shutdown server:", err)
	}

	log.Println("Server exited gracefully")
}
