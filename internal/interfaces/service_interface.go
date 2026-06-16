package interfaces

type UserServiceInterface interface {
	GetAllUsers()
	GetUserById()
	UpdateUserById()
	DeleteUserById()
}

type AuthServiceInterface interface {
	RegisterUser()
	LoginUser()
}

type FoodItemServiceInterface interface {
	CreateFoodItem()
	GetAllFoodItems()
	GetFoodItemById()
	UpdateFoodItemById()
	DeleteFoodItemById()
}
