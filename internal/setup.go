package internal

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nikitsingh/forky/backend/internal/config"
	"github.com/nikitsingh/forky/backend/internal/db"
)

func RunSetup() error {
	addr := ":" + config.Envs.SERVER_PORT

	dbConn, err := db.Connect()
	if err != nil {
		return fmt.Errorf("error connecting to database: %w", err)
	}

	r := NewRouter(dbConn)
	r.SetupRouter()

	server := &http.Server{
		Addr:    addr,
		Handler: r.router,
	}

	// Channel to listen for errors coming from the listener.
	serverErrors := make(chan error, 1)

	// Start the server in a goroutine
	go func() {
		log.Printf("Starting API server on %s", addr)
		serverErrors <- server.ListenAndServe()
	}()

	// Channel to listen for an interrupt or terminate signal from the OS
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	// Blocking select
	select {
	case err := <-serverErrors:
		return fmt.Errorf("error starting server: %w", err)

	case <-shutdown:
		log.Println("Starting shutdown...")

		// Disable keep-alives on existing connections
		server.SetKeepAlivesEnabled(false)

		// Give outstanding requests a deadline for completion
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// Shutdown the server
		err := server.Shutdown(ctx)
		if err != nil {
			return fmt.Errorf("could not stop server gracefully: %w", err)
		}
	}

	return nil
}
