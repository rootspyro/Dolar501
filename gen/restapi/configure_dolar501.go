// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/rs/cors"

	"github.com/rootspyro/Dolar501/gen/restapi/operations"
	"github.com/rootspyro/Dolar501/gen/restapi/operations/dolar"
	"github.com/rootspyro/Dolar501/middlewares"
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

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.DolarGetDolarAverageHandler == nil {
		api.DolarGetDolarAverageHandler = dolar.GetDolarAverageHandlerFunc(func(params dolar.GetDolarAverageParams) middleware.Responder {
			return middleware.NotImplemented("operation dolar.GetDolarAverage has not yet been implemented")
		})
	}
	if api.DolarGetDolarPlatformsHandler == nil {
		api.DolarGetDolarPlatformsHandler = dolar.GetDolarPlatformsHandlerFunc(func(params dolar.GetDolarPlatformsParams) middleware.Responder {
			return middleware.NotImplemented("operation dolar.GetDolarPlatforms has not yet been implemented")
		})
	}
	if api.DolarGetDolarPriceHandler == nil {
		api.DolarGetDolarPriceHandler = dolar.GetDolarPriceHandlerFunc(func(params dolar.GetDolarPriceParams) middleware.Responder {
			return middleware.NotImplemented("operation dolar.GetDolarPrice has not yet been implemented")
		})
	}

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
		AllowCredentials: true,
	})

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	
		middlw.RequestLogger(r)

		//Next
		handler = c.Handler(handler)
		handler.ServeHTTP(w,r)

	})
}
