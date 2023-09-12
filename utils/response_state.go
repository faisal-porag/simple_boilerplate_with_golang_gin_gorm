package utils

import "github.com/gin-gonic/gin"

type CommonResponse struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Lang    string      `json:"lang"`
	Data    interface{} `json:"data,omitempty"`
}

type ResponseState struct {
	StatusCode int
	Code       string
	MessageEn  string
	MessageBn  string
}

func (rs ResponseState) CommonResponse(lang string, data interface{}) CommonResponse {
	var message string
	if lang == "bn" {
		message = rs.MessageBn
	} else {
		message = rs.MessageEn
	}

	return CommonResponse{
		Code:    rs.Code,
		Message: message,
		Lang:    lang,
		Data:    data,
	}
}

func (rs ResponseState) WriteToResponse(c *gin.Context, lang string, data interface{}) {
	c.JSON(rs.StatusCode, rs.CommonResponse(lang, data))
}
