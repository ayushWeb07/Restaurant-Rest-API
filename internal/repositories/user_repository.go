package repositories

import (
	"github.com/ayushWeb07/Restaurant-Rest-API/internal/config"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/zap"
)

type UserRepository struct {
	ServerConfig *config.ServerConfig
	Logger       *zap.Logger
	MongoClient  *mongo.Client
}

func (userRepository *UserRepository) GetAllUsers() {
	userRepository.Logger.Info("User Repo -> GetAllUsers")
}

func (userRepository *UserRepository) GetUserById() {
	userRepository.Logger.Info("User Repo -> GetUserById")
}

func (userRepository *UserRepository) UpdateUserById() {
	userRepository.Logger.Info("User Repo -> UpdateUserById")
}

func (userRepository *UserRepository) DeleteUserById() {
	userRepository.Logger.Info("User Repo -> DeleteUserById")
}
