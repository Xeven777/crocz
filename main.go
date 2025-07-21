package main

import (
	"crocz/handlers"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New(fiber.Config{
		BodyLimit: 100 * 1024 * 1024, // 100MB limit
	})

	app.Use(recover.New())
	app.Use(cors.New())

	app.Post("/upload", handlers.UploadHandler)
	app.Get("/download/:code", handlers.DownloadHandler)
	app.Static("/frontend", "./frontend")

	log.Println("Server starting on :3000")
	log.Fatal(app.Listen(":3000"))
}
