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
	uploadRouter.HandleFunc("/upload/{filename}", handlers.StorageHandler)

	downloadRouter := sm.Methods(http.MethodGet).Subrouter()
	downloadRouter.HandleFunc("/download/{id}", handlers.StorageHandler)

	getAllRouter := sm.Methods(http.MethodGet).Subrouter()
	getAllRouter.HandleFunc("/record", handlers.StorageHandler)

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/record/{id}", handlers.StorageHandler)

	//analyseRouter := sm.Methods(http.MethodGet).Subrouter()
	//analyseRouter.HandleFunc("/originality/{id}", handlers.AnalyseHandler)

	statsRouter := sm.Methods(http.MethodGet).Subrouter()
	statsRouter.HandleFunc("/stats/{id}", handlers.AnalyseHandler)

	allStatsRouter := sm.Methods(http.MethodGet).Subrouter()
	allStatsRouter.HandleFunc("/stats", handlers.AnalyseHandler)

	wordCloudRouter := sm.Methods(http.MethodGet).Subrouter()
	wordCloudRouter.HandleFunc("/wordcloud/{id}", handlers.AnalyseHandler)

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
