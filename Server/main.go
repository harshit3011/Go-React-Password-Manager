package main

import (
	"context"
	"log"
	"os"

	"github.com/harshit3011/Go-React-Password-Manager/Server/database"
	"github.com/joho/godotenv"
)

func main() {
	err:=godotenv.Load(".env")
	if err != nil {
		log.Println("No .env file found, using environment variables")
	}
	client:=database.ConnectDB()
	database.Client=client
	port:=os.Getenv("PORT")
	if port==""{
		port="8081"
	}
	err=client.Ping(context.Background(),nil)
	if err!=nil{
		log.Fatalf("Couldn't connect to the database: %v",err)
	}
	log.Println("Successfully connected to the database!!")
}