package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type LoginCredentials struct {
	Name     string `bson:"email,omitempty"`
	Password string `bson:"password,omitempty"`
}

type RegisterUser struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name              string             `bson:"name,omitempty" json:"name"`
	Password          string             `bson:"-" json:"-"`
	EncryptedPassword []byte             `bson:"encryptedpassword,omitempty" json:"-"`
	Role              string             `bson:"role,omitempty" json:"role"`
	Mail              string             `bson:"mail,omitempty" json:"mail"`
}
