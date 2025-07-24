FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o crocz main.go

# Use a minimal image for running
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/crocz .
COPY frontend ./frontend
COPY tmp ./tmp
COPY handlers ./handlers
COPY utils ./utils
EXPOSE 3000
CMD ["./crocz"]