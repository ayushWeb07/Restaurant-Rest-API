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

type OrderRouter struct {
	ServerConfig    *config.ServerConfig
	Logger          *zap.Logger
	OrderController interfaces.OrderControllerInterface
	MongoClient     *mongo.Client
}

func (orderRouter *OrderRouter) Register(router *gin.RouterGroup) {
	orderRouterGroup := router.Group("/orders")

	// create / -> create order
	orderRouterGroup.POST("", orderRouter.OrderController.CreateOrder())

	// get / -> get all orders
	orderRouterGroup.GET("", orderRouter.OrderController.GetAllOrders())

	// get /:id -> get order by id
	orderRouterGroup.GET("/:id", orderRouter.OrderController.GetOrderById())

	// put /:id -> update order by id
	orderRouterGroup.PUT("/:id", orderRouter.OrderController.UpdateOrderById())

	// delete /:id -> delete order by id
	orderRouterGroup.DELETE("/:id", orderRouter.OrderController.DeleteOrderById())
}

func NewOrderRouter(serverConfig *config.ServerConfig, logger *zap.Logger, mongoClient *mongo.Client) interfaces.RouterInterface {
	orderRepository := repositories.NewOrderRepository(serverConfig, logger, mongoClient)
	orderService := services.NewOrderService(serverConfig, logger, orderRepository, mongoClient)
	orderController := controllers.NewOrderController(serverConfig, logger, orderService, mongoClient)

	orderRouter := &OrderRouter{
		ServerConfig:    serverConfig,
		Logger:          logger,
		OrderController: orderController,
		MongoClient:     mongoClient,
	}

	return orderRouter
}
