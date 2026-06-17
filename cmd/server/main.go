package server

import (
	"github.com/ayushWeb07/Restaurant-Rest-API/internal/config"
	"github.com/ayushWeb07/Restaurant-Rest-API/internal/routers"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/zap"
)

type App struct {
	ServerConfig *config.ServerConfig
	Logger       *zap.Logger
	Router       *gin.Engine
	MongoClient  *mongo.Client
}

func (app *App) RegisterRouters() {
	// create the version router
	v1Router := app.Router.Group("/api/v1")

	// setup user router
	userRouter := routers.NewUserRouter(app.ServerConfig, app.Logger, app.MongoClient)
	userRouter.Register(v1Router)

	// setup auth router
	authRouter := routers.NewAuthRouter(app.ServerConfig, app.Logger, app.MongoClient)
	authRouter.Register(v1Router)

	// setup food item router
	foodItemRouter := routers.NewFoodItemRouter(app.ServerConfig, app.Logger, app.MongoClient)
	foodItemRouter.Register(v1Router)

	// setup invoice router
	invoiceRouter := routers.NewInvoiceRouter(app.ServerConfig, app.Logger, app.MongoClient)
	invoiceRouter.Register(v1Router)

	// setup menu router
	menuRouter := routers.NewMenuRouter(app.ServerConfig, app.Logger, app.MongoClient)
	menuRouter.Register(v1Router)

	// setup order router
	routers.SetupOrderRouter(v1Router, app.Logger)

	// setup order item router
	routers.SetupOrderItemRouter(v1Router, app.Logger)

	// setup table router
	routers.SetupTableRouter(v1Router, app.Logger)
}

func (app *App) Run() {
	// start the server
	app.Logger.Info("Starting the server...",
		zap.String("port", app.ServerConfig.Port))

	app.Router.Run(":" + app.ServerConfig.Port)
}
