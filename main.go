package main

import (
	"log"

	"github.com/ayushWeb07/Restaurant-Rest-API/cmd/server"
	"github.com/ayushWeb07/Restaurant-Rest-API/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	// create the server config instance
	serverCfg, err := config.LoadServerConfig()

	if err != nil {
		log.Fatal(err)
	}

	// validate the server config
	validate := validator.New()

	if validationErr := validate.Struct(serverCfg); validationErr != nil {
		log.Fatal("Failed while validating the server config: " + validationErr.Error())
	}

	// create the logger instance
	logger := config.GetLogger(serverCfg.AppEnv)

	// create the router instance
	router := gin.Default()
	router.SetTrustedProxies(nil)

	// create the app instance
	serverApp := &server.App{
		ServerConfig: serverCfg,
		Logger:       logger,
		Router:       router,
	}

	// run the app
	serverApp.RegisterRouters()
	serverApp.Run()
}
