package repositories

import (
	"github.com/ayushWeb07/Restaurant-Rest-API/internal/config"
	"github.com/ayushWeb07/Restaurant-Rest-API/internal/interfaces"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/zap"
)

type InvoiceRepository struct {
	ServerConfig *config.ServerConfig
	Logger       *zap.Logger
	MongoClient  *mongo.Client
}

func (invoiceRepository *InvoiceRepository) CreateInvoice() {
	invoiceRepository.Logger.Info("Invoice Repo -> CreateInvoice")
}

func (invoiceRepository *InvoiceRepository) GetAllInvoices() {
	invoiceRepository.Logger.Info("Invoice Repo -> GetAllInvoices")
}

func (invoiceRepository *InvoiceRepository) GetInvoiceById() {
	invoiceRepository.Logger.Info("Invoice Repo -> GetInvoiceById")
}

func (invoiceRepository *InvoiceRepository) UpdateInvoiceById() {
	invoiceRepository.Logger.Info("Invoice Repo -> UpdateInvoiceById")
}

func (invoiceRepository *InvoiceRepository) DeleteInvoiceById() {
	invoiceRepository.Logger.Info("Invoice Repo -> DeleteInvoiceById")
}

func NewInvoiceRepository(serverConfig *config.ServerConfig, logger *zap.Logger, mongoClient *mongo.Client) interfaces.InvoiceRepositoryInterface {
	invoiceRepository := &InvoiceRepository{
		ServerConfig: serverConfig,
		Logger:       logger,
		MongoClient:  mongoClient,
	}

	return invoiceRepository
}
