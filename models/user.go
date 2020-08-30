package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// User es la estructura de un usuario
type User struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name string             `bson:"name, omitempty" json:"name"`
	Age  int                `bson:"age,omitempty" json:"age"`
}
