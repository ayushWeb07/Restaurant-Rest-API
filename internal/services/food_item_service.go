package services

import (
	"github.com/ayushWeb07/Restaurant-Rest-API/internal/config"
	"github.com/ayushWeb07/Restaurant-Rest-API/internal/interfaces"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/zap"
)

type FoodItemService struct {
	ServerConfig       *config.ServerConfig
	Logger             *zap.Logger
	FoodItemRepository interfaces.FoodItemRepositoryInterface
	MongoClient        *mongo.Client
}

func (foodItemService *FoodItemService) CreateFoodItem() {
	foodItemService.Logger.Info("Food Item Service -> CreateFoodItem")
	foodItemService.FoodItemRepository.CreateFoodItem()
}

func (foodItemService *FoodItemService) GetAllFoodItems() {
	foodItemService.Logger.Info("Food Item Service -> GetAllFoodItems")
	foodItemService.FoodItemRepository.GetAllFoodItems()
}

func (foodItemService *FoodItemService) GetFoodItemById() {
	foodItemService.Logger.Info("Food Item Service -> GetFoodItemById")
	foodItemService.FoodItemRepository.GetFoodItemById()
}

func (foodItemService *FoodItemService) UpdateFoodItemById() {
	foodItemService.Logger.Info("Food Item Service -> UpdateFoodItemById")
	foodItemService.FoodItemRepository.UpdateFoodItemById()
}

func (foodItemService *FoodItemService) DeleteFoodItemById() {
	foodItemService.Logger.Info("Food Item Service -> DeleteFoodItemById")
	foodItemService.FoodItemRepository.DeleteFoodItemById()
}

func NewFoodItemService(serverConfig *config.ServerConfig, logger *zap.Logger, foodItemRepository interfaces.FoodItemRepositoryInterface, mongoClient *mongo.Client) interfaces.FoodItemServiceInterface {
	foodItemService := &FoodItemService{
		ServerConfig:       serverConfig,
		Logger:             logger,
		FoodItemRepository: foodItemRepository,
		MongoClient:        mongoClient,
	}

	return foodItemService
}
