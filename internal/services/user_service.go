package services

import (
	"github.com/ayushWeb07/Restaurant-Rest-API/internal/config"
	"github.com/ayushWeb07/Restaurant-Rest-API/internal/interfaces"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/zap"
)

type UserService struct {
	ServerConfig   *config.ServerConfig
	Logger         *zap.Logger
	UserRepository interfaces.UserRepositoryInterface
	MongoClient    *mongo.Client
}

func (userService *UserService) GetAllUsers() {
	userService.Logger.Info("User Service -> GetAllUsers")
	userService.UserRepository.GetAllUsers()
}

func (userService *UserService) GetUserById() {
	userService.Logger.Info("User Service -> GetAllUsers")
	userService.UserRepository.GetUserById()
}

func (userService *UserService) UpdateUserById() {
	userService.Logger.Info("User Service -> GetAllUsers")
	userService.UserRepository.UpdateUserById()
}

func (userService *UserService) DeleteUserById() {
	userService.Logger.Info("User Service -> GetAllUsers")
	userService.UserRepository.DeleteUserById()
}
