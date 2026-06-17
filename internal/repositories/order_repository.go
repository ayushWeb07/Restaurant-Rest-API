package repositories

import (
	"github.com/ayushWeb07/Restaurant-Rest-API/internal/config"
	"github.com/ayushWeb07/Restaurant-Rest-API/internal/interfaces"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/zap"
)

type OrderRepository struct {
	ServerConfig *config.ServerConfig
	Logger       *zap.Logger
	MongoClient  *mongo.Client
}

func (orderRepository *OrderRepository) CreateOrder() {
	orderRepository.Logger.Info("Order Repo -> CreateOrder")
}

func (orderRepository *OrderRepository) GetAllOrders() {
	orderRepository.Logger.Info("Order Repo -> GetAllOrders")
}

func (orderRepository *OrderRepository) GetOrderById() {
	orderRepository.Logger.Info("Order Repo -> GetOrderById")
}

func (orderRepository *OrderRepository) UpdateOrderById() {
	orderRepository.Logger.Info("Order Repo -> UpdateOrderById")
}

func (orderRepository *OrderRepository) DeleteOrderById() {
	orderRepository.Logger.Info("Order Repo -> DeleteOrderById")
}

func NewOrderRepository(serverConfig *config.ServerConfig, logger *zap.Logger, mongoClient *mongo.Client) interfaces.OrderRepositoryInterface {
	orderRepository := &OrderRepository{
		ServerConfig: serverConfig,
		Logger:       logger,
		MongoClient:  mongoClient,
	}

	return orderRepository
}
