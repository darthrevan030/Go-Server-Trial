package user

import (
	"context"
	"fmt"
	"log/slog"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type MongoRepository struct {
	collection *mongo.Collection
}

func NewRepository(db *mongo.Database, collectionName string) Repository {
	return &MongoRepository{
		collection: db.Collection(collectionName),
	}
}

func (c MongoRepository) CreateUser(user User) (string, error) {
	result, err := c.collection.InsertOne(context.Background(), user)
	if err != nil {
		return "", err
	}
	return result.InsertedID.(bson.ObjectID).Hex(), nil
}

func (c MongoRepository) GetUserByID(id string) (User, error) {
	docID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return User{}, fmt.Errorf("Invalid ID")
	}

	var user User
	filter := bson.D{{Key: "_id", Value: docID}}

	err = c.collection.FindOne(context.Background(), filter).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return User{}, fmt.Errorf("Record Does Not Exist")
	} else if err != nil {
		return User{}, err
	}

	return user, nil
}

func (c MongoRepository) GetAllUsers() ([]User, error) {
	filter := bson.D{}

	cur, err := c.collection.Find(context.Background(), filter)
	if err != nil {
		return []User{}, err
	}
	defer cur.Close(context.Background())

	var users []User

	for cur.Next(context.Background()) {
		var user User

		err := cur.Decode(&user)

		if err != nil {
			slog.Error("Error while decoding users", slog.String("error", err.Error()))
			continue
		}

		users = append(users, user)
	}

	return users, nil
}

func (c MongoRepository) UpdateUserAgeByID(id string, age int) (int, error) {
	docID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return 0, fmt.Errorf("Invalid ID")
	}

	filter := bson.D{{Key: "_id", Value: docID}}
	updateStatement := bson.D{{Key: "$set", Value: bson.D{{Key: "age", Value: age}}}}

	result, err := c.collection.UpdateOne(context.Background(), filter, updateStatement)
	if err != nil {
		return 0, err
	}

	return int(result.ModifiedCount), nil
}

func (c MongoRepository) DeleteUserByID(id string) (int, error) {
	docID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return 0, fmt.Errorf("Invalid ID")
	}

	filter := bson.D{{Key: "_id", Value: docID}}

	result, err := c.collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return 0, err
	}

	return int(result.DeletedCount), nil
}

func (c MongoRepository) DeleteAllUsers() (int, error) {
	filter := bson.D{}

	result, err := c.collection.DeleteMany(context.Background(), filter)
	if err != nil {
		return 0, err
	}

	return int(result.DeletedCount), nil
}
