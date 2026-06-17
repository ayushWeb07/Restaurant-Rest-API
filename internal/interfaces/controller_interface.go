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

type FoodItemControllerInterface interface {
	CreateFoodItem() gin.HandlerFunc
	GetAllFoodItems() gin.HandlerFunc
	GetFoodItemById() gin.HandlerFunc
	UpdateFoodItemById() gin.HandlerFunc
	DeleteFoodItemById() gin.HandlerFunc
}

type InvoiceControllerInterface interface {
	CreateInvoice() gin.HandlerFunc
	GetAllInvoices() gin.HandlerFunc
	GetInvoiceById() gin.HandlerFunc
	UpdateInvoiceById() gin.HandlerFunc
	DeleteInvoiceById() gin.HandlerFunc
}

type MenuControllerInterface interface {
	CreateMenu() gin.HandlerFunc
	GetAllMenus() gin.HandlerFunc
	GetMenuById() gin.HandlerFunc
	UpdateMenuById() gin.HandlerFunc
	DeleteMenuById() gin.HandlerFunc
}
