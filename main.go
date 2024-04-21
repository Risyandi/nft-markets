package main

import (
	"nftmarkets/config"
	"nftmarkets/middleware"
	"nftmarkets/routes"
)

func main() {
	/**
	*   1. Config environment
	*   2. Routes
	*   3. Middleware
	**/

	config := config.ConfigViper()
	router := routes.SetupRouter()

	if config.GetBool("server.log") {
		router.Use(middleware.RequestLoggerMiddleware())
	}

	router.Run(config.GetString("server.port"))
}
