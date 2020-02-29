package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish")
	flag.Parse()

	router := mux.NewRouter()

	// Register Middleware
	router.Use(loggingMiddleware)

	// Register Routes
	router.HandleFunc("/versions", geVersionsHandler).Methods("GET")

	srv := &http.Server{
		Addr: "0.0.0.0:8080",
		// Timeouts
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}

	// Start go server
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accpet gracefull shutdowns when quit via SIGNINT,
	// however SIGKILL, SIGQUIT will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block undil we recieve our signal
	<-c

	// Create a deadline to wait for
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wiat until the timeout deadline
	srv.Shutdown(ctx)

	log.Println("Server shutting down")
	os.Exit(0)
}

// Middleware

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		// Log clients routes
		log.Printf("%s - %s\n", req.Method, req.RequestURI)

		next.ServeHTTP(res, req)
	})
}

// Route Handlers

func geVersionsHandler(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	fmt.Fprintf(res, "Looking good!")
}
