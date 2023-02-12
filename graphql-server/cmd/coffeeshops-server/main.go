package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/adamyeats/coffeeshops/graphql-server/internal/server"
	"github.com/adamyeats/coffeeshops/graphql-server/pkg/config"
)

// main is the entrypoint for the server
// this is a bit of a mess right now
func main() {
	ctx := context.Background()

	srv, err := server.New(ctx, config.FromEnv())

	if err != nil {
		panic(err)
	}

	// LFG
	if err := srv.Serve(); err != nil {
		log.Fatal("There was an error whilst serving:", err)
	}

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Server is shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Stop(ctx); err != nil {
		log.Fatal("There was an error whilst shutting down:", err)
	}

	// catching ctx.Done() timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}

	log.Println("Server exiting")
}
