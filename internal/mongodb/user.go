package mongodb

import (
	"github.com/darthrevan030/go-server-trial/internal/model"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type MongoClient struct {
	Client mongo.Collection
}

func (c MongoClient) CreateUser(user model.User) (string, error) {
	result, err := c.Client.InsertOne(context.Background(), user)
	if err!= nil {
		return "", err
	}
	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (c MongoClient) GetUserByID(id string) (model.User, error) {
	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.User{}, fmt.Errorf("Invalid ID")
	}
	
	var user model.User
	filter := bson.D{{Key: "_id", Value: docID }}

	err = c.Client.FindOne(context.Background(), filter).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return model.User{}, fmt.Errorf("Record Does Not Exist")
	} else if err != nil {
		return model.User{}, err
	}

	return user, nil
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
