package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type LoginCredentials struct {
	Email    string `bson:"email,omitempty"`
	Password string `bson:"password,omitempty"`
}

type RegisterUser struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"name,omitempty"`
	Password string             `bson:"password,omitempty"`
	Role     string             `bson:"role,omitempty"`
	Mail     string             `bson:"mail,omitempty"`
	UserID   interface{}        `bson:"userid,omitempty"`
}
type RegisterBasicUser struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"name,omitempty"`
	Password []byte             `bson:"password,omitempty"`
	UserID   interface{}        `bson:"userid,omitempty"`
}
