package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/space-event/email-service/internal"
	"github.com/space-event/email-service/internal/logger"
	"github.com/space-event/email-service/internal/service"
	pb "github.com/space-event/email-service/pkg/emailpb"
	"google.golang.org/grpc/reflection"

	"github.com/pelletier/go-toml/v2"
	"google.golang.org/grpc"
	"gopkg.in/gomail.v2"
)

func LoadConfig() (*internal.Config, error) {
	doc, err := os.ReadFile("config/config.toml")

	if err != nil {
		return nil, err
	}

	expanded := os.ExpandEnv(string(doc))

	var config internal.Config
	err = toml.Unmarshal([]byte(expanded), &config)

	if err != nil {
		return nil, err
	}

	return &config, nil
}

func main() {

	config, err := LoadConfig()

	if err != nil {
		log.Fatalf("Error to load config: %v", err)
	}

	logger.Init(config.LogLevel)

	listener, err := net.Listen("tcp", config.Addr)
	if err != nil {
		logger.Error("Error to create tcp connect", "error", err.Error())
		return
	}

	portSMTP, err := strconv.Atoi(config.Smtp.PortSMTP)
	if err != nil {
		logger.Error("Error to convert SMTP_PORT", "error", err.Error())
		return
	}

	mail := gomail.NewDialer(config.Smtp.HostSMTP, portSMTP, config.Smtp.EmailSMTP,
		config.Smtp.PasswordSMTP)

	grpcServer := grpc.NewServer()
	pb.RegisterEmailServiceServer(grpcServer, service.NewEmailServiceServer(mail, config.Smtp.EmailSMTP))

	reflection.Register(grpcServer)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		logger.Info("Email service server on", "address", config.Addr)
		if err = grpcServer.Serve(listener); err != nil {
			logger.Error("Error to server auth-service", "error", err.Error())
		}
	}()

	<-signalChan
	grpcServer.GracefulStop()
	logger.Info("Server is stopping...", "status", "graceful")
}
