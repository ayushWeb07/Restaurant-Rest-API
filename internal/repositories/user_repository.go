package repositories

import (
	"github.com/ayushWeb07/Restaurant-Rest-API/internal/config"
	"github.com/ayushWeb07/Restaurant-Rest-API/internal/interfaces"
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

func NewUserRepository(serverConfig *config.ServerConfig, logger *zap.Logger, mongoClient *mongo.Client) interfaces.UserRepositoryInterface {
	userRepository := &UserRepository{
		ServerConfig: serverConfig,
		Logger:       logger,
		MongoClient:  mongoClient,
	}

	return userRepository
}
