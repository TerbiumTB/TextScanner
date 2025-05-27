package main

import (
	"context"
	_ "fileanalysis/docs"
	"fileanalysis/handlers"
	"fileanalysis/infrastructure"
	"fileanalysis/pkg/postgres"
	"fileanalysis/service"
	"github.com/gorilla/mux"
	"github.com/swaggo/http-swagger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// @title File Analysis API
// @version 1.0
// @description API для анализа файлов

// @host localhost:8081
// @BasePath /
func main() {
	l := log.New(os.Stdout, "API gateway ", log.LstdFlags)
	c := &http.Client{Timeout: 10 * time.Second}

	db, err := postgres.Init()
	if err != nil {
		l.Fatalf("[ERROR] couldn't establish postgress connection: %s", err)
	}

	stats, err := infrastructure.NewFileStatsDBX(db)

	images := infrastructure.NewLocalStorage(os.Getenv("IMAGES_PATH"))
	if err != nil {
		l.Fatalf("[ERROR] couldn't initialize schema: %s", err)
	}

	s := service.NewService(c, stats, images)
	//r := infrastructure.NewFileOriginalityDBX()

	h := handlers.NewHandler(l, c, s)
	sm := mux.NewRouter()

	//downloadRouter := sm.Methods(http.MethodGet).Subrouter()
	//downloadRouter.HandleFunc("/originality/{id}", h.CheckOriginalityHandler)

	statsRouter := sm.Methods(http.MethodGet).Subrouter()
	statsRouter.HandleFunc("/stats/{id}", h.GetStatsHandler)

	allStatsRouter := sm.Methods(http.MethodGet).Subrouter()
	allStatsRouter.HandleFunc("/stats", h.GetAllStatsHandler)

	wordCloudRouter := sm.Methods(http.MethodGet).Subrouter()
	wordCloudRouter.HandleFunc("/wordcloud/{id}", h.GetWordCloud)

	sm.PathPrefix("/documentation/").HandlerFunc(httpSwagger.WrapHandler)

	server := http.Server{
		Addr:         ":8080",
		Handler:      sm,
		ErrorLog:     l,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		l.Println("Starting server on port 8080")

		err := server.ListenAndServe()
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
	server.Shutdown(ctx)
}
