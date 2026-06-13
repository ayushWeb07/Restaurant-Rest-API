package main

import (
	"context"
	"log"
	"time"

	"github.com/ayushWeb07/Restaurant-Rest-API/cmd/server"
	"github.com/ayushWeb07/Restaurant-Rest-API/internal/config"
	"github.com/ayushWeb07/Restaurant-Rest-API/internal/database"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
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

	// create the db instance
	mongoClient, err := database.GetDatabaseInstance(serverCfg)

	if err != nil {
		logger.Fatal("Something went wrong while getting the database instance",
			zap.String("error", err.Error()))
	} else {
		logger.Info("Successfully connected to the database")
	}

	defer func() {
		// create a delayed context to disconnect the mongo client
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := mongoClient.Disconnect(ctx); err != nil {
			logger.Fatal("Failed to disconnect from the mongo client",
				zap.String("error", err.Error()))
		}
	}()

	// create the app instance
	serverApp := &server.App{
		ServerConfig: serverCfg,
		Logger:       logger,
		Router:       router,
		MongoClient:  mongoClient,
	}

	// run the app
	serverApp.RegisterRouters()
	serverApp.Run()
}
