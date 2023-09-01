package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"task/internal/delivery"
	"task/internal/service"
	"task/internal/storage"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
		<-ch
		cancel()
	}()
	Run(ctx)
}

func Run(ctx context.Context) {
	db, err := storage.InitDB()
	if err != nil {
		log.Println(err)
		return
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("can't close db err: %v\n", err)
		} else {
			log.Printf("db closed")
		}
	}()

	storages := storage.NewStorage(db)
	services := service.NewService(storages)
	handlers := delivery.NewHandler(services)
	handlers.InitRoutes()

	server := http.Server{
		Addr:         ":8080",
		Handler:      handlers.Mux,
		WriteTimeout: 2 * time.Second,
		ReadTimeout:  2 * time.Second,
	}

	fmt.Println("Starting server on http://localhost:8080")
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Println(err)
			cancel()
			return
		}
	}()
	<-ctx.Done()
	ctx, cancel = context.WithTimeout(ctx, 3*time.Minute)
	defer cancel()
	if err = server.Shutdown(ctx); err != nil {
		log.Println(err)
		return
	}
}
