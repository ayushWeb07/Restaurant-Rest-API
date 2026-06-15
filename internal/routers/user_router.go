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

type UserRouter struct {
	ServerConfig   *config.ServerConfig
	Logger         *zap.Logger
	UserController interfaces.UserControllerInterface
	MongoClient    *mongo.Client
}

func (userRouter *UserRouter) Register(router *gin.RouterGroup) {
	userRouterGroup := router.Group("/users")

	// get / -> get all users
	userRouterGroup.GET("", userRouter.UserController.GetAllUsers())

	// get /:id -> get user by id
	userRouterGroup.GET("/:id", userRouter.UserController.GetUserById())

	// put /:id -> update user by id
	userRouterGroup.PUT("/:id", userRouter.UserController.UpdateUserById())

	// delete /:id -> delete user by id
	userRouterGroup.DELETE("/:id", userRouter.UserController.DeleteUserById())
}

func NewUserRouter(serverConfig *config.ServerConfig, logger *zap.Logger, mongoClient *mongo.Client) interfaces.RouterInterface {
	userRepository := repositories.NewUserRepository(serverConfig, logger, mongoClient)
	userService := services.NewUserService(serverConfig, logger, userRepository, mongoClient)
	userController := controllers.NewUserController(serverConfig, logger, userService, mongoClient)

	userRouter := &UserRouter{
		ServerConfig:   serverConfig,
		Logger:         logger,
		UserController: userController,
		MongoClient:    mongoClient,
	}

	return userRouter
}
