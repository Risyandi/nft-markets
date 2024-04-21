package nftmarketplace

import (
	"context"
	"net/http"
	"nftmarkets/config"
	"nftmarkets/database/mongodb"
	"nftmarkets/entity"

	"github.com/gin-gonic/gin"
)

// TODO Rest API:
// - I can see all the items for the related user [-]
// - I can get a single item [-]
// - I can create new entries [X]
// - I can update information of any of my items [-]
// - I can delete any item [-]
// - A purchasing endpoint that whenever is called reduces the availability of the item, and that fails if there is no availability [-]

type NftMarketplaceController struct {
	Config config.KeyViperConfig
}

// CreateItem creates a new item
func (h *NftMarketplaceController) CreateItem(ctx *gin.Context) {
	var formdataProductNft entity.ProductNft

	// Bind the request body to the newItem variable
	if err := ctx.BindJSON(&formdataProductNft); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Insert the new item
	// Accessing collections within the database
	_, err := mongodb.Database.Collection("product_nft").InsertOne(context.Background(), formdataProductNft)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create item"})
		return
	}

	ctx.JSON(http.StatusCreated, formdataProductNft)
}
