package router_helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AppResponse struct {
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data"`
}

func GenerateResponseMessage(c *gin.Context, content interface{}) {
	formattedMessage := AppResponse{
		StatusCode: http.StatusOK,
		Data:       content,
	}
	c.JSON(http.StatusOK, formattedMessage)
}
