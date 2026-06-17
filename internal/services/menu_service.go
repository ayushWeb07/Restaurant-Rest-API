package services

import (
	"github.com/ayushWeb07/Restaurant-Rest-API/internal/config"
	"github.com/ayushWeb07/Restaurant-Rest-API/internal/interfaces"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/zap"
)

type MenuService struct {
	ServerConfig   *config.ServerConfig
	Logger         *zap.Logger
	MenuRepository interfaces.MenuRepositoryInterface
	MongoClient    *mongo.Client
}

func (menuService *MenuService) CreateMenu() {
	menuService.Logger.Info("Menu Service -> CreateMenu")
	menuService.MenuRepository.CreateMenu()
}

func (menuService *MenuService) GetAllMenus() {
	menuService.Logger.Info("Menu Service -> GetAllMenus")
	menuService.MenuRepository.GetAllMenus()
}

func (menuService *MenuService) GetMenuById() {
	menuService.Logger.Info("Menu Service -> GetMenuById")
	menuService.MenuRepository.GetMenuById()
}

func (menuService *MenuService) UpdateMenuById() {
	menuService.Logger.Info("Menu Service -> UpdateMenuById")
	menuService.MenuRepository.UpdateMenuById()
}

func (menuService *MenuService) DeleteMenuById() {
	menuService.Logger.Info("Menu Service -> DeleteMenuById")
	menuService.MenuRepository.DeleteMenuById()
}

func NewMenuService(serverConfig *config.ServerConfig, logger *zap.Logger, menuRepository interfaces.MenuRepositoryInterface, mongoClient *mongo.Client) interfaces.MenuServiceInterface {
	menuService := &MenuService{
		ServerConfig:   serverConfig,
		Logger:         logger,
		MenuRepository: menuRepository,
		MongoClient:    mongoClient,
	}

	return menuService
}
