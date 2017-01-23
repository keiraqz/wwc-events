package api

import "github.com/go-ozzo/ozzo-routing"

// HealthCheck implements a simple health check so that an external third party
// can verify that the application is currently up and functional
func HelloWorld(c *routing.Context) error {
	c.Write("Hellow World for WWC!")
	return nil
}
