package main

import (
	"crocz/handlers"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/olahol/melody"
)

var m = melody.New()

func main() {
	app := fiber.New()

	app.Post("/upload", handlers.UploadHandler)
	app.Get("/download/:code", handlers.DownloadHandler)
	app.Static("/frontend", "./frontend")

	app.Get("/ws/:code", adaptor.HTTPHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m.HandleRequest(w, r)
	})))

	app.Listen(":3000")
}
