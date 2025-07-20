package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")
	
	// Sanitize name for logging by removing control characters
	controlCharsRegex := regexp.MustCompile(`[\x00-\x1F\x7F]`)
	sanitizedName := controlCharsRegex.ReplaceAllString(name, " ")
	
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
	
	// Limit accepted name length to 100 characters
	if len(name) > 100 {
		name = name[:100]
	}
	
	// Strip newline and control characters, replace with space
	// This regex matches control characters (ASCII 0-31) except space (32)
	controlCharsRegex := regexp.MustCompile(`[\x00-\x1F\x7F]`)
	name = controlCharsRegex.ReplaceAllString(name, " ")
	
	return "Hello, " + name + "\n"
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
