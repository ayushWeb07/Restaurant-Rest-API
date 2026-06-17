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

type InvoiceRepositoryInterface interface {
	CreateInvoice()
	GetAllInvoices()
	GetInvoiceById()
	UpdateInvoiceById()
	DeleteInvoiceById()
}

type MenuRepositoryInterface interface {
	CreateMenu()
	GetAllMenus()
	GetMenuById()
	UpdateMenuById()
	DeleteMenuById()
}

type OrderRepositoryInterface interface {
	CreateOrder()
	GetAllOrders()
	GetOrderById()
	UpdateOrderById()
	DeleteOrderById()
}

type TableRepositoryInterface interface {
	CreateTable()
	GetAllTables()
	GetTableById()
	UpdateTableById()
	DeleteTableById()
}
