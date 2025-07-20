package utils

import (
	"io"
	"os"
	"path/filepath"

	"filippo.io/age"
)

// Encrypts the file at inputPath and saves it at tmp/<code>.age
func EncryptFile(inputPath, code, passphrase string) (string, error) {
	// Open input file
	in, err := os.Open(inputPath)
	if err != nil {
		return "", err
	}
	defer in.Close()

	// Create encrypted file
	encryptedPath := filepath.Join("tmp", code+".age")
	out, err := os.Create(encryptedPath)
	if err != nil {
		return "", err
	}
	defer out.Close()

	// Create age writer
	recipient, err := age.NewScryptRecipient(passphrase)
	if err != nil {
		return "", err
	}
	encryptor, err := age.Encrypt(out, recipient)
	if err != nil {
		return "", err
	}

	// Copy contents
	_, err = io.Copy(encryptor, in)
	if err != nil {
		return "", err
	}
	encryptor.Close()

	os.Remove(inputPath)

	return encryptedPath, nil
}

// DecryptFile decrypts an .age file and returns a read stream
func DecryptFile(encryptedPath, passphrase string) (io.ReadCloser, error) {
	file, err := os.Open(encryptedPath)
	if err != nil {
		return nil, err
	}

	identity, err := age.NewScryptIdentity(passphrase)
	if err != nil {
		file.Close()
		return nil, err
	}

	reader, err := age.Decrypt(file, identity)
	if err != nil {
		file.Close()
		return nil, err
	}

	return io.NopCloser(reader), nil
}
