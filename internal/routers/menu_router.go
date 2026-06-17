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

type MenuRouter struct {
	ServerConfig   *config.ServerConfig
	Logger         *zap.Logger
	MenuController interfaces.MenuControllerInterface
	MongoClient    *mongo.Client
}

func (menuRouter *MenuRouter) Register(router *gin.RouterGroup) {
	menuRouterGroup := router.Group("/menus")

	// create / -> create menu
	menuRouterGroup.POST("", menuRouter.MenuController.CreateMenu())

	// get / -> get all menus
	menuRouterGroup.GET("", menuRouter.MenuController.GetAllMenus())

	// get /:id -> get menu by id
	menuRouterGroup.GET("/:id", menuRouter.MenuController.GetMenuById())

	// put /:id -> update menu by id
	menuRouterGroup.PUT("/:id", menuRouter.MenuController.UpdateMenuById())

	// delete /:id -> delete menu by id
	menuRouterGroup.DELETE("/:id", menuRouter.MenuController.DeleteMenuById())
}

func NewMenuRouter(serverConfig *config.ServerConfig, logger *zap.Logger, mongoClient *mongo.Client) interfaces.RouterInterface {
	menuRepository := repositories.NewMenuRepository(serverConfig, logger, mongoClient)
	menuService := services.NewMenuService(serverConfig, logger, menuRepository, mongoClient)
	menuController := controllers.NewMenuController(serverConfig, logger, menuService, mongoClient)

	menuRouter := &MenuRouter{
		ServerConfig:   serverConfig,
		Logger:         logger,
		MenuController: menuController,
		MongoClient:    mongoClient,
	}

	return menuRouter
}
