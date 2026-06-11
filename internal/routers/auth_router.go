package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SetupAuthRouter(router *gin.RouterGroup, logger *zap.Logger) {
	authRouter := router.Group("/auth")

	// post /signup -> signup or register
	authRouter.POST("/signup", func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, gin.H{
			"status": "UP",
		})
	})

	// post /signin -> signin or login
	authRouter.POST("/signin", func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, gin.H{
			"status": "UP",
		})
	})
}
