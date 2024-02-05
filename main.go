package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/slaughtlaught/web-service-api/internal/infrasctructure/persistance"
	"github.com/slaughtlaught/web-service-api/internal/server"
	"github.com/slaughtlaught/web-service-api/internal/service"
	"github.com/slaughtlaught/web-service-api/pkg/dbx"
)

func main() {

	ctx := context.Background()
	pool, closefunc, err := dbx.NewStorage(ctx)
	if err != nil {
		log.Printf("error connecting to a database %v", err)
	}
	defer closefunc()

	noteSource := persistance.NewStorage(pool)
	noteService := service.NewNotes(noteSource)
	noteServer := server.NewNoteHandler(noteService)
	router := server.NewRouter(noteServer)

	httpServer := http.Server{
		Addr:              ":8080",
		WriteTimeout:      time.Second * 5,
		ReadTimeout:       time.Second * 5,
		ReadHeaderTimeout: time.Second * 5,
		IdleTimeout:       time.Minute,
		Handler:           router,
	}
	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatalf("http.ListenAndServe: %v", err)
	}
}
