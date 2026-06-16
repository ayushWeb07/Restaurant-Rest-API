package repositories

import (
	"github.com/ayushWeb07/Restaurant-Rest-API/internal/config"
	"github.com/ayushWeb07/Restaurant-Rest-API/internal/interfaces"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/zap"
)

type FoodItemRepository struct {
	ServerConfig *config.ServerConfig
	Logger       *zap.Logger
	MongoClient  *mongo.Client
}

func (foodItemRepository *FoodItemRepository) CreateFoodItem() {
	foodItemRepository.Logger.Info("Food Item Repo -> CreateFoodItem")
}

func (foodItemRepository *FoodItemRepository) GetAllFoodItems() {
	foodItemRepository.Logger.Info("Food Item Repo -> GetAllFoodItems")
}

func (foodItemRepository *FoodItemRepository) GetFoodItemById() {
	foodItemRepository.Logger.Info("Food Item Repo -> GetFoodItemById")
}

func (foodItemRepository *FoodItemRepository) UpdateFoodItemById() {
	foodItemRepository.Logger.Info("Food Item Repo -> UpdateFoodItemById")
}

func (foodItemRepository *FoodItemRepository) DeleteFoodItemById() {
	foodItemRepository.Logger.Info("Food Item Repo -> DeleteFoodItemById")
}

func NewFoodItemRepository(serverConfig *config.ServerConfig, logger *zap.Logger, mongoClient *mongo.Client) interfaces.FoodItemRepositoryInterface {
	foodItemRepository := &FoodItemRepository{
		ServerConfig: serverConfig,
		Logger:       logger,
		MongoClient:  mongoClient,
	}

	return foodItemRepository
}
