package app

import (
	// "fmt"
	"net/http"

	routing "github.com/go-ozzo/ozzo-routing"
	"github.com/go-ozzo/ozzo-routing/slash"
	// "github.com/go-ozzo/ozzo-routing/fault"

	"github.com/keiraqz/WWC/api"
)

type launchStep func()

var launchSteps = []launchStep{
	setupDB,
	initializeDAOs,
	initializeHTTPClient,
	setupAPIClients,
	serveAPI,
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

// Run handles the main logic for the application.  It cycles through all of the
// launch steps that the application needs to perform as it is launching.
func Run() {

	for _, ls := range launchSteps {
		ls()
	}
}
