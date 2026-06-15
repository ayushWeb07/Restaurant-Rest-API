package interfaces

import "github.com/gin-gonic/gin"

type RouterInterface interface {
	Register(router *gin.RouterGroup)
}
