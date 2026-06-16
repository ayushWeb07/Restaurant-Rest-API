package routers

import (
	"github.com/ayushWeb07/Restaurant-Rest-API/internal/config"
	"github.com/ayushWeb07/Restaurant-Rest-API/internal/controllers"
	"github.com/ayushWeb07/Restaurant-Rest-API/internal/interfaces"
	"github.com/ayushWeb07/Restaurant-Rest-API/internal/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/zap"
)

type AuthRouter struct {
	ServerConfig   *config.ServerConfig
	Logger         *zap.Logger
	AuthController interfaces.AuthControllerInterface
	MongoClient    *mongo.Client
}

func (authRouter *AuthRouter) Register(router *gin.RouterGroup) {
	authRouterGroup := router.Group("/auth")

	// post /register -> register user
	authRouterGroup.POST("/register", authRouter.AuthController.RegisterUser())

	// post /login -> login user
	authRouterGroup.POST("/login", authRouter.AuthController.LoginUser())
}

func NewAuthRouter(serverConfig *config.ServerConfig, logger *zap.Logger, mongoClient *mongo.Client) interfaces.RouterInterface {
	authService := services.NewAuthService(serverConfig, logger, mongoClient)
	authController := controllers.NewAuthController(serverConfig, logger, authService, mongoClient)

	authRouter := &AuthRouter{
		ServerConfig:   serverConfig,
		Logger:         logger,
		AuthController: authController,
		MongoClient:    mongoClient,
	}

	return authRouter
}
