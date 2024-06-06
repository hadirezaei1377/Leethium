package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize MongoDB connection
	// Initialize other dependencies such as repositories and services

	router := gin.Default()

	// Initialize controllers
	userController := controllers.NewUserController()

	// Define routes
	router.POST("/register", userController.Register)
	router.POST("/login", userController.Login)

	// Run the server
	router.Run(":8080")
}
