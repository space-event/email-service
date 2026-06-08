# Email Service

## 📖 English

[![Go Version](https://img.shields.io/badge/Go-1.26-blue.svg)](https://golang.org/)
[![gRPC](https://img.shields.io/badge/gRPC-latest-brightgreen.svg)](https://grpc.io/)
[![License](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

**🇷🇺 [Читать на русском](./README.ru.md)**

Email Service is a Go microservice for sending emails via gRPC with SMTP support (Gmail, any SMTP servers).

## Features
- 📧 Send emails via gRPC API
- 📝 Support plain text and HTML emails
- ⏱️ Graceful shutdown with context cancellation
- 📊 Structured logging (slog)
- 🧪 Full test coverage
- 🐳 Docker support


## 🚀Quick Start

### Installation
```bash
# Clone repository
git clone https://github.com/space-event/email-service.git
cd email-service

# Install dependencies
go mod download

# Build
go build -o email-service cmd/main.go

# Run
./email-service
```

### Docker

```bash
# Build image
docker build -t email-service .

# Run image
docker run -d \
  -p 8081:8081 \
  -e ADDR=":8081"
  -e SMTP_HOST="smtp.gmail.com" \
  -e SMTP_PORT="587" \
  -e SMTP_EMAIL="sender@gmail.com" \
  -e SMTP_PASSWORD="your-password" \
  email-service
```

## Environment Variables

| Variable | Description | Default | Required |
|------------|----------|-----------------------|--------------|
| `ADDR` | gRPC server address and port | `-`                      |  ✅ |
| `LOG_LEVEL` | Logging level (debug, info, warn, error) | `info`                | ❌ |
| `SMTP_HOST` | SMTP server host | `-`                     | ✅ |
| `SMTP_PORT` | SMTP server port | `-`                     | ✅ |
| `SMTP_EMAIL` | Sender email address | `-`                   | ✅ |
| `SMTP_PASSWORD` | SMTP app password | `-`                   | ✅ |

### Usage Example

```bash
export ADDR="localhost:50051"
export LOG_LEVEL="debug"
export SMTP_HOST="smtp.gmail.com"
export SMTP_PORT="587"
export SMTP_EMAIL="sender@gmail.com"
export SMTP_PASSWORD="your-app-password"
```

For detailed API documentation and usage examples, see [`/docs`](./docs):

## Contents of the documentation

- **[API Reference](./docs/API.md)** — Full gRPC API specification
- **[Примеры использования](./docs/EXAMPLES.md)** — Code examples in Go, gRPCurl and other languages


