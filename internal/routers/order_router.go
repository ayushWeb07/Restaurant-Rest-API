package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SetupOrderRouter(router *gin.RouterGroup, logger *zap.Logger) {
	orderRouter := router.Group("/orders")

	// post / -> create a new order
	orderRouter.POST("", func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, gin.H{
			"status": "UP",
		})
	})

	// get / -> get all orders
	orderRouter.GET("", func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, gin.H{
			"status": "UP",
		})
	})

	// get /:id -> get order by id
	orderRouter.GET("/:id", func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, gin.H{
			"status": "UP",
		})
	})

	// put /:id -> update order by id
	orderRouter.PUT("/:id", func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, gin.H{
			"status": "UP",
		})
	})

	// delete /:id -> delete order by id
	orderRouter.DELETE("/:id", func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, gin.H{
			"status": "UP",
		})
	})
}
