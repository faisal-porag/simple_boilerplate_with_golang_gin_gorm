package utils

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func GetLanguageFromRequest(c *gin.Context) string {
	lang := c.GetHeader("Accept-Language")
	lang = strings.ToLower(lang)
	if lang == "bn" || lang == "ban" || lang == "bang" || lang == "bangla" || lang == "bengali" {
		return "bn"
	} else {
		return "en"
	}
}
