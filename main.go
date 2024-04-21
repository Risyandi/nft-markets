package main

import (
	"fmt"
	"nftmarkets/config"
	"nftmarkets/middleware"
	"nftmarkets/routes"
	"strings"

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
	// Registering the custom validator
	if err := validate.RegisterValidation("forbiddenwords", forbiddenWordsValidator); err != nil {
		fmt.Printf("Failed to register custom validation: %v\n", err)
		return
	}

	config := config.ConfigViper()
	router := routes.SetupRouter(validate)

	if config.GetBool("server.log") {
		router.Use(middleware.RequestLoggerMiddleware())
	}

	router.Run(config.GetString("server.port"))
}

func forbiddenWordsValidator(fl validator.FieldLevel) bool {
	// List of forbidden words
	forbiddenWords := []string{"Sex", "Gay", "Lesbian"}

	// Get the field value
	fieldValue := fl.Field().String()

	// Convert the string to lowercase for case-insensitive matching
	fieldValue = strings.ToLower(fieldValue)

	// Check if any of the forbidden words are present
	for _, word := range forbiddenWords {
		if strings.Contains(fieldValue, strings.ToLower(word)) {
			return false
		}
	}

	return true
}
