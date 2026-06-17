package controllers

import (
	"net/http"

	"github.com/ayushWeb07/Restaurant-Rest-API/internal/config"
	"github.com/ayushWeb07/Restaurant-Rest-API/internal/interfaces"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/zap"
)

type OrderController struct {
	ServerConfig *config.ServerConfig
	Logger       *zap.Logger
	OrderService interfaces.OrderServiceInterface
	MongoClient  *mongo.Client
}

func (orderController *OrderController) CreateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		orderController.OrderService.CreateOrder()
		c.JSON(http.StatusOK, gin.H{"method": "GET"})
	}
}

func (orderController *OrderController) GetAllOrders() gin.HandlerFunc {
	return func(c *gin.Context) {
		orderController.OrderService.GetAllOrders()
		c.JSON(http.StatusOK, gin.H{"method": "GET"})
	}
}

func (orderController *OrderController) GetOrderById() gin.HandlerFunc {
	return func(c *gin.Context) {
		orderController.OrderService.GetOrderById()
		c.JSON(http.StatusOK, gin.H{"method": "GET"})
	}
}

func (orderController *OrderController) UpdateOrderById() gin.HandlerFunc {
	return func(c *gin.Context) {
		orderController.OrderService.UpdateOrderById()
		c.JSON(http.StatusOK, gin.H{"method": "GET"})
	}
}

func (orderController *OrderController) DeleteOrderById() gin.HandlerFunc {
	return func(c *gin.Context) {
		orderController.OrderService.DeleteOrderById()
		c.JSON(http.StatusOK, gin.H{"method": "GET"})
	}
}

func NewOrderController(serverConfig *config.ServerConfig, logger *zap.Logger, orderService interfaces.OrderServiceInterface, mongoClient *mongo.Client) interfaces.OrderControllerInterface {
	orderController := &OrderController{
		ServerConfig: serverConfig,
		Logger:       logger,
		OrderService: orderService,
		MongoClient:  mongoClient,
	}

	return orderController
}
