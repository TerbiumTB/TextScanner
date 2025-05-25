package main

import (
	"context"
	"filestorage/handlers"
	"filestorage/infrastructure"
	"filestorage/pkg/postgres"
	"filestorage/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	lg := log.New(os.Stdout, "File storage ", log.LstdFlags)

	storage := infrastructure.NewLocalStorage(os.Getenv("STORAGE_ROOT"))
	//repo := infrastructure.NewFileMap()

	db, err := postgres.Init()
	if err != nil {
		lg.Fatalf("[ERROR] couldn't establish postgress connection: %s", err)
	}

	repo, err := infrastructure.NewFileDBX(db)

	if err != nil {
		lg.Fatalf("[ERROR] couldn't create schema: %s", err)
	}

	ser := service.NewService(repo, storage)

	h := handlers.NewHandler(lg, ser)
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
		ErrorLog:     lg,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		lg.Println("Starting server on port 8080")

		err := s.ListenAndServe()
		if err != nil {
			lg.Fatalf("Error starting server: %s", err)
			//log.Printf("Error starting server: %s\n", err)
			//os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <-c
	lg.Println("Got signal:", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
