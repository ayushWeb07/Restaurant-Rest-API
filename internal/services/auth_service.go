package services

import (
	"github.com/ayushWeb07/Restaurant-Rest-API/internal/config"
	"github.com/ayushWeb07/Restaurant-Rest-API/internal/interfaces"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/zap"
)

type AuthService struct {
	ServerConfig *config.ServerConfig
	Logger       *zap.Logger
	MongoClient  *mongo.Client
}

func (authService *AuthService) RegisterUser() {
	authService.Logger.Info("Auth Service -> RegisterUser")
}

func (authService *AuthService) LoginUser() {
	authService.Logger.Info("Auth Service -> LoginUser")
}

func NewAuthService(serverConfig *config.ServerConfig, logger *zap.Logger, mongoClient *mongo.Client) interfaces.AuthServiceInterface {
	authService := &AuthService{
		ServerConfig: serverConfig,
		Logger:       logger,
		MongoClient:  mongoClient,
	}

	return authService
}
