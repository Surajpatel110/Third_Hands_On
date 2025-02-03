package main

import (
	"go-restful-app/handler"
	"go-restful-app/middleware"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Set Gin mode for production
	gin.SetMode(gin.ReleaseMode)

	// Created router
	r := gin.Default()

	// Setted trusted proxies (can adjust as needed)
	err := r.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		log.Fatalf("Failed to set trusted proxies: %v", err)
	}

	// Adding middleware to the router..This middleware is a custom function from the middleware package that logs details of incoming HTTP requests for monitoring and debugging.
	r.Use(middleware.RequestLogger())

	// Register routes
	handler.RegisterRoutes(r) //This function is expected to define and register all the application's API endpoints with the router r

	// Start server
	r.Run(":8081")
}
