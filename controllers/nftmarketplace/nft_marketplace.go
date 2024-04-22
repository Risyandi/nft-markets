package nftmarketplace

import (
	"context"
	"log"
	"net/http"
	"nftmarkets/config"
	"nftmarkets/database/mongodb"
	"nftmarkets/entity"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/**
* ===============================
* TODO Rest API:
* ===============================
* a I can see all the items for the related user [X]
* b I can get a single item [X]
* - Retrieve my items with rating = X
* - Retrieve my items with X reputationBadge
* - Retrieve my items with availability of more/less than X
* - Retrieve my items with X category
* c I can create new entries [X]
* d I can update information of any of my items [x]
* e I can delete any item [x]
* f A purchasing endpoint that whenever is called reduces the availability of the item, and that fails if there is no availability [x]
* ===============================
**/

var Validate *validator.Validate

type NftMarketplaceController struct {
	Config   config.KeyViperConfig
	Validate *validator.Validate
}

// GetReputationBadgeFilter returns the reputation badge filter
func GetReputationBadgeFilter(repBadge string) interface{} {
	switch repBadge {
	case "red":
		return bson.M{"$lte": 500}
	case "yellow":
		return bson.M{"$lte": 799, "$gt": 500}
	case "green":
		return bson.M{"$gt": 799}
	default:
		return nil
	}
}

// CreateItem creates a new item
func (h *NftMarketplaceController) CreateItem(ctx *gin.Context) {
	var formdataProductNft entity.ProductNft

	// Bind the request body to the newItem variable
	if err := ctx.BindJSON(&formdataProductNft); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Validate the input
	if err := h.Validate.Struct(formdataProductNft); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error_validate": err.Error()})
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

// GetAllItems returns all items
func (h *NftMarketplaceController) GetAllItems(ctx *gin.Context) {
	var formdataProductNft []entity.ProductNft

	filter := bson.M{}

	// Optional query parameters
	// ?rating=
	if ratingStr := ctx.Query("rating"); ratingStr != "" {
		rating, err := strconv.Atoi(ratingStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid rating parameter"})
			return
		}
		filter["rating"] = rating
	}

	// ?reputation=
	if repBadge := ctx.Query("reputationBadge"); repBadge != "" {
		repBadgeFilter := GetReputationBadgeFilter(repBadge)
		if repBadgeFilter != nil {
			filter["reputation"] = repBadgeFilter
		}
	}

	// ?availability=
	if availabilityStr := ctx.Query("availability"); availabilityStr != "" {
		availability, err := strconv.Atoi(availabilityStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid availability parameter"})
			return
		}
		filter["availability"] = availability
	}

	// ?category=
	if category := ctx.Query("category"); category != "" {
		filter["category"] = category
	}

	// Find the items in the database
	cursor, err := mongodb.Database.Collection("product_nft").Find(context.Background(), filter)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get items"})
		return
	}
	defer cursor.Close(context.Background())

	// Iterate over the cursor and decode each item into a new item variable
	for cursor.Next(context.Background()) {
		var itemProductNft entity.ProductNft
		if err := cursor.Decode(&itemProductNft); err != nil {
			log.Printf("Failed to decode item: %v", err)
			continue
		}
		formdataProductNft = append(formdataProductNft, itemProductNft)
	}

	// Check if there are any items
	lengthData := len(formdataProductNft)
	if lengthData == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return

	}

	ctx.JSON(http.StatusOK, formdataProductNft)
}

// GetItemByID returns an item by its ID
func (h *NftMarketplaceController) GetItemByID(ctx *gin.Context) {
	// Get the item ID from the URL params
	id := ctx.Param("id")
	itemID, err := primitive.ObjectIDFromHex(id)

	// Check if the item ID is valid
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
		return
	}

	// Find the item in the database
	var itemProductNft entity.ProductNft
	err = mongodb.Database.Collection("product_nft").FindOne(context.Background(), bson.M{"_id": itemID}).Decode(&itemProductNft)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	ctx.JSON(http.StatusOK, itemProductNft)
}

// UpdateItem updates an item in the database
func (h *NftMarketplaceController) UpdateItem(ctx *gin.Context) {
	id := ctx.Param("id")
	itemID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
		return
	}

	var updatedItem entity.ProductNft
	if err := ctx.BindJSON(&updatedItem); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Validate the input
	if err := h.Validate.Struct(updatedItem); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error_validate": err.Error()})
		return
	}

	// Update the item
	_, err = mongodb.Database.Collection("product_nft").ReplaceOne(context.Background(), bson.M{"_id": itemID}, updatedItem)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update item"})
		return
	}

	ctx.JSON(http.StatusOK, updatedItem)
}

// DeleteItem deletes an item from the database
func (h *NftMarketplaceController) DeleteItem(ctx *gin.Context) {
	id := ctx.Param("id")
	itemID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
		return
	}

	// Delete the item
	_, err = mongodb.Database.Collection("product_nft").DeleteOne(context.Background(), bson.M{"_id": itemID})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete item"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Item deleted"})
}

// PurchaseItem purchases an item from the database
func (h *NftMarketplaceController) PurchaseItem(c *gin.Context) {
	id := c.Param("id")
	itemID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
		return
	}

	var item entity.ProductNft
	err = mongodb.Database.Collection("product_nft").FindOne(context.Background(), bson.M{"_id": itemID}).Decode(&item)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	if item.Availability <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Item is out of stock"})
		return
	}

	// Reduce availability
	item.Availability--
	_, err = mongodb.Database.Collection("product_nft").ReplaceOne(context.Background(), bson.M{"_id": itemID}, item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update item"})
		return
	}

	c.JSON(http.StatusOK, item)
}
