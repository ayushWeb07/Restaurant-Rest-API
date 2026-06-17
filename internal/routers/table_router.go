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

type TableRouter struct {
	ServerConfig    *config.ServerConfig
	Logger          *zap.Logger
	TableController interfaces.TableControllerInterface
	MongoClient     *mongo.Client
}

func (tableRouter *TableRouter) Register(router *gin.RouterGroup) {
	tableRouterGroup := router.Group("/tables")

	// create / -> create table
	tableRouterGroup.POST("", tableRouter.TableController.CreateTable())

	// get / -> get all tables
	tableRouterGroup.GET("", tableRouter.TableController.GetAllTables())

	// get /:id -> get table by id
	tableRouterGroup.GET("/:id", tableRouter.TableController.GetTableById())

	// put /:id -> update table by id
	tableRouterGroup.PUT("/:id", tableRouter.TableController.UpdateTableById())

	// delete /:id -> delete table by id
	tableRouterGroup.DELETE("/:id", tableRouter.TableController.DeleteTableById())
}

func NewTableRouter(serverConfig *config.ServerConfig, logger *zap.Logger, mongoClient *mongo.Client) interfaces.RouterInterface {
	tableRepository := repositories.NewTableRepository(serverConfig, logger, mongoClient)
	tableService := services.NewTableService(serverConfig, logger, tableRepository, mongoClient)
	tableController := controllers.NewTableController(serverConfig, logger, tableService, mongoClient)

	tableRouter := &TableRouter{
		ServerConfig:    serverConfig,
		Logger:          logger,
		TableController: tableController,
		MongoClient:     mongoClient,
	}

	return tableRouter
}
