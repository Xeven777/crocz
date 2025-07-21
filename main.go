package main

import (
	"crocz/handlers"
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Post("/upload", handlers.UploadHandler)
	app.Get("/download/:code", handlers.DownloadHandler)
	app.Static("/frontend", "./frontend")

	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws/:code", websocket.New(func(c *websocket.Conn) {
		code := c.Params("code")
		log.Printf("WebSocket connected for code: %s", code)

		defer func() {
			log.Printf("WebSocket disconnected for code: %s", code)
			c.Close()
		}()

		for {
			_, _, err := c.ReadMessage()
			if err != nil {
				log.Printf("WebSocket read error: %v", err)
				break
			}
		}
	}))

	log.Println("Server starting on :3000")
	log.Fatal(app.Listen(":3000"))
}
