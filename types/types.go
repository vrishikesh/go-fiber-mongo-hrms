package types

import "go.mongodb.org/mongo-driver/mongo"

type MongoInstance struct {
	Client *mongo.Client
	DB     *mongo.Database
}

type Employee struct {
	ID          string  `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string  `json:"name,omitempty"`
	Salary      float64 `json:"salary,omitempty"`
	Age         int     `json:"age,omitempty"`
}
