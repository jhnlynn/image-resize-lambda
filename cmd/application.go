package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/umtkas/image-resizer-lambda/internal/aws"
)

// build and upload to lambda:
// GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main cmd/application.go
// zip -r main.zip main
func main() {
	lambda.Start(aws.Handler)
}
