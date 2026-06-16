package interfaces

type UserRepositoryInterface interface {
	GetAllUsers()
	GetUserById()
	UpdateUserById()
	DeleteUserById()
}

type FoodItemRepositoryInterface interface {
	CreateFoodItem()
	GetAllFoodItems()
	GetFoodItemById()
	UpdateFoodItemById()
	DeleteFoodItemById()
}
