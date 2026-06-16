package controllers

import (
	"net/http"

	"github.com/ayushWeb07/Restaurant-Rest-API/internal/config"
	"github.com/ayushWeb07/Restaurant-Rest-API/internal/interfaces"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/zap"
)

type FoodItemController struct {
	ServerConfig    *config.ServerConfig
	Logger          *zap.Logger
	FoodItemService interfaces.FoodItemServiceInterface
	MongoClient     *mongo.Client
}

func (foodItemController *FoodItemController) CreateFoodItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		foodItemController.FoodItemService.CreateFoodItem()
		c.JSON(http.StatusOK, gin.H{"method": "GET"})
	}
}

func (foodItemController *FoodItemController) GetAllFoodItems() gin.HandlerFunc {
	return func(c *gin.Context) {
		foodItemController.FoodItemService.GetAllFoodItems()
		c.JSON(http.StatusOK, gin.H{"method": "GET"})
	}
}

func (foodItemController *FoodItemController) GetFoodItemById() gin.HandlerFunc {
	return func(c *gin.Context) {
		foodItemController.FoodItemService.GetFoodItemById()
		c.JSON(http.StatusOK, gin.H{"method": "GET"})
	}
}

func (foodItemController *FoodItemController) UpdateFoodItemById() gin.HandlerFunc {
	return func(c *gin.Context) {
		foodItemController.FoodItemService.UpdateFoodItemById()
		c.JSON(http.StatusOK, gin.H{"method": "GET"})
	}
}

func (foodItemController *FoodItemController) DeleteFoodItemById() gin.HandlerFunc {
	return func(c *gin.Context) {
		foodItemController.FoodItemService.DeleteFoodItemById()
		c.JSON(http.StatusOK, gin.H{"method": "GET"})
	}
}

func NewFoodItemController(serverConfig *config.ServerConfig, logger *zap.Logger, foodItemService interfaces.FoodItemServiceInterface, mongoClient *mongo.Client) interfaces.FoodItemControllerInterface {
	foodItemController := &FoodItemController{
		ServerConfig:    serverConfig,
		Logger:          logger,
		FoodItemService: foodItemService,
		MongoClient:     mongoClient,
	}

	return foodItemController
}
