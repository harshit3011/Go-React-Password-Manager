package database

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client{
	err:=godotenv.Load(".env")
	if err!=nil{
		log.Fatalf("Error loading the environment variables: %v",err)
	}
	mongoUrl:=os.Getenv("MONGODB_URI")
	if mongoUrl==""{
		log.Fatalf("Mongo URL not set in the .env file")
	}
	clientOptions:= options.Client().ApplyURI(mongoUrl)
    client, err:=mongo.Connect(context.Background(),clientOptions)
	if err!=nil{
		log.Fatalf("Error Connecting to the database: %v",err)
	}
	return client
}

var Client *mongo.Client

func OpenCollections(collection string, client mongo.Client) *mongo.Collection{
	err:=godotenv.Load(".env")
	if err!=nil{
		log.Fatalf("Error loading the environment variables: %v",err)
	}
	database:=os.Getenv("DATABASE_NAME")
	if database==""{
		log.Fatalf("Database name not set in the environment variables")
	}
	databaseCollection:=client.Database(database).Collection(collection)
	return databaseCollection
}