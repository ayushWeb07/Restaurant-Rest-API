package repositories

import (
	"github.com/ayushWeb07/Restaurant-Rest-API/internal/config"
	"github.com/ayushWeb07/Restaurant-Rest-API/internal/interfaces"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/zap"
)

type MenuRepository struct {
	ServerConfig *config.ServerConfig
	Logger       *zap.Logger
	MongoClient  *mongo.Client
}

func (menuRepository *MenuRepository) CreateMenu() {
	menuRepository.Logger.Info("Menu Repo -> CreateMenu")
}

func (menuRepository *MenuRepository) GetAllMenus() {
	menuRepository.Logger.Info("Menu Repo -> GetAllMenus")
}

func (menuRepository *MenuRepository) GetMenuById() {
	menuRepository.Logger.Info("Menu Repo -> GetMenuById")
}

func (menuRepository *MenuRepository) UpdateMenuById() {
	menuRepository.Logger.Info("Menu Repo -> UpdateMenuById")
}

func (menuRepository *MenuRepository) DeleteMenuById() {
	menuRepository.Logger.Info("Menu Repo -> DeleteMenuById")
}

func NewMenuRepository(serverConfig *config.ServerConfig, logger *zap.Logger, mongoClient *mongo.Client) interfaces.MenuRepositoryInterface {
	menuRepository := &MenuRepository{
		ServerConfig: serverConfig,
		Logger:       logger,
		MongoClient:  mongoClient,
	}

	return menuRepository
}
