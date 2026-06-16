package controllers

import (
	"net/http"

	"github.com/ayushWeb07/Restaurant-Rest-API/internal/config"
	"github.com/ayushWeb07/Restaurant-Rest-API/internal/interfaces"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/zap"
)

type InvoiceController struct {
	ServerConfig   *config.ServerConfig
	Logger         *zap.Logger
	InvoiceService interfaces.InvoiceServiceInterface
	MongoClient    *mongo.Client
}

func (invoiceController *InvoiceController) CreateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		invoiceController.InvoiceService.CreateInvoice()
		c.JSON(http.StatusOK, gin.H{"method": "GET"})
	}
}

func (invoiceController *InvoiceController) GetAllInvoices() gin.HandlerFunc {
	return func(c *gin.Context) {
		invoiceController.InvoiceService.GetAllInvoices()
		c.JSON(http.StatusOK, gin.H{"method": "GET"})
	}
}

func (invoiceController *InvoiceController) GetInvoiceById() gin.HandlerFunc {
	return func(c *gin.Context) {
		invoiceController.InvoiceService.GetInvoiceById()
		c.JSON(http.StatusOK, gin.H{"method": "GET"})
	}
}

func (invoiceController *InvoiceController) UpdateInvoiceById() gin.HandlerFunc {
	return func(c *gin.Context) {
		invoiceController.InvoiceService.UpdateInvoiceById()
		c.JSON(http.StatusOK, gin.H{"method": "GET"})
	}
}

func (invoiceController *InvoiceController) DeleteInvoiceById() gin.HandlerFunc {
	return func(c *gin.Context) {
		invoiceController.InvoiceService.DeleteInvoiceById()
		c.JSON(http.StatusOK, gin.H{"method": "GET"})
	}
}

func NewInvoiceController(serverConfig *config.ServerConfig, logger *zap.Logger, invoiceService interfaces.InvoiceServiceInterface, mongoClient *mongo.Client) interfaces.InvoiceControllerInterface {
	invoiceController := &InvoiceController{
		ServerConfig:   serverConfig,
		Logger:         logger,
		InvoiceService: invoiceService,
		MongoClient:    mongoClient,
	}

	return invoiceController
}
