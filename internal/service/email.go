package service

import (
	"context"
	"net/mail"
	"strings"

	"github.com/space-event/email-service/internal/logger"
	pb "github.com/space-event/email-service/pkg/emailpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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

	err := validRequest(request)

	if err != nil {
		logger.Error("Invalid request", "error", err.Error())
		return nil, err
	}

	err = validEmail(request.EmailTarget)
	if err != nil {
		logger.Error("Invalid email", "error", err.Error())
		return nil, err
	}

	m := gomail.NewMessage()
	m.SetHeader("From", s.fromEmail)
	m.SetHeader("To", request.EmailTarget)
	m.SetHeader("Subject", request.Subject)
	m.SetBody(request.ContentType, request.MessageText)

	logger.Info("Sending email",
		"from", s.fromEmail,
		"to", maskEmail(request.EmailTarget),
		"subject", request.Subject,
		"content_type", request.ContentType,
	)

	if err = s.sendWithContext(ctx, m); err != nil {
		logger.Error("Failed to send email", "to", maskEmail(request.EmailTarget), "error", err.Error())
		return &pb.EmailResponse{Success: false}, err
	}

	return &pb.EmailResponse{Success: true}, nil
}

func (s *EmailServiceServer) sendWithContext(ctx context.Context, message *gomail.Message) error {

	errChan := make(chan error, 1)

	go func() {
		errChan <- s.mail.DialAndSend(message)
	}()

	select {
	case err := <-errChan:
		if err != nil {
			return status.Errorf(codes.Internal, "failed to send email: %v", err)
		}
		return nil
	case <-ctx.Done():
		return status.Error(codes.Canceled, "email sending cancelled")
	}
}

func validRequest(request *pb.EmailRequest) error {

	if request == nil {
		return status.Error(codes.InvalidArgument, "request is nil")
	}

	if request.EmailTarget == "" {
		return status.Error(codes.InvalidArgument, "email target is required")
	}

	if request.MessageText == "" {
		return status.Error(codes.InvalidArgument, "message is required")
	}

	if request.ContentType == "" {
		return status.Error(codes.InvalidArgument, "content type is required")
	}

	if request.ContentType != "text/plain" && request.ContentType != "text/html" {
		return status.Error(codes.InvalidArgument, "unknown type content type")
	}

	return nil
}

func validEmail(email string) error {
	_, err := mail.ParseAddress(email)
	return err
}

func maskEmail(email string) string {
	parts := strings.Split(email, "@")
	return parts[0][:2] + "***@" + parts[1]
}
