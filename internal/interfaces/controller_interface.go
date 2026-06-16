package interfaces

import "github.com/gin-gonic/gin"

type UserControllerInterface interface {
	GetAllUsers() gin.HandlerFunc
	GetUserById() gin.HandlerFunc
	UpdateUserById() gin.HandlerFunc
	DeleteUserById() gin.HandlerFunc
}

type AuthControllerInterface interface {
	RegisterUser() gin.HandlerFunc
	LoginUser() gin.HandlerFunc
}
