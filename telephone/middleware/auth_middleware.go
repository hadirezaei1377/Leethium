package middleware

import (
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Check authentication token or session
		// If not authenticated, return unauthorized error
		// Otherwise, continue to next handler
		ctx.Next()
	}
}
