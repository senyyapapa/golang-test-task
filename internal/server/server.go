package server

import (
	"log"
	"main/internal/handlers"

	"github.com/gofiber/fiber/v3"
)

func StartServer(h *handlers.Handler) {
	app := fiber.New()

	app.Post("/post_number", h.PostNum)
	app.Get("/get_numbers", h.GetNumber)

	log.Fatal(app.Listen(":8080"))

}
