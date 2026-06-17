package controllers

import (
	"net/http"

	"github.com/ayushWeb07/Restaurant-Rest-API/internal/config"
	"github.com/ayushWeb07/Restaurant-Rest-API/internal/interfaces"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/zap"
)

type MenuController struct {
	ServerConfig *config.ServerConfig
	Logger       *zap.Logger
	MenuService  interfaces.MenuServiceInterface
	MongoClient  *mongo.Client
}

func (menuController *MenuController) CreateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		menuController.MenuService.CreateMenu()
		c.JSON(http.StatusOK, gin.H{"method": "GET"})
	}
}

func (menuController *MenuController) GetAllMenus() gin.HandlerFunc {
	return func(c *gin.Context) {
		menuController.MenuService.GetAllMenus()
		c.JSON(http.StatusOK, gin.H{"method": "GET"})
	}
}

func (menuController *MenuController) GetMenuById() gin.HandlerFunc {
	return func(c *gin.Context) {
		menuController.MenuService.GetMenuById()
		c.JSON(http.StatusOK, gin.H{"method": "GET"})
	}
}

func (menuController *MenuController) UpdateMenuById() gin.HandlerFunc {
	return func(c *gin.Context) {
		menuController.MenuService.UpdateMenuById()
		c.JSON(http.StatusOK, gin.H{"method": "GET"})
	}
}

func (menuController *MenuController) DeleteMenuById() gin.HandlerFunc {
	return func(c *gin.Context) {
		menuController.MenuService.DeleteMenuById()
		c.JSON(http.StatusOK, gin.H{"method": "GET"})
	}
}

func NewMenuController(serverConfig *config.ServerConfig, logger *zap.Logger, menuService interfaces.MenuServiceInterface, mongoClient *mongo.Client) interfaces.MenuControllerInterface {
	menuController := &MenuController{
		ServerConfig: serverConfig,
		Logger:       logger,
		MenuService:  menuService,
		MongoClient:  mongoClient,
	}

	return menuController
}
