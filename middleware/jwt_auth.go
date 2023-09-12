package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang_boilerplate_with_gin/utils"
	"net/http"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		err := utils.ValidateJWT(context)
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
			fmt.Println(err)
			context.Abort()
			return
		}
		context.Next()
	}
}
