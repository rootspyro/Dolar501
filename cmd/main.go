package main

import (
	"log"
	"os"
	"strconv"

	"github.com/go-openapi/loads"
	"github.com/joho/godotenv"
	"github.com/rootspyro/Dollar-VzlAPI/gen/restapi"
	"github.com/rootspyro/Dollar-VzlAPI/gen/restapi/operations"
)

func main(){

	// load environment variables
	godotenv.Load()	

	swaggerSpec, _ := loads.Analyzed(restapi.FlatSwaggerJSON, "")

	api := operations.NewDollarVzlAPIAPI(swaggerSpec)
	server := restapi.NewServer(api)

	defer server.Shutdown()

	// server port 
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	server.Port = port

	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		log.Panicln(err.Error())
	}
	
}
