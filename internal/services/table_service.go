package services

import (
	"github.com/ayushWeb07/Restaurant-Rest-API/internal/config"
	"github.com/ayushWeb07/Restaurant-Rest-API/internal/interfaces"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/zap"
)

type TableService struct {
	ServerConfig    *config.ServerConfig
	Logger          *zap.Logger
	TableRepository interfaces.TableRepositoryInterface
	MongoClient     *mongo.Client
}

func (tableService *TableService) CreateTable() {
	tableService.Logger.Info("Table Service -> CreateTable")
	tableService.TableRepository.CreateTable()
}

func (tableService *TableService) GetAllTables() {
	tableService.Logger.Info("Table Service -> GetAllTables")
	tableService.TableRepository.GetAllTables()
}

func (tableService *TableService) GetTableById() {
	tableService.Logger.Info("Table Service -> GetTableById")
	tableService.TableRepository.GetTableById()
}

func (tableService *TableService) UpdateTableById() {
	tableService.Logger.Info("Table Service -> UpdateTableById")
	tableService.TableRepository.UpdateTableById()
}

func (tableService *TableService) DeleteTableById() {
	tableService.Logger.Info("Table Service -> DeleteTableById")
	tableService.TableRepository.DeleteTableById()
}

func NewTableService(serverConfig *config.ServerConfig, logger *zap.Logger, tableRepository interfaces.TableRepositoryInterface, mongoClient *mongo.Client) interfaces.TableServiceInterface {
	tableService := &TableService{
		ServerConfig:    serverConfig,
		Logger:          logger,
		TableRepository: tableRepository,
		MongoClient:     mongoClient,
	}

	return tableService
}
