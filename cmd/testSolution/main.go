package main

import (
	"main/internal/handlers"
	"main/internal/server"
	"main/internal/storage"
	"os"
)

func main() {
	storage, err := storage.NewStorage()
	if err != nil {
		os.Exit(1)
	}
	storage.Migrate()

	handler := handlers.NewHandler(storage)

	server.StartServer(handler)
}
