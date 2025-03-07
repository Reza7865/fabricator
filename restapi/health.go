package restapi

import "github.com/gin-gonic/gin"

func NewHealthHandler(r *gin.Engine) {
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "healthy",
			"message": "Fabric API is running",
		})
	})
}
