# API Usage Examples

## Send Plain Text Email

```go
req := &emailpb.EmailRequest{
    EmailTarget: "user@example.com",
    Subject:     "Hello from Email Service",
    MessageText: "This is a plain text email",
    ContentType: "text/plain",
}

res, err := client.Send(ctx, req)
```

## Send HTML Email
```go
req := &emailpb.EmailRequest{
    EmailTarget: "user@example.com",
    Subject:     "HTML Email Example",
    MessageText: "<h1>Welcome!</h1><p>This is <b>HTML</b> email</p>",
    ContentType: "text/html",
}

res, err := client.Send(ctx, req)
```

## gRPCurl Example

```bash
grpcurl -plaintext -d '{
  "emailTarget": "test@example.com",
  "subject": "Test",
  "messageText": "Hello",
  "contentType": "text/plain"
}' localhost:50051 email.EmailService/Send
```




