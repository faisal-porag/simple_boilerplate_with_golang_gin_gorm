package main

import (
	"fmt"
	"golang_boilerplate_with_gin/controllers"
	"golang_boilerplate_with_gin/db"
	"golang_boilerplate_with_gin/middleware"
	"golang_boilerplate_with_gin/models"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	loadDatabase()
	serveApplication()
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func loadDatabase() {
	_ = db.ConnectPostgres()
	_ = db.Database.AutoMigrate(&models.User{})
	_ = db.Database.AutoMigrate(&models.Entry{})
}

func serveApplication() {
	router := gin.Default()

	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", controllers.Register)
	publicRoutes.POST("/login", controllers.Login)

	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(middleware.JWTAuthMiddleware())
	protectedRoutes.POST("/entry", controllers.AddEntry)
	protectedRoutes.GET("/entry", controllers.GetAllEntries)

	_ = router.Run(os.Getenv("APPLICATION_PORT"))
	fmt.Printf("Server running on port %v", os.Getenv("APPLICATION_PORT"))
}
