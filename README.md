# Crocz - Secure File Sharing

A simple and secure file sharing application built with Go and Fiber. Upload files with encryption and share them using temporary codes.

## Features

- **Secure File Upload**: Upload files with passphrase protection
- **End-to-End Encryption**: Files are encrypted using AES encryption
- **Temporary Sharing**: Files expire after 10 minutes for security
- **One-Time Download**: Files are automatically deleted after download
- **Simple Web Interface**: Clean HTML frontend for easy file sharing
- **Containerized**: Docker support for easy deployment

## Tech Stack

- **Backend**: Go with Fiber framework
- **Encryption**: AES encryption with age library
- **Caching**: In-memory cache with automatic expiration
- **Frontend**: Pure HTML/CSS/JavaScript
- **Containerization**: Docker with multi-stage builds

## Quick Start

### Option 1: Run with Docker (Recommended)

1. Build the Docker image:

   ```bash
   docker build -t crocz-app .
   ```

2. Run the container:

   ```bash
   docker run -p 3000:3000 crocz-app
   ```

3. Open your browser and navigate to:
   ```
   http://localhost:3000/frontend/upload.html
   ```

### Option 2: Run Locally

1. Make sure you have Go 1.24+ installed

2. Clone and navigate to the project:

   ```bash
   git clone <your-repo-url>
   cd crocz
   ```

3. Install dependencies:

   ```bash
   go mod download
   ```

4. Run the application:

   ```bash
   go run main.go
   ```

5. Open your browser and navigate to:
   ```
   http://localhost:3000/frontend/upload.html
   ```

## How to Use

### Upload a File

1. Go to `/frontend/upload.html`
2. Select a file to upload
3. Enter a secure passphrase
4. Click "Upload File"
5. Copy the generated code and share it securely

### Download a File

1. Go to `/frontend/download.html`
2. Enter the 6-character code
3. Enter the correct passphrase
4. Click "Download File"
5. The file will download automatically

## API Endpoints

### POST /upload

Upload a file with encryption.

**Request:**

- `Content-Type`: `multipart/form-data`
- `file`: File to upload
- `passphrase`: String for encryption

**Response:**

```json
{
  "code": "abc123",
  "filename": "document.pdf"
}
```

### GET /download/:code

Download an encrypted file.

**Parameters:**

- `code`: 6-character file code
- `passphrase`: Query parameter for decryption

**Response:**

- File stream with proper headers for download

## Security Features

- Files are encrypted with AES encryption before storage
- Passphrases are required for both upload and download
- Files automatically expire after 10 minutes
- Files are deleted immediately after download
- Temporary storage in encrypted format only

## File Structure

```
crocz/
├── main.go              # Main application entry point
├── go.mod               # Go module dependencies
├── dockerfile           # Docker container configuration
├── frontend/            # Web interface
│   ├── upload.html      # File upload page
│   └── download.html    # File download page
├── handlers/            # HTTP request handlers
│   ├── upload.go        # File upload logic
│   └── download.go      # File download logic
├── utils/               # Utility functions
│   └── encryption.go    # File encryption/decryption
└── tmp/                 # Temporary encrypted file storage
```

## Environment Configuration

The application runs on port 3000 by default. You can modify this in `main.go` if needed.

## Docker Deployment

The included Dockerfile uses multi-stage builds for optimal image size:

- Build stage: Compiles the Go application
- Runtime stage: Minimal Alpine Linux with the compiled binary

## Security Notes

- Never share passphrases through insecure channels
- Files are automatically cleaned up for security
- Use strong passphrases for better protection
- The application is designed for temporary file sharing only

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test thoroughly
5. Submit a pull request
