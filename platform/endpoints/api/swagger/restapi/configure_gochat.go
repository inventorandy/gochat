// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"log"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"gochat/platform/endpoints/api/controller"
	"gochat/platform/endpoints/api/swagger/models"
	"gochat/platform/endpoints/api/swagger/restapi/operations"
	"gochat/platform/endpoints/api/swagger/restapi/operations/stable"
	"gochat/platform/internal/proto/pb"
)

//go:generate swagger generate server --target ../../swagger --name Gochat --spec ../swagger.yml --principal models.Principal --exclude-main

// user - Create the User for all methods
var user *pb.User

func configureFlags(api *operations.GochatAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.GochatAPI) http.Handler {
	// Create the Controller Instance
	handlers, err := controller.NewHandlerController()
	if err != nil {
		log.Fatalln(err.Error())
	}

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

	// Applies when the "Authorization" header is set
	api.JwtAuth = func(token string) (*models.Principal, error) {
		user, err = handlers.AuthenticateJWT(token)
		if err != nil {
			return nil, err
		}

		// Return the Token
		return (*models.Principal)(&token), nil
		//return nil, errors.NotImplemented("api key auth (jwt) Authorization from header param [Authorization] has not yet been implemented")
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()

	// [POST /auth/login]
	api.StablePostAuthLoginHandler = stable.PostAuthLoginHandlerFunc(func(params stable.PostAuthLoginParams, principal *models.Principal) middleware.Responder {
		return handlers.AuthLoginPost(params)
	})

	// [POST /account]
	api.StablePostAccountHandler = stable.PostAccountHandlerFunc(func(params stable.PostAccountParams, principal *models.Principal) middleware.Responder {
		return handlers.AccountPost(params)
	})

	// [PUT /account]
	api.StablePutAccountHandler = stable.PutAccountHandlerFunc(func(params stable.PutAccountParams, principal *models.Principal) middleware.Responder {
		return handlers.AccountPut(user, params)
	})

	// [GET /account]
	api.StableGetAccountHandler = stable.GetAccountHandlerFunc(func(params stable.GetAccountParams, principal *models.Principal) middleware.Responder {
		return handlers.AccountGet(user, params)
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
