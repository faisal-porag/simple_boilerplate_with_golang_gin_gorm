package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"golang_boilerplate_with_gin/controllers/example_controller"
	"golang_boilerplate_with_gin/db"
	"log"
)

func main() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	router := gin.Default()
	dbHandler, dbErr := db.Init(dbUrl)
	if dbErr != nil {
		log.Fatal(dbErr)
	}

	example_controller.RegisterRoutes(router, dbHandler)

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"port":  port,
			"dbUrl": dbUrl,
		})
	})

	_ = router.Run(port)
}
