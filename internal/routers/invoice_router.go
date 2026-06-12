package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SetupInvoiceRouter(router *gin.RouterGroup, logger *zap.Logger) {
	invoiceRouter := router.Group("/invoices")

	// post / -> create a new invoice
	invoiceRouter.POST("", func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, gin.H{
			"status": "UP",
		})
	})

	// get / -> get all invoices
	invoiceRouter.GET("", func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, gin.H{
			"status": "UP",
		})
	})

	// get /:id -> get invoice by id
	invoiceRouter.GET("/:id", func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, gin.H{
			"status": "UP",
		})
	})

	// put /:id -> update invoice by id
	invoiceRouter.PUT("/:id", func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, gin.H{
			"status": "UP",
		})
	})

	// delete /:id -> delete invoice by id
	invoiceRouter.DELETE("/:id", func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, gin.H{
			"status": "UP",
		})
	})
}
