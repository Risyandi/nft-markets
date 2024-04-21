package mongodb

import (
	"context"
	"log"
	"nftmarkets/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Database *mongo.Database

func Connect() {
	// get value config environment
	config := config.ConfigViper()
	mongoDatabaseName := config.GetString("database.mongo.database")
	mongoHost := config.GetString("database.mongo.host")
	mongoPort := config.GetString("database.mongo.port")

	ctx, cancel := context.WithTimeout(context.Background(), 25*time.Second)
	defer cancel()

	// Connect to MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://" + mongoHost + mongoPort)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Accessing specific database
	Database = client.Database(mongoDatabaseName)

	log.Println("Successfully Connected to MongoDB!")
}
