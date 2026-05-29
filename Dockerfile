FROM golang:latest

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY .. .

RUN go build -o email-service ./cmd/email-service/main.go

CMD ["./email-service"]

