package main

import (
	"context"
	"fileanalysis/handlers"
	"fileanalysis/infrastructure"
	"fileanalysis/pkg/postgres"
	"fileanalysis/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	l := log.New(os.Stdout, "API gateway ", log.LstdFlags)
	c := &http.Client{Timeout: 10 * time.Second}

	db, err := postgres.Init()
	if err != nil {
		l.Fatalf("[ERROR] couldn't establish postgress connection: %s", err)
	}

	stats, err := infrastructure.NewFileStatsDBX(db)
	if err != nil {
		l.Fatalf("[ERROR] couldn't initialize schema: %s", err)
	}

	s := service.NewService(c, stats)
	//r := infrastructure.NewFileOriginalityDBX()

	h := handlers.NewHandler(l, c, s)
	sm := mux.NewRouter()

	//downloadRouter := sm.Methods(http.MethodGet).Subrouter()
	//downloadRouter.HandleFunc("/originality/{id}", h.CheckOriginalityHandler)

	statsRouter := sm.Methods(http.MethodGet).Subrouter()
	statsRouter.HandleFunc("/stats/{id}", h.GetStatsHandler)

	allStatsRouter := sm.Methods(http.MethodGet).Subrouter()
	allStatsRouter.HandleFunc("/stats", h.GetAllStatsHandler)

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
