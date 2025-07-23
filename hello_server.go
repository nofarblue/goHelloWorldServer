package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
	"unicode"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")
	// Sanitize name for logging by removing control characters
	sanitizedName := sanitizeForLogging(name)
	log.Printf("Received request for %s\n", sanitizedName)
	w.Write([]byte(CreateGreeting(name)))
}

func CreateGreeting(name string) string {
	// Trim leading and trailing whitespace
	name = strings.TrimSpace(name)
	
	// Return "Hello, Guest" if input is empty after trim
	if name == "" {
		return "Hello, Guest\n"
	}
	
	// Limit name length to 100 characters
	if len(name) > 100 {
		name = name[:100]
	}
	
	// Strip control characters (including newlines)
	name = sanitizeControlChars(name)
	
	return "Hello, " + name + "\n"
}

// sanitizeControlChars removes or replaces control characters from input
func sanitizeControlChars(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsControl(r) {
			return -1 // Remove control characters
		}
		return r
	}, s)
}

// sanitizeForLogging sanitizes input for safe logging
func sanitizeForLogging(s string) string {
	if s == "" {
		return "<empty>"
	}
	return sanitizeControlChars(s)
}

func main() {
	// Create Server and Route Handlers
	r := mux.NewRouter()

	r.HandleFunc("/", handler)

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Start Server
	go func() {
		log.Println("Starting Server")
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	// Graceful Shutdown
	waitForShutdown(srv)
}

func waitForShutdown(srv *http.Server) {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal.
	<-interruptChan

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	srv.Shutdown(ctx)

	log.Println("Shutting down")
	os.Exit(0)
}
