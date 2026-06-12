package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SetupTableRouter(router *gin.RouterGroup, logger *zap.Logger) {
	tableRouter := router.Group("/tables")

	// post / -> create a new table
	tableRouter.POST("", func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, gin.H{
			"status": "UP",
		})
	})

	// get / -> get all tables
	tableRouter.GET("", func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, gin.H{
			"status": "UP",
		})
	})

	// get /:id -> get table by id
	tableRouter.GET("/:id", func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, gin.H{
			"status": "UP",
		})
	})

	// put /:id -> update table by id
	tableRouter.PUT("/:id", func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, gin.H{
			"status": "UP",
		})
	})

	// delete /:id -> delete table by id
	tableRouter.DELETE("/:id", func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, gin.H{
			"status": "UP",
		})
	})
}
