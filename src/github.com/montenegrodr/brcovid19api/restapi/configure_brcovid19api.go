// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"strings"
	"crypto/tls"
	"net/http"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/montenegrodr/brcovid19api/restapi/operations"
	"github.com/montenegrodr/brcovid19api/models"
	"github.com/go-redis/redis"
)

//go:generate swagger generate server --target ../../brcovid19api --name Brcovid19api --spec ../../../../../swagger.yaml

func configureFlags(api *operations.Brcovid19apiAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}


func configureAPI(api *operations.Brcovid19apiAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	client := redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
	})

	api.GetCovid19ReportDataHandler = operations.GetCovid19ReportDataHandlerFunc(func(params operations.GetCovid19ReportDataParams) middleware.Responder {

		val, err := client.Get("data").Result()
		if err != nil {
			panic(err)
		}

		s := strings.Split(val, ";")

		if len(s) < 3 {
			panic("Could not fetch data")
		}

		confirmed, err := strconv.ParseInt(s[0], 10, 64)
		if err != nil {
			panic(err)
		}

		deceased, err := strconv.ParseInt(s[1], 10, 64)
		if err != nil {
			panic(err)
		}

		recovered, err := strconv.ParseInt(s[2], 10, 64)
		if err != nil {
			panic(err)
		}

		return operations.NewGetCovid19ReportDataOK().WithPayload(
			&models.Response{
				Confirmed: confirmed,
				Deceased: deceased,
				Recovered: recovered,
			})
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
