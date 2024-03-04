package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

func init() {
	config.LoadEnv()
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// initiate server starting
	pool := database.New(ctx)

	// check if the database is reachable
	if err := pool.Ping(ctx); err != nil {
		log.Fatalln("Error Connecting to DB: ", err)
	}
	defer pool.Close() // close the pool when the server exits

	port := config.Env.PORT

	// HTTP server initialization
	server := &http.Server{
		Addr:         "0.0.0.0:" + port,
		Handler:      server.NewHandler(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// Server run context
	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	// Listen for syscall signals for process to interrupt/quit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig // Block until a signal is received.

		// Shutdown signal with grace period of 20 seconds
		shutdownCtx, cancelFunc := context.WithTimeout(serverCtx, 20*time.Second)
		defer cancelFunc()
		go func() {
			<-shutdownCtx.Done() // Block until context is canceled
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatal("graceful shutdown timed out.. forcing exit.")
			}
		}()

		// Trigger graceful shutdown
		if err := server.Shutdown(shutdownCtx); err != nil {
			log.Fatal(err)
		}
		serverStopCtx()
	}()

	log.Printf("Server running on port %s....", port)
	// Run our server and handle any errors
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}

	// Wait for server context to be stopped
	<-serverCtx.Done()
}
