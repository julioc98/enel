package main

import (
	"bytes"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/julioc98/enel/enel/api"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse
type Request events.APIGatewayProxyRequest

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(req Request) (Response, error) {
	var buf bytes.Buffer

	cpf := req.PathParameters["cpf"]
	id := req.PathParameters["id"]

	// body, err := json.Marshal(map[string]interface{}{
	// 	"message": "Go Serverless v1.0! Your function executed successfully!",
	// })
	body, err := getInfos(cpf, id)

	if err != nil {
		return Response{StatusCode: 404}, err
	}
	json.HTMLEscape(&buf, body)

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type":           "application/json",
			"X-MyCompany-Func-Reply": "hello-handler",
		},
	}

	return resp, nil
}

func getInfos(cpf, id string) ([]byte, error) {
	token, err := api.Login(cpf, id)
	if err != nil {
		return api.Resp, nil
	}
	if token == "" {
		return api.Resp, nil
	}

	info, err := api.Info(token)
	if err != nil {
		return api.Resp, nil
	}
	return []byte(info), nil
}

func main() {
	lambda.Start(Handler)
}
