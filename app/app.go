package app

import (
	// "fmt"
	"net/http"

	routing "github.com/go-ozzo/ozzo-routing"
	"github.com/go-ozzo/ozzo-routing/slash"
	// "github.com/go-ozzo/ozzo-routing/fault"

	"github.com/keiraqz/wwc-events/api"
)

type launchStep func()

var launchSteps = []launchStep{
	// setupDB,
	// initializeDAOs,
	// initializeHTTPClient,
	// setupAPIClients,
	serveAPI,
}

var done chan error

// Run handles the main logic for the application.  It cycles through all of the
// launch steps that the application needs to perform as it is launching.
func Run() {

	for _, ls := range launchSteps {
		ls()
	}
}

// ServeAPI serves the API up over HTTP
func serveAPI() {
	router := routing.New()

	// declare routing groups
	health := router.Group("/wwc")
	main := router.Group("/wwc")

	configureHealthCheckRouteGroup(health)
	configureMainRouteGroup(main)

	http.Handle("/", router)

	done = make(chan error)

	go func() {
		done <- http.ListenAndServe(":3939", nil)
	}()

	// result := <-done

	// encountered an error while server is up -- stop immediately
	// logger.Fatalf("An unexpected error occurred: %v", result)
}

func configureHealthCheckRouteGroup(health *routing.RouteGroup) {
	health.Use(
		slash.Remover(http.StatusMovedPermanently),
		// fault.Recovery(logger.Errorf),
	)

	health.Get(
		"/health-check",
		api.HealthCheck,
	)
}

func configureMainRouteGroup(co *routing.RouteGroup) {
	co.Use(
		slash.Remover(http.StatusMovedPermanently),
		// fault.Recovery(logger.Errorf),
	)

	co.Get(
		"/hello-world",
		api.HelloWorld,
	)
}

// Shutdown shuts down the application
func Shutdown() {
	done <- nil
}
