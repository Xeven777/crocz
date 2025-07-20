package handlers

import (
	"crocz/utils"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/olahol/melody"
)

var m = melody.New()

func DownloadHandler(c *fiber.Ctx) error {
	code := c.Params("code")
	passphrase := c.Query("passphrase")

	// Validate inputs
	if code == "" || passphrase == "" {
		return c.Status(400).SendString("Code and passphrase are required")
	}

	// Lookup the file metadata in cache
	value, err := FileCache.Get(code)
	if err != nil {
		return c.Status(404).SendString("Invalid or expired code")
	}

	meta := value.(FileMeta)
	if passphrase != meta.Passphrase {
		return c.Status(403).SendString("Incorrect passphrase")
	}

	// Decrypt file stream
	stream, err := utils.DecryptFile(meta.Path, passphrase)
	if err != nil {
		return c.Status(500).SendString("Decryption failed")
	}
	defer stream.Close()

	// Set headers and stream the decrypted file
	c.Set("Content-Type", "application/octet-stream")
	c.Set("Content-Disposition", "attachment; filename=\""+meta.Filename+"\"")

	m.BroadcastFilter([]byte("downloaded:"+code), func(q *melody.Session) bool {
		return q.Request.URL.Path == "/ws/"+code
	})

	// Optional: delete after download
	defer func() {
		FileCache.Remove(code)
		os.Remove(meta.Path)
	}()

	return c.SendStream(stream)
}
