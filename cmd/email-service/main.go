package main

import (
	"github.com/space-event/email-service/internal"
	"github.com/space-event/email-service/internal/service"
	pb "github.com/space-event/email-service/pkg/emailpb"
	"log"
	"net"
	"os"
	"strconv"

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

	listener, err := net.Listen("tcp", config.Addr)
	if err != nil {
		log.Fatal(err.Error())
	}

	portSMTP, err := strconv.Atoi(config.Smtp.PortSMTP)
	if err != nil {
		log.Fatalf("Error to convert SMTP_PORT: %v", err)
	}

	mail := gomail.NewDialer(config.Smtp.HostSMTP, portSMTP, config.Smtp.EmailSMTP,
		config.Smtp.PasswordSMTP)

	grpcServer := grpc.NewServer()
	pb.RegisterEmailServiceServer(grpcServer, service.NewEmailServiceServer(mail, config.Smtp.EmailSMTP))

	log.Printf("Start server on  %s", config.Addr)
	if err = grpcServer.Serve(listener); err != nil {
		log.Fatalf("Error to start auth-service: %v", err)
	}

}
