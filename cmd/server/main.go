package server

import (
	"github.com/ayushWeb07/Restaurant-Rest-API/internal/config"
	"github.com/ayushWeb07/Restaurant-Rest-API/internal/controllers"
	"github.com/ayushWeb07/Restaurant-Rest-API/internal/repositories"
	"github.com/ayushWeb07/Restaurant-Rest-API/internal/routers"
	"github.com/ayushWeb07/Restaurant-Rest-API/internal/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/zap"
)

type AppInterface interface {
	Run()
	RegisterRouters()
}

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
	userRepository := &repositories.UserRepository{
		ServerConfig: app.ServerConfig,
		Logger:       app.Logger,
		MongoClient:  app.MongoClient,
	}

	userService := &services.UserService{
		ServerConfig:   app.ServerConfig,
		Logger:         app.Logger,
		UserRepository: userRepository,
		MongoClient:    app.MongoClient,
	}

	userController := &controllers.UserController{
		ServerConfig: app.ServerConfig,
		Logger:       app.Logger,
		UserService:  userService,
		MongoClient:  app.MongoClient,
	}

	userRouter := &routers.UserRouter{
		ServerConfig:   app.ServerConfig,
		Logger:         app.Logger,
		UserController: userController,
		MongoClient:    app.MongoClient,
	}

	userRouter.Register(v1Router)

	// setup auth router
	routers.SetupAuthRouter(v1Router, app.Logger)

	// setup food item router
	routers.SetupFoodItemRouter(v1Router, app.Logger)

	// setup invoice router
	routers.SetupInvoiceRouter(v1Router, app.Logger)

	// setup menu router
	routers.SetupMenuRouter(v1Router, app.Logger)

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
