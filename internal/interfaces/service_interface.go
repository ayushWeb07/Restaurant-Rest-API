package interfaces

type UserServiceInterface interface {
	GetAllUsers()
	GetUserById()
	UpdateUserById()
	DeleteUserById()
}
