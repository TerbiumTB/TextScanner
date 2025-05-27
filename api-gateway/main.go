package main

import (
	//_ "apigateway/docs/"
	"apigateway/handlers"
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// @title API Gateway
// @version 1.0
// @description API Gateway

// @host localhost:8000
// @BasePath /
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
	http.FileServer(http.Dir("./static/"))

	wordCloudRouter := sm.Methods(http.MethodGet).Subrouter()
	wordCloudRouter.HandleFunc("/wordcloud/{id}", handlers.AnalyseHandler)

	//sm.PathPrefix("/documentation/").HandlerFunc(httpSwagger.Handler(
	//	httpSwagger.URL("./docs/swagger.json")))
	//sm.HandleFunc("./docs/swagger.json", func(w http.ResponseWriter, r *http.Request) {
	//	http.ServeFile(w, r, "./docs/swagger.json")
	//})

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
