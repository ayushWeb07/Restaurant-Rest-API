package services

import (
	"github.com/ayushWeb07/Restaurant-Rest-API/internal/config"
	"github.com/ayushWeb07/Restaurant-Rest-API/internal/interfaces"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/zap"
)

type OrderItemService struct {
	ServerConfig        *config.ServerConfig
	Logger              *zap.Logger
	OrderItemRepository interfaces.OrderItemRepositoryInterface
	MongoClient         *mongo.Client
}

func (orderItemService *OrderItemService) CreateOrderItem() {
	orderItemService.Logger.Info("Order Item Service -> CreateOrderItem")
	orderItemService.OrderItemRepository.CreateOrderItem()
}

func (orderItemService *OrderItemService) GetAllOrderItemsByOrderId() {
	orderItemService.Logger.Info("Order Item Service -> GetAllOrderItemsByOrderId")
	orderItemService.OrderItemRepository.GetAllOrderItemsByOrderId()
}

func (orderItemService *OrderItemService) GetOrderItemById() {
	orderItemService.Logger.Info("Order Item Service -> GetOrderItemById")
	orderItemService.OrderItemRepository.GetOrderItemById()
}

func (orderItemService *OrderItemService) UpdateOrderItemById() {
	orderItemService.Logger.Info("Order Item Service -> UpdateOrderItemById")
	orderItemService.OrderItemRepository.UpdateOrderItemById()
}

func (orderItemService *OrderItemService) DeleteOrderItemById() {
	orderItemService.Logger.Info("Order Item Service -> DeleteOrderItemById")
	orderItemService.OrderItemRepository.DeleteOrderItemById()
}

func NewOrderItemService(serverConfig *config.ServerConfig, logger *zap.Logger, orderItemRepository interfaces.OrderItemRepositoryInterface, mongoClient *mongo.Client) interfaces.OrderItemServiceInterface {
	orderItemService := &OrderItemService{
		ServerConfig:        serverConfig,
		Logger:              logger,
		OrderItemRepository: orderItemRepository,
		MongoClient:         mongoClient,
	}

	return orderItemService
}
