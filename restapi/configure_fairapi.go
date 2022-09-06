// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"unico-challenge/restapi/operations"

	"github.com/sirupsen/logrus"
)

//go:generate swagger generate server --target ../../unico-challenge --name Fairapi --spec ../api/fair/fair.yml --principal interface{}

func configureFlags(api *operations.FairapiAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.FairapiAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})

	logger := logrus.New()
	logger.Formatter = &logrus.JSONFormatter{}
	logger.Out = os.Stdout
	file, err := os.OpenFile("../../log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logger.Out = file
	} else {
		logger.Info("Failed to log to file, using default stderr")
	}
	log.SetOutput(logger.Writer())
	api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()
	api.TxtProducer = runtime.TextProducer()

	if api.GetHealthHandler == nil {
		api.GetHealthHandler = operations.GetHealthHandlerFunc(func(params operations.GetHealthParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetHealth has not yet been implemented")
		})
	}
	if api.CreateFairHandler == nil {
		api.CreateFairHandler = operations.CreateFairHandlerFunc(func(params operations.CreateFairParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.CreateFair has not yet been implemented")
		})
	}
	if api.DeleteFairHandler == nil {
		api.DeleteFairHandler = operations.DeleteFairHandlerFunc(func(params operations.DeleteFairParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.DeleteFair has not yet been implemented")
		})
	}
	if api.GetFairHandler == nil {
		api.GetFairHandler = operations.GetFairHandlerFunc(func(params operations.GetFairParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetFair has not yet been implemented")
		})
	}
	if api.GetFairsHandler == nil {
		api.GetFairsHandler = operations.GetFairsHandlerFunc(func(params operations.GetFairsParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetFairs has not yet been implemented")
		})
	}
	if api.UpdateFairHandler == nil {
		api.UpdateFairHandler = operations.UpdateFairHandlerFunc(func(params operations.UpdateFairParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.UpdateFair has not yet been implemented")
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
	return handler
}
