package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/harshit3011/Go-React-Password-Manager/Server/internal/routes"
	database "github.com/harshit3011/Go-React-Password-Manager/Server/pkg/mongodb"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("No .env file found, using environment variables")
	}
	client := database.ConnectDB()
	database.Client = client
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalf("Couldn't connect to the database: %v", err)
	}
	log.Println("Successfully connected to the database!!")
	router := gin.Default()

	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Fatalf("Failed to disconnect from MongoDB: %v", err)
		}
	}()

	routes.BackendRoutes(router)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Server couldn't be connected!!")
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	log.Println("Shutting down server...")
}
