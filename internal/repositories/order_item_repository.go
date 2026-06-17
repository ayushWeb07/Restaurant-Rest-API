package repositories

import (
	"github.com/ayushWeb07/Restaurant-Rest-API/internal/config"
	"github.com/ayushWeb07/Restaurant-Rest-API/internal/interfaces"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/zap"
)

type OrderItemRepository struct {
	ServerConfig *config.ServerConfig
	Logger       *zap.Logger
	MongoClient  *mongo.Client
}

func (orderItemRepository *OrderItemRepository) CreateOrderItem() {
	orderItemRepository.Logger.Info("Order Item Repo -> CreateOrderItem")
}

func (orderItemRepository *OrderItemRepository) GetAllOrderItemsByOrderId() {
	orderItemRepository.Logger.Info("Order Item Repo -> GetAllOrderItemsByOrderId")
}

func (orderItemRepository *OrderItemRepository) GetOrderItemById() {
	orderItemRepository.Logger.Info("Order Item Repo -> GetOrderItemById")
}

func (orderItemRepository *OrderItemRepository) UpdateOrderItemById() {
	orderItemRepository.Logger.Info("Order Item Repo -> UpdateOrderItemById")
}

func (orderItemRepository *OrderItemRepository) DeleteOrderItemById() {
	orderItemRepository.Logger.Info("Order Item Repo -> DeleteOrderItemById")
}

func NewOrderItemRepository(serverConfig *config.ServerConfig, logger *zap.Logger, mongoClient *mongo.Client) interfaces.OrderItemRepositoryInterface {
	orderItemRepository := &OrderItemRepository{
		ServerConfig: serverConfig,
		Logger:       logger,
		MongoClient:  mongoClient,
	}

	return orderItemRepository
}
