package services

import (
	"github.com/ayushWeb07/Restaurant-Rest-API/internal/config"
	"github.com/ayushWeb07/Restaurant-Rest-API/internal/interfaces"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/zap"
)

type OrderService struct {
	ServerConfig    *config.ServerConfig
	Logger          *zap.Logger
	OrderRepository interfaces.OrderRepositoryInterface
	MongoClient     *mongo.Client
}

func (orderService *OrderService) CreateOrder() {
	orderService.Logger.Info("Order Service -> CreateOrder")
	orderService.OrderRepository.CreateOrder()
}

func (orderService *OrderService) GetAllOrders() {
	orderService.Logger.Info("Order Service -> GetAllOrders")
	orderService.OrderRepository.GetAllOrders()
}

func (orderService *OrderService) GetOrderById() {
	orderService.Logger.Info("Order Service -> GetOrderById")
	orderService.OrderRepository.GetOrderById()
}

func (orderService *OrderService) UpdateOrderById() {
	orderService.Logger.Info("Order Service -> UpdateOrderById")
	orderService.OrderRepository.UpdateOrderById()
}

func (orderService *OrderService) DeleteOrderById() {
	orderService.Logger.Info("Order Service -> DeleteOrderById")
	orderService.OrderRepository.DeleteOrderById()
}

func NewOrderService(serverConfig *config.ServerConfig, logger *zap.Logger, orderRepository interfaces.OrderRepositoryInterface, mongoClient *mongo.Client) interfaces.OrderServiceInterface {
	orderService := &OrderService{
		ServerConfig:    serverConfig,
		Logger:          logger,
		OrderRepository: orderRepository,
		MongoClient:     mongoClient,
	}

	return orderService
}
