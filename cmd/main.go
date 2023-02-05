package main

import (
	"os"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-redis/redis/v9"
	"github.com/joho/godotenv"
	"github.com/rootspyro/Dolar501/gen/models"
	"github.com/rootspyro/Dolar501/gen/restapi"
	"github.com/rootspyro/Dolar501/gen/restapi/operations"
	"github.com/rootspyro/Dolar501/github"
	"github.com/rootspyro/Dolar501/handlers"
	"github.com/rootspyro/Dolar501/services"
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

	//Github Authentication
	gh := github.NewGHAuth(
		os.Getenv("GH_CLIENT_ID"),
		os.Getenv("GH_CLIENT_SECRET"),
	)

	// services setup
	dolarSrv := services.NewDolarServices(rdClient)
	authSrv := services.NewAuthServices(gh)

	api.OauthSecurityAuth = func(token string, scopes []string) (interface{}, error) {
		ok, err := authSrv.ValidateToken(token)
		if err != nil {
			return nil, errors.New(401, "error authenticate")
		}
		if !ok {
			return nil, errors.New(401, "invalid token")
		}
		prin := models.Principal(token)
		return &prin, nil
	}

	// handlers setup
	api.DolarGetDolarCurrenciesHandler= handlers.NewGetCurrenciesImpl(dolarSrv)
	api.DolarGetCurrencyPlatformsHandler = handlers.NewGetPlatformsImpl(dolarSrv)
	api.DolarGetDolarPriceHandler = handlers.NewGetDolarPriceImpl(dolarSrv)
	api.AuthAuthLoginHandler = handlers.NewLoginImpl(gh)
	api.AuthGetAuthTokenHandler = handlers.NewCallbackImpl(gh, authSrv)

	// server port 
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	server.Port = port

	server.ConfigureAPI()

	server.Serve()
}
