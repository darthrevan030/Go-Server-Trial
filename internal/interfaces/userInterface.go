package interfaces

import "github.com/darthrevan030/go-server-trial/internal/model"

type UserInterface interface {
	CreateUser(model.User) (string, error)
	GetUserByID(string) (model.User, error)
	GetAllUsers() ([]model.User, error)
	UpdateUserAgeByID(string, int) (int, error)
	DeleteUserByID(string, int) (int, error)
	DeleteAllUsers() (int, error)
	
}