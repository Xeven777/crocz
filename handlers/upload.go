package handlers

import (
	"crocz/utils"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/bluele/gcache"
	"github.com/gofiber/fiber/v2"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type FileMeta struct {
	Path       string
	Filename   string
	Passphrase string
}

var FileCache gcache.Cache = gcache.New(100).
	LRU().
	Expiration(time.Minute * 10).
	EvictedFunc(func(key interface{}, value interface{}) {
		meta := value.(FileMeta)
		os.Remove(meta.Path)
	}).Build()

func UploadHandler(c *fiber.Ctx) error {
	log.Println("Upload request received")

	fileHeader, err := c.FormFile("file")
	if err != nil {
		log.Printf("Error getting file: %v", err)
		return c.Status(fiber.StatusBadRequest).SendString("No file uploaded")
	}

	passphrase := c.FormValue("passphrase")
	if passphrase == "" {
		log.Println("No passphrase provided")
		return c.Status(400).SendString("Passphrase required")
	}

	code, _ := gonanoid.New(6)
	log.Printf("Generated code: %s", code)

	// Save temporarily before encryption
	tempPath := filepath.Join("tmp", code+"_"+fileHeader.Filename)
	err = c.SaveFile(fileHeader, tempPath)
	if err != nil {
		log.Printf("Error saving file: %v", err)
		return c.Status(500).SendString("Failed to save file")
	}
	log.Printf("File saved to: %s", tempPath)

	// Encrypt the file
	encryptedPath, err := utils.EncryptFile(tempPath, code, passphrase)
	if err != nil {
		log.Printf("Encryption failed: %v", err)
		return c.Status(500).SendString("Encryption failed")
	}
	log.Printf("File encrypted to: %s", encryptedPath)

	// Store in cache
	FileCache.Set(code, FileMeta{
		Path:       encryptedPath,
		Filename:   fileHeader.Filename,
		Passphrase: passphrase,
	})

	log.Printf("Upload successful for code: %s", code)
	return c.JSON(fiber.Map{
		"code":     code,
		"filename": fileHeader.Filename,
	})
}
