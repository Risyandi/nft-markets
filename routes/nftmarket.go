package routes

import (
	"nftmarkets/controllers/nftmarketplace"

	"github.com/gin-gonic/gin"
)

func NftMarketsRoutes(route *gin.RouterGroup, nftmarketController *nftmarketplace.NftMarketplaceController) {
	route.GET("/items", nftmarketController.GetAllItems)
	route.GET("/items/:id", nftmarketController.GetItemByID)
	route.PUT("/items/:id", nftmarketController.UpdateItem)
	route.DELETE("/items/:id", nftmarketController.DeleteItem)
	route.POST("/items/create", nftmarketController.CreateItem)
	route.PUT("/items/purchase/:id", nftmarketController.PurchaseItem)
}
