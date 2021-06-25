package dto

type LoginCredentials struct {
	Email    string `bson:"email,omitempty"`
	Password string `bson:"password,omitempty"`
}
