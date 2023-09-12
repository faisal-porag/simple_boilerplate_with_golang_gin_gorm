package httpserver

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang_boilerplate_with_gin/state"
	"net/http"
)

func Server(s *state.State) {
	router := gin.Default()

	// Enable CORS with the specified options
	router.Use(corsMiddleware())

	// Serve static files from the "uploads" directory
	router.Static("/uploads", "./uploads")

	// SERVICE HEALTH CHECK
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"app_port": s.Cfg.ApplicationPort,
			"status":   "Application successfully running .....",
		})
	})
	
	_ = router.Run(s.Cfg.ApplicationPort)
}

func corsMiddleware() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "PATCH"}
	config.AllowHeaders = []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Accept-Language"}

	return cors.New(config)
}
