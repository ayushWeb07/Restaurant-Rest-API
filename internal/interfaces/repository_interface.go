package interfaces

type UserRepositoryInterface interface {
	GetAllUsers()
	GetUserById()
	UpdateUserById()
	DeleteUserById()
}
