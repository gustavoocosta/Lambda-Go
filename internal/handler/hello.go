package handler

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

func Hello(
	ctx context.Context,
	request events.APIGatewayProxyRequest,
) (events.APIGatewayProxyResponse, error) {

	name := request.QueryStringParameters["name"]

	if name == "" {
		name = "Visitante"
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body: fmt.Sprintf(
			`{"message":"Olá %s! Lambda em Go funcionando 🚀"}`,
			name,
		),
	}, nil
}
