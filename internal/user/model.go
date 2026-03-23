package user

import "go.mongodb.org/mongo-driver/v2/bson"

type User struct {
	ID      bson.ObjectID `bson:"_id,omitempt"`
	Name    string        `bson:"name,omitempty"`
	Age     int           `bson:"age,omitempty"`
	Country string        `bson:"country,omitempty"`
}

type UserRequest struct {
	Name    string `json:"name,omitempty"`
	Age     int    `json:"age,omitempty"`
	Country string `json:"country,omitempty"`
}

type UserResponse struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}
