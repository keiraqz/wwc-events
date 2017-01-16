package api

import "github.com/go-ozzo/ozzo-routing"

// HealthCheck implements a simple health check so that an external third party
// can verify that the application is currently up and functional
func HealthCheck(c *routing.Context) error {
	c.Response.Header().Add("Version", "v0.1")
	return nil
}
