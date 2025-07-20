package handlers

import (
	"crocz/utils"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type FileMeta struct {
	Path       string
	Filename   string
	Passphrase string
}

var FileStore = make(map[string]FileMeta) // code → metadata

func UploadHandler(c *fiber.Ctx) error {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("No file uploaded")
	}

	passphrase := c.FormValue("passphrase")
	if passphrase == "" {
		return c.Status(400).SendString("Passphrase required")
	}

	code, _ := gonanoid.New(6)

	// Save temporarily before encryption
	tempPath := filepath.Join("tmp", code+"_"+fileHeader.Filename)
	err = c.SaveFile(fileHeader, tempPath)
	if err != nil {
		return c.Status(500).SendString("Failed to save file")
	}

	// Encrypt the file
	encryptedPath, err := utils.EncryptFile(tempPath, code, passphrase)
	if err != nil {
		return c.Status(500).SendString("Encryption failed")
	}

	// Store in map
	FileStore[code] = FileMeta{
		Path:       encryptedPath,
		Filename:   fileHeader.Filename,
		Passphrase: passphrase,
	}

	return c.JSON(fiber.Map{
		"code":     code,
		"filename": fileHeader.Filename,
	})
}
