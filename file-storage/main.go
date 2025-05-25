package main

import (
	"context"
	"filestorage/handlers"
	"filestorage/infrastructure"
	"filestorage/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	log := log.New(os.Stdout, "File storage ", log.LstdFlags)

	storage := infrastructure.NewLocalStorage(os.Getenv("STORAGE_ROOT"))
	repo := infrastructure.NewFileMap()
	ser := service.NewService(repo, storage)

	h := handlers.NewHandler(log, ser)
	sm := mux.NewRouter()

	uploadRouter := sm.Methods(http.MethodPost).Subrouter()
	uploadRouter.HandleFunc("/upload/{filename}", h.Upload)

	downloadRouter := sm.Methods(http.MethodGet).Subrouter()
	downloadRouter.HandleFunc("/download/{id}", h.Download)

	downloadAllRouter := sm.Methods(http.MethodGet).Subrouter()
	downloadAllRouter.HandleFunc("/download", h.DownloadAll)

	s := http.Server{
		Addr:         ":8080",
		Handler:      sm,
		ErrorLog:     log,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		log.Println("Starting server on port 8080")

		err := s.ListenAndServe()
		if err != nil {
			log.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <-c
	log.Println("Got signal:", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
