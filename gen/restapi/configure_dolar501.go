// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"
	"os"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-redis/redis/v9"
	"github.com/rs/cors"

	"github.com/rootspyro/Dolar501/gen/models"
	"github.com/rootspyro/Dolar501/gen/restapi/operations"
	"github.com/rootspyro/Dolar501/github"
	"github.com/rootspyro/Dolar501/handlers"
	"github.com/rootspyro/Dolar501/middlewares"
	"github.com/rootspyro/Dolar501/services"
)

//go:generate swagger generate server --target ../../gen --name Dolar501 --spec ../../swagger/swagger.yml --principal interface{} --exclude-main

func configureFlags(api *operations.Dolar501API) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.Dolar501API) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

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

	// Authentication
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

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	middlw := middlewares.NewMiddlewares()

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"Origin","Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	
		middlw.RequestLogger(r)

		//Next
		handler = c.Handler(handler)
		handler.ServeHTTP(w,r)

	})
}
