package main

import (
	"blog/middleware"
	"blog/repository"
	"blog/service"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// Init DB
	repository.Init()
	postService := service.NewPostService()

	// Request
	id := request.PathParameters["id"]
	var result interface{}
	var err error

	switch true {
	case request.HTTPMethod == "GET" && len(id) > 0:
		result, err = postService.GetOne(id)
	case request.HTTPMethod == "GET" && len(id) == 0:
		result, err = postService.GetAll()
	default:
	}
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	if isAuthorized := middleware.CheckJWT(); !isAuthorized {
		return events.APIGatewayProxyResponse{
			Body:       string("Access Denied"),
			StatusCode: http.StatusUnauthorized,
		}, nil
	}

	switch true {
	case request.HTTPMethod == "POST":
		result, err = postService.Post(request.Body)
	case request.HTTPMethod == "PATCH":
		result, err = postService.Edit(request.Body)
	case request.HTTPMethod == "DELETE":
		result, err = postService.Delete(id)
	default:
		return events.APIGatewayProxyResponse{
			Body:       "Route Not Found",
			StatusCode: http.StatusNotFound,
		}, nil
	}

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	// Json
	bytes, err := json.Marshal(result)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		Body:       string(bytes),
		StatusCode: http.StatusOK,
	}, nil
}

func main() {
	lambda.Start(handler)
}
