package server

import (
	"log"
	"net/http"

	"github.com/ayushWeb07/Restaurant-Rest-API/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
	userV1Router := v1Router.Group("/users")

	userV1Router.GET("", func(c *gin.Context) {
		app.Logger.Info("Working fine...")

		c.JSON(http.StatusNotImplemented, gin.H{
			"status": "UP",
		})
	})

	// setup auth router

	// setup food item router
}

func (app *App) Run() {
	// validate the server config
	validate := validator.New()

	if validationErr := validate.Struct(app.ServerConfig); validationErr != nil {
		log.Fatal("Failed while validating the server config:" + validationErr.Error())
	}

	// start the server
	app.Logger.Info("Starting the server...",
		zap.String("port", app.ServerConfig.Port))

	app.Router.Run(":" + app.ServerConfig.Port)
}
