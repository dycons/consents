// This file is safe to edit. Once it exists it will not be overwritten
package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gobuffalo/pop"

	"github.com/dycons/consents/consents-service/api/restapi/handlers"
	"github.com/dycons/consents/consents-service/api/restapi/operations"
	"github.com/dycons/consents/consents-service/utilities/log"
)

//go:generate swagger generate server --target ../../swagger-tests --name ConsentsService --spec ../swagger.yaml

func configureFlags(api *operations.ConsentsServiceAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.ConsentsServiceAPI) http.Handler {
	// Initialize custom logger. Configuration to the logger can be made here through this (or a similar) function
	log.Init()

	// Connect to the database
	tx, err := pop.Connect("development")
	if err != nil {
		log.Write(nil, 0, err).Panic("Failed to connect to database")
	}

	// configure the api here
	api.ServeError = errors.ServeError

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.GetOneDefaultConsentHandler = operations.GetOneDefaultConsentHandlerFunc(func(params operations.GetOneDefaultConsentParams) middleware.Responder {
		return handlers.GetOneDefaultConsent(params, tx)
	})
	api.GetProjectConsentsByParticipantHandler = operations.GetProjectConsentsByParticipantHandlerFunc(func(params operations.GetProjectConsentsByParticipantParams) middleware.Responder {
		return handlers.GetProjectConsentsByParticipant(params, tx)
	})
	api.AddRemsEntitlementHandler = operations.AddRemsEntitlementHandlerFunc(func(params operations.AddRemsEntitlementParams) middleware.Responder {
		return handlers.AddRemsEntitlement(params, tx)
	})
	api.PostParticipantHandler = operations.PostParticipantHandlerFunc(func(params operations.PostParticipantParams) middleware.Responder {
		return handlers.PostParticipant(params, tx)
	})
	api.PutProjectConsentHandler = operations.PutProjectConsentHandlerFunc(func(params operations.PutProjectConsentParams) middleware.Responder {
		return handlers.PutProjectConsent(params, tx)
	})

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
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
