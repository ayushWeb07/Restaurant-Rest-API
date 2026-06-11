package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SetupUserRouter(router *gin.RouterGroup, logger *zap.Logger) {
	userRouter := router.Group("/users")

	// get / -> get all users
	userRouter.GET("", func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, gin.H{
			"status": "UP",
		})
	})

	// get /:id -> get user by id
	userRouter.GET("/:id", func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, gin.H{
			"status": "UP",
		})
	})

	// put /:id -> update user by id
	userRouter.PUT("/:id", func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, gin.H{
			"status": "UP",
		})
	})

	// delete /:id -> delete user by id
	userRouter.DELETE("/:id", func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, gin.H{
			"status": "UP",
		})
	})
}
