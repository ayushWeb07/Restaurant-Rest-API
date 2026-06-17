package repositories

import (
	"github.com/ayushWeb07/Restaurant-Rest-API/internal/config"
	"github.com/ayushWeb07/Restaurant-Rest-API/internal/interfaces"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/zap"
)

type TableRepository struct {
	ServerConfig *config.ServerConfig
	Logger       *zap.Logger
	MongoClient  *mongo.Client
}

func (tableRepository *TableRepository) CreateTable() {
	tableRepository.Logger.Info("Table Repo -> CreateTable")
}

func (tableRepository *TableRepository) GetAllTables() {
	tableRepository.Logger.Info("Table Repo -> GetAllTables")
}

func (tableRepository *TableRepository) GetTableById() {
	tableRepository.Logger.Info("Table Repo -> GetTableById")
}

func (tableRepository *TableRepository) UpdateTableById() {
	tableRepository.Logger.Info("Table Repo -> UpdateTableById")
}

func (tableRepository *TableRepository) DeleteTableById() {
	tableRepository.Logger.Info("Table Repo -> DeleteTableById")
}

func NewTableRepository(serverConfig *config.ServerConfig, logger *zap.Logger, mongoClient *mongo.Client) interfaces.TableRepositoryInterface {
	tableRepository := &TableRepository{
		ServerConfig: serverConfig,
		Logger:       logger,
		MongoClient:  mongoClient,
	}

	return tableRepository
}
