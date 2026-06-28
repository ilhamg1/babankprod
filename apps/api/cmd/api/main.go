package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func setupRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "OK")
	})
	return mux
}

func main() {
	// Configure structured JSON logging
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	port := ":8080"
	server := &http.Server{
		Addr:    port,
		Handler: setupRouter(),
	}

	// Channel to listen for errors coming from the listener
	serverErrors := make(chan error, 1)

	// Start the server in a goroutine
	go func() {
		slog.Info("Starting core API server", "port", port)
		serverErrors <- server.ListenAndServe()
	}()

	// Channel to listen for an interrupt or terminate signal from the OS
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	// Block until a signal is received or an error occurs
	select {
	case err := <-serverErrors:
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("Server error", "error", err)
			os.Exit(1)
		}
	case sig := <-shutdown:
		slog.Info("Starting graceful shutdown", "signal", sig)
		
		// Create a context with a timeout for the shutdown process
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		
		// Attempt graceful shutdown
		if err := server.Shutdown(ctx); err != nil {
			slog.Error("Graceful shutdown failed", "error", err)
			
			// Force shutdown
			if err := server.Close(); err != nil {
				slog.Error("Force shutdown failed", "error", err)
			}
			os.Exit(1)
		}
		slog.Info("Server stopped gracefully")
	}
}
