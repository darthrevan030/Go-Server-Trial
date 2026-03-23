package user

type Repository interface {
	CreateUser(User) (string, error)
	GetUserByID(string) (User, error)
	GetAllUsers() ([]User, error)
	UpdateUserAgeByID(string, int) (int, error)
	DeleteUserByID(string) (int, error)
	DeleteAllUsers() (int, error)
}
