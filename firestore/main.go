package main

import (
	"bytes"
	"encoding/json"
	"errors"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/julioc98/enel/firestore/api"
	joker "github.com/julioc98/enel/firestore/joker"
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
	var body []byte
	var err error
	var in joker.Joker
	var out joker.Joker

	switch req.HTTPMethod {
	case "POST":
		in, err = ByteToMap([]byte(req.Body))
		out, err = api.Add(in)
		body, err = json.Marshal(out)
	case "GET":
		body, err = json.Marshal(joker.Joker{
			"message": "Go JC GET!",
		})
	default:
		body, err = json.Marshal(joker.Joker{
			"message": "Go JC Default!",
		})
	}

	if err != nil {
		return Response{StatusCode: 404}, err
	}

	if body == nil {
		return Response{StatusCode: 500}, errors.New("Foi aqui 499")
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

// ByteToMap ...
func ByteToMap(byt []byte) (joker.Joker, error) {

	// byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	// We need to provide a variable where the JSON
	// package can put the decoded data. This
	// `map[string]interface{}` will hold a map of strings
	// to arbitrary data types.

	var dat map[string]interface{}

	// Here's the actual decoding, and a check for
	// associated errors.
	if err := json.Unmarshal(byt, &dat); err != nil {
		return nil, err
	}
	// fmt.Println(dat)
	return dat, nil

	// In order to use the values in the decoded map,
	// we'll need to convert them to their appropriate type.
	// For example here we convert the value in `num` to
	// the expected `float64` type.
	// num := dat["num"].(float64)
	// fmt.Println(num)

	// // Accessing nested data requires a series of
	// // conversions.
	// strs := dat["strs"].([]interface{})
	// str1 := strs[0].(string)
	// fmt.Println(str1)

}

func main() {
	lambda.Start(Handler)
}
