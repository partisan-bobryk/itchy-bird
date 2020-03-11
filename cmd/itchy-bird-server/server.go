package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

func StartServer() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish")
	flag.Parse()

	// Figure out where files are located
	binaryRepo, binaryRepoErr := MakeLocalBinaryRepository(DefaultFilePath)

	if binaryRepoErr != nil {
		log.Panic(binaryRepoErr)
		return
	}

	apiHandler := &ApiRouteHandlers{binaryRepo}

	router := mux.NewRouter()

	// Register Middleware
	router.Use(LoggingMiddleware)

	// Register Routes
	router.HandleFunc("/versions", apiHandler.GeVersionsHandler).Methods("GET")
	router.HandleFunc("/download/{fileName}", apiHandler.DownloadHandler).Methods("POST")

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
		log.Printf("Starting server on port %s \n", "8080")
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
