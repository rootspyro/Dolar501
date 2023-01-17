package main

import (
	"log"
	"os"
	"strconv"
	"github.com/go-openapi/loads"
	"github.com/go-redis/redis/v9"
	"github.com/joho/godotenv"
	"github.com/rootspyro/Dollar-VzlAPI/gen/restapi"
	"github.com/rootspyro/Dollar-VzlAPI/gen/restapi/operations"
	"github.com/rootspyro/Dollar-VzlAPI/handlers"
	"github.com/rootspyro/Dollar-VzlAPI/services"
)

func main(){

	// load environment variables
	godotenv.Load()	

	swaggerSpec, _ := loads.Analyzed(restapi.FlatSwaggerJSON, "")

	api := operations.NewDolar501API(swaggerSpec)
	server := restapi.NewServer(api)


	defer server.Shutdown()

	// redis client
	db, _ := strconv.Atoi(os.Getenv("REDIS_DB"))

	rdClient := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB: db,
	}) 

	// services setup
	dolarSrv := services.NewDolarServices(rdClient)

	// handlers setup
	api.DolarGetDolarPlatformsHandler = handlers.NewGetPlatformsImpl(dolarSrv)
	api.DolarGetDolarPriceHandler = handlers.NewGetDolarPriceImpl(dolarSrv)
	api.DolarGetDolarAverageHandler = handlers.NewGetDolarAverageImpl(dolarSrv)

	// server port 
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	server.Port = port

	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		log.Panicln(err.Error())
	}
	
}
