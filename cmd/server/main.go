package server

import (
	"github.com/ayushWeb07/Restaurant-Rest-API/internal/config"
	"github.com/ayushWeb07/Restaurant-Rest-API/internal/routers"
	"github.com/gin-gonic/gin"
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
}

func (app *App) RegisterRouters() {
	// create the version router
	v1Router := app.Router.Group("/api/v1")

	// setup user router
	routers.SetupUserRouter(v1Router, app.Logger)

	// setup auth router
	routers.SetupAuthRouter(v1Router, app.Logger)

	// setup food item router
	routers.SetupFoodItemRouter(v1Router, app.Logger)

	// setup invoice router
	routers.SetupInvoiceRouter(v1Router, app.Logger)
}

func (app *App) Run() {
	// start the server
	app.Logger.Info("Starting the server...",
		zap.String("port", app.ServerConfig.Port))

	app.Router.Run(":" + app.ServerConfig.Port)
}
