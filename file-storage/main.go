package main

import (
	"context"
	_ "filestorage/docs"
	"filestorage/handlers"
	"filestorage/infrastructure"
	"filestorage/pkg/postgres"
	"filestorage/service"
	"github.com/gorilla/mux"
	swag "github.com/swaggo/http-swagger"
	"log"
	"net/http"
	"os"
	"os/signal"
	//_ "swagger-mux/docs"
	//_ "go-swag-demo-api/docs"
	//"github.com/swaggo/swag/"
	"time"
)

// @title File Storage API
// @version 1.0
// @description API для хранения и управления файлами

// @host localhost:8080
// @BasePath /
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

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/record/{id}", h.GetRecord)

	getAllRouter := sm.Methods(http.MethodGet).Subrouter()
	getAllRouter.HandleFunc("/record", h.GetAllRecords)

	sm.PathPrefix("/documentation/").Handler(swag.WrapHandler)

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
