package routers

import (
	"github.com/ayushWeb07/Restaurant-Rest-API/internal/config"
	"github.com/ayushWeb07/Restaurant-Rest-API/internal/controllers"
	"github.com/ayushWeb07/Restaurant-Rest-API/internal/interfaces"
	"github.com/ayushWeb07/Restaurant-Rest-API/internal/repositories"
	"github.com/ayushWeb07/Restaurant-Rest-API/internal/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/zap"
)

type InvoiceRouter struct {
	ServerConfig      *config.ServerConfig
	Logger            *zap.Logger
	InvoiceController interfaces.InvoiceControllerInterface
	MongoClient       *mongo.Client
}

func (invoiceRouter *InvoiceRouter) Register(router *gin.RouterGroup) {
	invoiceRouterGroup := router.Group("/invoices")

	// create / -> create invoice
	invoiceRouterGroup.POST("", invoiceRouter.InvoiceController.CreateInvoice())

	// get / -> get all invoices
	invoiceRouterGroup.GET("", invoiceRouter.InvoiceController.GetAllInvoices())

	// get /:id -> get invoice by id
	invoiceRouterGroup.GET("/:id", invoiceRouter.InvoiceController.GetInvoiceById())

	// put /:id -> update invoice by id
	invoiceRouterGroup.PUT("/:id", invoiceRouter.InvoiceController.UpdateInvoiceById())

	// delete /:id -> delete invoice by id
	invoiceRouterGroup.DELETE("/:id", invoiceRouter.InvoiceController.DeleteInvoiceById())
}

func NewInvoiceRouter(serverConfig *config.ServerConfig, logger *zap.Logger, mongoClient *mongo.Client) interfaces.RouterInterface {
	invoiceRepository := repositories.NewInvoiceRepository(serverConfig, logger, mongoClient)
	invoiceService := services.NewInvoiceService(serverConfig, logger, invoiceRepository, mongoClient)
	invoiceController := controllers.NewInvoiceController(serverConfig, logger, invoiceService, mongoClient)

	invoiceRouter := &InvoiceRouter{
		ServerConfig:      serverConfig,
		Logger:            logger,
		InvoiceController: invoiceController,
		MongoClient:       mongoClient,
	}

	return invoiceRouter
}
