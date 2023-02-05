package main

import (
	"os"
	"strconv"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"github.com/go-openapi/loads"
	"github.com/joho/godotenv"
	"github.com/rootspyro/Dolar501/gen/restapi"
	"github.com/rootspyro/Dolar501/gen/restapi/operations"
)

func main(){

	// load environment variables
	godotenv.Load()	

	swaggerSpec, _ := loads.Analyzed(restapi.FlatSwaggerJSON, "")

	api := operations.NewDolar501API(swaggerSpec)
	server := restapi.NewServer(api)

	defer server.Shutdown()

	// server port 
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	server.Port = port

	server.ConfigureAPI()

	adapter := httpadapter.New(server.GetHandler())
	lambda.Start(adapter.Proxy)

}
