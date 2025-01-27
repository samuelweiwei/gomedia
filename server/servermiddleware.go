package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Middleware to limit the concurency of the server
func ConcurrencyLimitMiddleware(c *gin.Context) {
	if err := sem.Acquire(c.Request.Context(), 1); err != nil {
		c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests"})
		c.Abort()
		return
	}
	defer sem.Release(1)
	c.Next()
}
