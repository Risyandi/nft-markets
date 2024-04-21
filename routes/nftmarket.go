package routes

import (
	"nftmarkets/controllers/nftmarketplace"

	"github.com/gin-gonic/gin"
)

func NftMarketsRoutes(route *gin.RouterGroup, nftmarketController *nftmarketplace.NftMarketplaceController) {
	route.POST("/create", nftmarketController.CreateItem)
}
