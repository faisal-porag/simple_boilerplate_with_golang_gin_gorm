package main

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"golang_boilerplate_with_gin/config"
	"golang_boilerplate_with_gin/controllers"
	"golang_boilerplate_with_gin/db"
	"golang_boilerplate_with_gin/middleware"
	"golang_boilerplate_with_gin/models"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	cfg := loadEnv()
	loadDatabase(cfg)
	serveApplication(cfg)
}

func loadEnv() *config.Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal().Err(err).Msg("Error loading .env file")
	}

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Config parsing failed")
	}
	return cfg
}

func loadDatabase(cfg *config.Config) {
	_, err := db.ConnectPostgres(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("db connection error. please check db connection")
	}
	_ = db.Database.AutoMigrate(&models.User{})
	_ = db.Database.AutoMigrate(&models.Entry{})
}

func serveApplication(cfg *config.Config) {
	router := gin.Default()

	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", controllers.Register)
	publicRoutes.POST("/login", controllers.Login)

	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(middleware.JWTAuthMiddleware())
	protectedRoutes.POST("/entry", controllers.AddEntry)
	protectedRoutes.GET("/entry", controllers.GetAllEntries)

	_ = router.Run(os.Getenv("APPLICATION_PORT"))
	fmt.Printf("Server running on port %v", cfg.ApplicationPort)
}
