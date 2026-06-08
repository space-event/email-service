# Email Service

## 📖 Russian

[![Go Version](https://img.shields.io/badge/Go-1.26-blue.svg)](https://golang.org/)
[![gRPC](https://img.shields.io/badge/gRPC-latest-brightgreen.svg)](https://grpc.io/)
[![License](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

Email Service — это микросервис на Go для отправки электронных писем через gRPC с поддержкой SMTP (Gmail, любые SMTP-серверы).

**🇬🇧 [Read in English](./README.md)**

## ✨ Возможности

- 📧 Отправка email через gRPC API
- 📝 Поддержка plain text и HTML писем
- ⏱️ Graceful shutdown и контекстная отмена операций
- 📊 Структурированное логирование (slog)
- 🧪 Полное тестовое покрытие
- 🐳 Docker support


## 🚀 Быстрый старт

### Установка
```bash
# Клонирование репозитория
git clone https://github.com/space-event/email-service.git
cd email-service

# Установка зависимостей
go mod download

# Сборка
go build -o email-service cmd/main.go

# Запуск
./email-service
```

### Docker

```bash
# Сборка образа
docker build -t email-service .

# Запуск
docker run -d \
  -p 8081:8081 \
  -e ADDR=":8081"
  -e SMTP_HOST="smtp.gmail.com" \
  -e SMTP_PORT="587" \
  -e SMTP_EMAIL="sender@gmail.com" \
  -e SMTP_PASSWORD="your-password" \
  email-service
```

##  Переменные окружения

| Переменная | Описание | Значение по умолчанию | Обязательная |
|------------|----------|-----------------------|--------------|
| `ADDR` | Адрес и порт для запуска gRPC сервера | `-`                      |  ✅ |
| `LOG_LEVEL` | Уровень логирования (debug, info, warn, error) | `info`                | ❌ |
| `SMTP_HOST` | SMTP сервер для отправки писем | `-`                     | ✅ |
| `SMTP_PORT` | Порт SMTP сервера | `-`                     | ✅ |
| `SMTP_EMAIL` | Email отправителя | `-`                   | ✅ |
| `SMTP_PASSWORD` | Пароль приложения для SMTP | `-`                   | ✅ |

### Пример использования

```bash
export ADDR="localhost:50051"
export LOG_LEVEL="debug"
export SMTP_HOST="smtp.gmail.com"
export SMTP_PORT="587"
export SMTP_EMAIL="sender@gmail.com"
export SMTP_PASSWORD="your-app-password"
```

Подробная документация API и примеры использования доступны в директории [`/docs`](./docs):

## Содержание документации

- **[API Reference](./docs/API.md)** — Полная спецификация gRPC API, сгенерированная из `.proto` файлов
- **[Примеры использования](./docs/EXAMPLES.md)** — Примеры кода на Go, gRPCurl и других языках


