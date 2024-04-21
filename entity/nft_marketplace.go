package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/**
* ===============================
* TODO validation API:
* ===============================
* When a user tries to save some data, you must validate against the following
* constraints:
* a. Item name: An item name cannot contain the words ["Sex", "Gay", "Lesbian"] and it should be longer than 10 characters [X]
* b. Rating: The rating MUST accept only integers, where rating is >= 0 and <= 5. [X]
* c. Category: The category can be one of [photo, sketch, cartoon, animation] and it should be a string [X]
* d. Image: The image MUST be a valid URL [X]
* e. Reputation: The reputation MUST be an integer >= 0 and <= 1000. [X]
* - If reputation is <= 500 the value is red
* - If reputation is <= 799 the value is yellow
* - Otherwise the value is green
* f. Price: The price must be an integer [X]
* g. Availability: The availability must be an integer [X]
* ===============================
**/

type ProductNft struct {
	ID           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	ItemName     string             `json:"item_name" bson:"item_name" validate:"required,min=10,forbiddenwords"`
	Rating       int                `json:"rating" bson:"rating" validate:"required,min=0,max=5"`
	Category     string             `json:"category" bson:"category" validate:"required,oneof=photo sketch cartoon animation"`
	Image        string             `json:"image" bson:"image" validate:"required,url"`
	Reputation   int                `json:"reputation" bson:"reputation" validate:"required,min=0,max=1000"`
	Price        int                `json:"price" bson:"price" validate:"required"`
	Availability int                `json:"availability" bson:"availability" validate:"required"`
}
