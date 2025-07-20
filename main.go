package main

import (
	"crocz/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Post("/upload", handlers.UploadHandler)
	app.Get("/download/:code", handlers.DownloadHandler)
	app.Static("/frontend", "./frontend")

	app.Listen(":3000")
}
