package handlers

import (
	"crocz/utils"
	"os"

	"github.com/gofiber/fiber/v2"
)

func DownloadHandler(c *fiber.Ctx) error {
	code := c.Params("code")
	passphrase := c.Query("passphrase")

	// Validate inputs
	if code == "" || passphrase == "" {
		return c.Status(400).SendString("Code and passphrase are required")
	}

	// Lookup the file metadata
	meta, ok := FileStore[code]
	if !ok {
		return c.Status(404).SendString("Invalid code")
	}

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

	// Optional: delete after download
	defer func() {
		os.Remove(meta.Path)
		delete(FileStore, code)
	}()

	return c.SendStream(stream)
}
