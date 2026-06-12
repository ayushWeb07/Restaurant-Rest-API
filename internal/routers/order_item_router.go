package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SetupOrderItemRouter(router *gin.RouterGroup, logger *zap.Logger) {
	orderItemRouter := router.Group("/order-items")

	// post / -> create a new order item
	orderItemRouter.POST("", func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, gin.H{
			"status": "UP",
		})
	})

	// get /order/:orderId -> get all order items of an order by id
	orderItemRouter.GET("/order/:orderId", func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, gin.H{
			"status": "UP",
		})
	})

	// get /:id -> get order item by id
	orderItemRouter.GET("/:id", func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, gin.H{
			"status": "UP",
		})
	})

	// put /:id -> update order item by id
	orderItemRouter.PUT("/:id", func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, gin.H{
			"status": "UP",
		})
	})

	// delete /:id -> delete order item by id
	orderItemRouter.DELETE("/:id", func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, gin.H{
			"status": "UP",
		})
	})
}
