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

type OrderItemRouter struct {
	ServerConfig        *config.ServerConfig
	Logger              *zap.Logger
	OrderItemController interfaces.OrderItemControllerInterface
	MongoClient         *mongo.Client
}

func (orderItemRouter *OrderItemRouter) Register(router *gin.RouterGroup) {
	orderItemRouterGroup := router.Group("/order-items")

	// create / -> create order item
	orderItemRouterGroup.POST("", orderItemRouter.OrderItemController.CreateOrderItem())

	// get / -> get all order items by order id
	orderItemRouterGroup.GET("/order/:orderId", orderItemRouter.OrderItemController.GetAllOrderItemsByOrderId())

	// get /:id -> get order item by id
	orderItemRouterGroup.GET("/:id", orderItemRouter.OrderItemController.GetOrderItemById())

	// put /:id -> update order item by id
	orderItemRouterGroup.PUT("/:id", orderItemRouter.OrderItemController.UpdateOrderItemById())

	// delete /:id -> delete order item by id
	orderItemRouterGroup.DELETE("/:id", orderItemRouter.OrderItemController.DeleteOrderItemById())
}

func NewOrderItemRouter(serverConfig *config.ServerConfig, logger *zap.Logger, mongoClient *mongo.Client) interfaces.RouterInterface {
	orderItemRepository := repositories.NewOrderItemRepository(serverConfig, logger, mongoClient)
	orderItemService := services.NewOrderItemService(serverConfig, logger, orderItemRepository, mongoClient)
	orderItemController := controllers.NewOrderItemController(serverConfig, logger, orderItemService, mongoClient)

	orderItemRouter := &OrderItemRouter{
		ServerConfig:        serverConfig,
		Logger:              logger,
		OrderItemController: orderItemController,
		MongoClient:         mongoClient,
	}

	return orderItemRouter
}
