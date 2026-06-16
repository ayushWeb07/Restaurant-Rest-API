package services

import (
	"github.com/ayushWeb07/Restaurant-Rest-API/internal/config"
	"github.com/ayushWeb07/Restaurant-Rest-API/internal/interfaces"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/zap"
)

type InvoiceService struct {
	ServerConfig      *config.ServerConfig
	Logger            *zap.Logger
	InvoiceRepository interfaces.InvoiceRepositoryInterface
	MongoClient       *mongo.Client
}

func (invoiceService *InvoiceService) CreateInvoice() {
	invoiceService.Logger.Info("Invoice Service -> CreateInvoice")
	invoiceService.InvoiceRepository.CreateInvoice()
}

func (invoiceService *InvoiceService) GetAllInvoices() {
	invoiceService.Logger.Info("Invoice Service -> GetAllInvoices")
	invoiceService.InvoiceRepository.GetAllInvoices()
}

func (invoiceService *InvoiceService) GetInvoiceById() {
	invoiceService.Logger.Info("Invoice Service -> GetInvoiceById")
	invoiceService.InvoiceRepository.GetInvoiceById()
}

func (invoiceService *InvoiceService) UpdateInvoiceById() {
	invoiceService.Logger.Info("Invoice Service -> UpdateInvoiceById")
	invoiceService.InvoiceRepository.UpdateInvoiceById()
}

func (invoiceService *InvoiceService) DeleteInvoiceById() {
	invoiceService.Logger.Info("Invoice Service -> DeleteInvoiceById")
	invoiceService.InvoiceRepository.DeleteInvoiceById()
}

func NewInvoiceService(serverConfig *config.ServerConfig, logger *zap.Logger, invoiceRepository interfaces.InvoiceRepositoryInterface, mongoClient *mongo.Client) interfaces.InvoiceServiceInterface {
	invoiceService := &InvoiceService{
		ServerConfig:      serverConfig,
		Logger:            logger,
		InvoiceRepository: invoiceRepository,
		MongoClient:       mongoClient,
	}

	return invoiceService
}
