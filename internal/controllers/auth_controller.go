package controllers

import (
	"net/http"

	"github.com/ayushWeb07/Restaurant-Rest-API/internal/config"
	"github.com/ayushWeb07/Restaurant-Rest-API/internal/interfaces"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/zap"
)

type AuthController struct {
	ServerConfig *config.ServerConfig
	Logger       *zap.Logger
	AuthService  interfaces.AuthServiceInterface
	MongoClient  *mongo.Client
}

func (authController *AuthController) RegisterUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		authController.AuthService.RegisterUser()
		c.JSON(http.StatusOK, gin.H{"method": "GET"})
	}
}

func (authController *AuthController) LoginUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		authController.AuthService.LoginUser()
		c.JSON(http.StatusOK, gin.H{"method": "GET"})
	}
}

func NewAuthController(serverConfig *config.ServerConfig, logger *zap.Logger, authService interfaces.AuthServiceInterface, mongoClient *mongo.Client) interfaces.AuthControllerInterface {
	authController := &AuthController{
		ServerConfig: serverConfig,
		Logger:       logger,
		AuthService:  authService,
		MongoClient:  mongoClient,
	}

	return authController
}
