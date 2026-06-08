package tests

import (
	"context"
	"net"
	"testing"

	smtpmock "github.com/mocktools/go-smtp-mock/v2"
	"github.com/space-event/email-service/internal/logger"
	"github.com/space-event/email-service/internal/service"
	pb "github.com/space-event/email-service/pkg/emailpb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
	"gopkg.in/gomail.v2"
)

const bufSize = 1024 * 1024

func setupTestServer(t *testing.T) (*grpc.ClientConn, pb.EmailServiceClient, func()) {

	logger.Init("debug")

	listner := bufconn.Listen(bufSize)

	smtpServer := smtpmock.New(smtpmock.ConfigurationAttr{
		MsgSizeLimit: 1024 * 1024,
	})

	if err := smtpServer.Start(); err != nil {
		t.Fatal(err)
	}

	mail := gomail.NewDialer("localhost", smtpServer.PortNumber(), "", "")
	emailServer := service.NewEmailServiceServer(mail, "test@example.com")

	grpcServer := grpc.NewServer()
	pb.RegisterEmailServiceServer(grpcServer, emailServer)

	go func() {
		if err := grpcServer.Serve(listner); err != nil {
			t.Logf("Server exited: %v", err)
		}
	}()

	conn, err := grpc.DialContext(context.Background(), "bufnet",
		grpc.WithContextDialer(func(ctx context.Context, s string) (net.Conn, error) {
			return listner.Dial()
		}),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	require.NoError(t, err)

	clientR := pb.NewEmailServiceClient(conn)

	cleanup := func() {
		conn.Close()
		grpcServer.Stop()
		smtpServer.Stop()
	}

	return conn, clientR, cleanup
}

func TestSend_ValidRequest(t *testing.T) {

	_, client, cleanup := setupTestServer(t)
	defer cleanup()

	req := &pb.EmailRequest{
		EmailTarget: "target@example.com",
		MessageText: "Test message",
		Subject:     "Test",
		ContentType: "text/html",
	}

	res, err := client.Send(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, true, res.Success)
}

func TestSend_InvalidEmail(t *testing.T) {

	_, client, cleanup := setupTestServer(t)
	defer cleanup()

	req := &pb.EmailRequest{
		EmailTarget: "targexample.com",
		MessageText: "Test message",
		Subject:     "Test",
		ContentType: "text/html",
	}

	res, err := client.Send(context.Background(), req)

	assert.NotNil(t, err)
	assert.Nil(t, res)
}

func TestSend_InvalidMessage(t *testing.T) {

	_, client, cleanup := setupTestServer(t)
	defer cleanup()

	req := &pb.EmailRequest{
		EmailTarget: "targexample.com",
		MessageText: "",
		Subject:     "Test",
		ContentType: "text/html",
	}

	res, err := client.Send(context.Background(), req)

	assert.Equal(t, status.Error(codes.InvalidArgument, "message is required"), err)
	assert.Nil(t, res)
}

func TestSend_EmptyContentType(t *testing.T) {

	_, client, cleanup := setupTestServer(t)
	defer cleanup()

	req := &pb.EmailRequest{
		EmailTarget: "targexample.com",
		MessageText: "Test message",
		Subject:     "Test",
		ContentType: "",
	}

	res, err := client.Send(context.Background(), req)

	assert.Equal(t, status.Error(codes.InvalidArgument, "content type is required"), err)
	assert.Nil(t, res)
}

func TestSend_UnknownContentType(t *testing.T) {

	_, client, cleanup := setupTestServer(t)
	defer cleanup()

	req := &pb.EmailRequest{
		EmailTarget: "targexample.com",
		MessageText: "Test message",
		Subject:     "Test",
		ContentType: "something",
	}

	res, err := client.Send(context.Background(), req)

	assert.Equal(t, status.Error(codes.InvalidArgument, "unknown type content type"), err)
	assert.Nil(t, res)
}
