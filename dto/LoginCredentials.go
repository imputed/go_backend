package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type LoginCredentials struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type ReturnUser struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name string             `bson:"name,omitempty" json:"name"`
	Role string             `bson:"role,omitempty" json:"role"`
	Mail string             `bson:"mail,omitempty" json:"mail"`
}

func (r *ReturnUser) ReadRegisterUser(user RegisterUser) {
	r.ID = user.ID
	r.Name = user.Name
	r.Role = user.Role
	r.Mail = user.Mail
}

type RegisterUser struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name              string             `bson:"name,omitempty" json:"name"`
	Password          string             `bson:"-" json:"password"`
	EncryptedPassword []byte             `bson:"encryptedpassword,omitempty" json:"-"`
	Role              string             `bson:"role,omitempty" json:"role"`
	Mail              string             `bson:"mail,omitempty" json:"mail"`
}

type TokenValidation struct {
	Token string `json:"token"`
}
