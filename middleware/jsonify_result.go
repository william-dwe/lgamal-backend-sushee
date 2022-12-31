package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func JSONifyResult() gin.HandlerFunc {
	return func(c *gin.Context) {
        c.Writer.Header().Set("Content-Type", "application/json")
        c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(http.StatusNoContent)
        }

        c.Next()
    }
} 