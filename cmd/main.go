package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/gustavoocosta/Lambda-Go/internal/handler"
)

func main() {
	response, err := handler.Hello(
		context.Background(),
		events.APIGatewayProxyRequest{
			QueryStringParameters: map[string]string{
				"name": "Gustavo",
			},
		},
	)

	if err != nil {
		panic(err)
	}

	fmt.Println(response.Body)
}
