package handlers

import (
	"crocz/utils"
	"os"

	"github.com/gofiber/fiber/v2"
)

func DownloadHandler(c *fiber.Ctx) error {
	code := c.Params("code")
	passphrase := c.Query("passphrase")

	if code == "" || passphrase == "" {
		return c.Status(400).SendString("Code and passphrase are required")
	}

	value, err := FileCache.Get(code)
	if err != nil {
		return c.Status(404).SendString("Invalid or expired code")
	}

	meta := value.(FileMeta)
	if passphrase != meta.Passphrase {
		return c.Status(403).SendString("Incorrect passphrase")
	}

	stream, err := utils.DecryptFile(meta.Path, passphrase)
	if err != nil {
		return c.Status(500).SendString("Decryption failed")
	}
	defer stream.Close()

	c.Set("Content-Type", "application/octet-stream")
	c.Set("Content-Disposition", "attachment; filename=\""+meta.Filename+"\"")

	defer func() {
		FileCache.Remove(code)
		os.Remove(meta.Path)
	}()

	return c.SendStream(stream)
}
