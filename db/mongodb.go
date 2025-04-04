package db

import (
	"context"
	"log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var CompanyCollection *mongo.Collection

// SetupDatabase initializes the MongoDB client and collection
func SetupDatabase() {
	clientOptions := options.Client().ApplyURI("mongodb+srv://iamviqar:Mva786@verifyhire.outzf6f.mongodb.net/?retryWrites=true&w=majority&appName=verifyhire")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	Client = client
	CompanyCollection = client.Database("companyDB").Collection("companies")
}
