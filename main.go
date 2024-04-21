package main

import (
	"nftmarkets/config"
	"nftmarkets/middleware"
	"nftmarkets/routes"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func main() {
	/**
	*   1. Config environment
	*   2. Routes
	*   3. Middleware
	**/

	// Initialize the validator
	validate = validator.New()
	config := config.ConfigViper()
	router := routes.SetupRouter(validate)

	if config.GetBool("server.log") {
		router.Use(middleware.RequestLoggerMiddleware())
	}

	router.Run(config.GetString("server.port"))
}
