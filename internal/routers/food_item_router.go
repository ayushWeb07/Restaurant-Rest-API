package routers

import (
	"github.com/ayushWeb07/Restaurant-Rest-API/internal/config"
	"github.com/ayushWeb07/Restaurant-Rest-API/internal/controllers"
	"github.com/ayushWeb07/Restaurant-Rest-API/internal/interfaces"
	"github.com/ayushWeb07/Restaurant-Rest-API/internal/repositories"
	"github.com/ayushWeb07/Restaurant-Rest-API/internal/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/zap"
)

type FoodItemRouter struct {
	ServerConfig       *config.ServerConfig
	Logger             *zap.Logger
	FoodItemController interfaces.FoodItemControllerInterface
	MongoClient        *mongo.Client
}

func (foodItemRouter *FoodItemRouter) Register(router *gin.RouterGroup) {
	foodItemRouterGroup := router.Group("/food-items")

	// create / -> create food item
	foodItemRouterGroup.POST("", foodItemRouter.FoodItemController.CreateFoodItem())

	// get / -> get all food items
	foodItemRouterGroup.GET("", foodItemRouter.FoodItemController.GetAllFoodItems())

	// get /:id -> get food item by id
	foodItemRouterGroup.GET("/:id", foodItemRouter.FoodItemController.GetFoodItemById())

	// put /:id -> update food item by id
	foodItemRouterGroup.PUT("/:id", foodItemRouter.FoodItemController.UpdateFoodItemById())

	// delete /:id -> delete food item by id
	foodItemRouterGroup.DELETE("/:id", foodItemRouter.FoodItemController.DeleteFoodItemById())
}

func NewFoodItemRouter(serverConfig *config.ServerConfig, logger *zap.Logger, mongoClient *mongo.Client) interfaces.RouterInterface {
	foodItemRepository := repositories.NewFoodItemRepository(serverConfig, logger, mongoClient)
	foodItemService := services.NewFoodItemService(serverConfig, logger, foodItemRepository, mongoClient)
	foodItemController := controllers.NewFoodItemController(serverConfig, logger, foodItemService, mongoClient)

	foodItemRouter := &FoodItemRouter{
		ServerConfig:       serverConfig,
		Logger:             logger,
		FoodItemController: foodItemController,
		MongoClient:        mongoClient,
	}

	return foodItemRouter
}
