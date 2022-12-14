package middleware

import "github.com/gin-gonic/gin"

func JSONifyResult() gin.HandlerFunc {
	return func(c *gin.Context) {
        c.Writer.Header().Set("Content-Type", "application/json")
        c.Next()
    }
} 