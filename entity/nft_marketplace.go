package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductNft struct {
	ID           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	ItemName     string             `json:"item_name" bson:"item_name"`
	Rating       int                `json:"rating" bson:"rating"`
	Category     string             `json:"category" bson:"category"`
	Image        string             `json:"image" bson:"image"`
	Reputation   int                `json:"reputation" bson:"reputation"`
	Price        int                `json:"price" bson:"price"`
	Availability int                `json:"availability" bson:"availability"`
}
