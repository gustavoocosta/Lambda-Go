package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gustavoocosta/Lambda-Go/internal/handler"
)

func main() {
	lambda.Start(handler.Hello)
}
