package controllers

import (
	"net/http"

	"github.com/ayushWeb07/Restaurant-Rest-API/internal/config"
	"github.com/ayushWeb07/Restaurant-Rest-API/internal/interfaces"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/zap"
)

type TableController struct {
	ServerConfig *config.ServerConfig
	Logger       *zap.Logger
	TableService interfaces.TableServiceInterface
	MongoClient  *mongo.Client
}

func (tableController *TableController) CreateTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		tableController.TableService.CreateTable()
		c.JSON(http.StatusOK, gin.H{"method": "GET"})
	}
}

func (tableController *TableController) GetAllTables() gin.HandlerFunc {
	return func(c *gin.Context) {
		tableController.TableService.GetAllTables()
		c.JSON(http.StatusOK, gin.H{"method": "GET"})
	}
}

func (tableController *TableController) GetTableById() gin.HandlerFunc {
	return func(c *gin.Context) {
		tableController.TableService.GetTableById()
		c.JSON(http.StatusOK, gin.H{"method": "GET"})
	}
}

func (tableController *TableController) UpdateTableById() gin.HandlerFunc {
	return func(c *gin.Context) {
		tableController.TableService.UpdateTableById()
		c.JSON(http.StatusOK, gin.H{"method": "GET"})
	}
}

func (tableController *TableController) DeleteTableById() gin.HandlerFunc {
	return func(c *gin.Context) {
		tableController.TableService.DeleteTableById()
		c.JSON(http.StatusOK, gin.H{"method": "GET"})
	}
}

func NewTableController(serverConfig *config.ServerConfig, logger *zap.Logger, tableService interfaces.TableServiceInterface, mongoClient *mongo.Client) interfaces.TableControllerInterface {
	tableController := &TableController{
		ServerConfig: serverConfig,
		Logger:       logger,
		TableService: tableService,
		MongoClient:  mongoClient,
	}

	return tableController
}
