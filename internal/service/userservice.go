package service

import (
	"net/http"

	"github.com/darthrevan030/go-server-trial/internal/interfaces"
)

type UserService struct {
	DBClient interfaces.UserInterface
}

func (service UserService) CreateUser(writer http.ResponseWriter, request *http.Request) {

}

func (service UserService) GetUserByID(writer http.ResponseWriter, request *http.Request) {

}

func (service UserService) GetAllUsers(writer http.ResponseWriter, request *http.Request) {

}

func (service UserService) UpdateUserAgeByID(writer http.ResponseWriter, request *http.Request) {

}

func (service UserService) DeleteUserByID(writer http.ResponseWriter, request *http.Request) {

}

func (service UserService) DeleteAllUsers(writer http.ResponseWriter, request *http.Request) {
	
}