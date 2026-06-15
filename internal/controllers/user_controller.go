package controllers

import (
	"net/http"

	"github.com/ayushWeb07/Restaurant-Rest-API/internal/config"
	"github.com/ayushWeb07/Restaurant-Rest-API/internal/interfaces"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/zap"
)

type UserController struct {
	ServerConfig *config.ServerConfig
	Logger       *zap.Logger
	UserService  interfaces.UserServiceInterface
	MongoClient  *mongo.Client
}

func (userController *UserController) GetAllUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		userController.UserService.GetAllUsers()
		c.JSON(http.StatusOK, gin.H{"method": "GET"})
	}
}

func (userController *UserController) GetUserById() gin.HandlerFunc {
	return func(c *gin.Context) {
		userController.UserService.GetUserById()
		c.JSON(http.StatusOK, gin.H{"method": "GET"})
	}
}

func (userController *UserController) UpdateUserById() gin.HandlerFunc {
	return func(c *gin.Context) {
		userController.UserService.UpdateUserById()
		c.JSON(http.StatusOK, gin.H{"method": "GET"})
	}
}

func (userController *UserController) DeleteUserById() gin.HandlerFunc {
	return func(c *gin.Context) {
		userController.UserService.DeleteUserById()
		c.JSON(http.StatusOK, gin.H{"method": "GET"})
	}
}
