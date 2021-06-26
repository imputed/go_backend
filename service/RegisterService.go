package service

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"webapp/api/dto"
	"webapp/api/utils"
)

const registerCollectionName = "credentials"

func Register(registerUser dto.RegisterUser) (*mongo.InsertOneResult, error) {
	q := utils.GetQuery()
	defer q.Close()
	collection := utils.GetCollection(q, registerCollectionName)
	basicUser := dto.RegisterBasicUser{Name: registerUser.Name, Password: utils.Encrypt(registerUser.Password), UserID: registerUser.UserID}
	res, err := collection.InsertOne(q.Ctx, bson.M{"user": basicUser})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return res, nil
}
