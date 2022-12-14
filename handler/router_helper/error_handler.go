package router_helper

import (
	"final-project-backend/errorlist"

	"github.com/gin-gonic/gin"
)

func GenerateErrorMessage(c *gin.Context, err error) {
	if appErr, isAppError := err.(errorlist.AppError); isAppError {
		c.AbortWithStatusJSON(appErr.StatusCode, appErr)
		return
	}
	serverErr := errorlist.InternalServerError()
	c.AbortWithStatusJSON(serverErr.StatusCode, serverErr)
}
