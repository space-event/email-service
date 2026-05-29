package service

import (
	"context"
	pb "github.com/space-event/email-service/pkg/emailpb"
	"log"

	"gopkg.in/gomail.v2"
)

type EmailServiceServer struct {
	pb.UnimplementedEmailServiceServer
	mail      *gomail.Dialer
	fromEmail string
}

func NewEmailServiceServer(mail *gomail.Dialer, fromEmail string) *EmailServiceServer {
	return &EmailServiceServer{mail: mail, fromEmail: fromEmail}
}

func (s *EmailServiceServer) Send(ctx context.Context, request *pb.EmailRequest) (*pb.EmailResponse, error) {
	m := gomail.NewMessage()
	m.SetHeader("From", s.fromEmail)
	m.SetHeader("To", request.EmailTarget)
	m.SetHeader("Subject", request.Subject)
	m.SetBody(request.ContentType, request.MessageText)

	log.Printf("Send message\nFrom: %s\nTo: %s\nSubject: %s\n",
		s.fromEmail, request.EmailTarget, request.Subject)

	if err := s.mail.DialAndSend(m); err != nil {
		log.Fatalf("Error to send message: %v", err)
		return &pb.EmailResponse{}, err
	}

	return &pb.EmailResponse{Success: true}, nil
}
