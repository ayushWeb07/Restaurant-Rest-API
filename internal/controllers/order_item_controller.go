package controllers

import (
	"net/http"

	"github.com/ayushWeb07/Restaurant-Rest-API/internal/config"
	"github.com/ayushWeb07/Restaurant-Rest-API/internal/interfaces"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/zap"
)

type OrderItemController struct {
	ServerConfig     *config.ServerConfig
	Logger           *zap.Logger
	OrderItemService interfaces.OrderItemServiceInterface
	MongoClient      *mongo.Client
}

func (orderItemController *OrderItemController) CreateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		orderItemController.OrderItemService.CreateOrderItem()
		c.JSON(http.StatusOK, gin.H{"method": "GET"})
	}
}

func (orderItemController *OrderItemController) GetAllOrderItemsByOrderId() gin.HandlerFunc {
	return func(c *gin.Context) {
		orderItemController.OrderItemService.GetAllOrderItemsByOrderId()
		c.JSON(http.StatusOK, gin.H{"method": "GET"})
	}
}

func (orderItemController *OrderItemController) GetOrderItemById() gin.HandlerFunc {
	return func(c *gin.Context) {
		orderItemController.OrderItemService.GetOrderItemById()
		c.JSON(http.StatusOK, gin.H{"method": "GET"})
	}
}

func (orderItemController *OrderItemController) UpdateOrderItemById() gin.HandlerFunc {
	return func(c *gin.Context) {
		orderItemController.OrderItemService.UpdateOrderItemById()
		c.JSON(http.StatusOK, gin.H{"method": "GET"})
	}
}

func (orderItemController *OrderItemController) DeleteOrderItemById() gin.HandlerFunc {
	return func(c *gin.Context) {
		orderItemController.OrderItemService.DeleteOrderItemById()
		c.JSON(http.StatusOK, gin.H{"method": "GET"})
	}
}

func NewOrderItemController(serverConfig *config.ServerConfig, logger *zap.Logger, orderItemService interfaces.OrderItemServiceInterface, mongoClient *mongo.Client) interfaces.OrderItemControllerInterface {
	orderItemController := &OrderItemController{
		ServerConfig:     serverConfig,
		Logger:           logger,
		OrderItemService: orderItemService,
		MongoClient:      mongoClient,
	}

	return orderItemController
}
