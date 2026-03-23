package mongodb

import (
	"github.com/darthrevan030/go-server-trial/internal/model"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type MongoClient struct {
	Client mongo.Collection
}

func (c MongoClient) CreateUser(user model.User) (string, error) {
	return "", nil
}

func (c MongoClient) GetUserByID(user string) (model.User, error) {
	return model.User{}, nil
}

func (c MongoClient) GetAllUsers() ([]model.User, error) {
	return []model.User{}, nil
}

func (c MongoClient) UpdateUserAgeByID(id string, age int) (int, error) {
	return 0, nil
}

func (c MongoClient) DeleteUserByID(id string) (int, error) {
	return 0, nil
}

func (c MongoClient) DeleteAllUsers() (int, error) {
	return 0, nil
}
