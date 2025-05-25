package main

import (
	"apigateway/handlers"
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	l := log.New(os.Stdout, "API gateway ", log.LstdFlags)
	sm := mux.NewRouter()

	uploadRouter := sm.Methods(http.MethodPost).Subrouter()
	uploadRouter.HandleFunc("/upload/{filename}", handlers.UploadHandler)

	downloadRouter := sm.Methods(http.MethodGet).Subrouter()
	downloadRouter.HandleFunc("/download/{id}", handlers.DownloadHandler)

	downloadAllRouter := sm.Methods(http.MethodGet).Subrouter()
	downloadAllRouter.HandleFunc("/download", handlers.DownloadHandler)

	analyseRouter := sm.Methods(http.MethodGet).Subrouter()
	analyseRouter.HandleFunc("/analyse/{id}", handlers.AnalyseHandler)

	s := http.Server{
		Addr:         ":8000",
		Handler:      sm,
		ErrorLog:     l,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		l.Println("Starting server on port 8000")

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	signal.Notify(ch, os.Kill)

	sig := <-ch
	log.Println("Got signal:", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
