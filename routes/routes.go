package routes

import (
	"net/http"
	"nftmarkets/controllers/nftmarketplace"
	"nftmarkets/database/mongodb"
	"nftmarkets/helpers"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func SetupRouter(validate *validator.Validate) *gin.Engine {
	router := gin.New()

	// connection mongodb database
	mongodb.Connect()

	// skip base url path
	router.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		SkipPaths: []string{"/"},
	}))

	// register error handler RFC7807 spec
	router.Use(helpers.GinErrorHandler())

	// gin recovery handle middleware error and enable cors configuration
	router.Use(gin.Recovery())
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "PUT", "PATCH", "DELETE", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// endpoint API for health checking
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusCreated, "Everything is working fine..")
	})

	// endpoint for group api
	group := router.Group("api/v1")
	{
		NftMarketGroup := group.Group("nftmarkets")
		NftMarketsRoutes(NftMarketGroup, &nftmarketplace.NftMarketplaceController{Validate: validate})
	}

	return router
}
