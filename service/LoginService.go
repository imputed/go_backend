package service

import (
	"bytes"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"os"
	"webapp/api/dto"
	"webapp/api/utils"
)

const envCollectionSpecifier = "UserCollectionName"

type LoginService interface {
	LoginUser(name string, password string) (bool, dto.ReturnUser)
}
type loginInformation struct {
	name     string
	password string
}

func StaticLoginService() LoginService {
	return &loginInformation{
		name:     "admin@wesionary.team",
		password: "admin",
	}
}
func (info *loginInformation) LoginUser(name string, password string) (bool, dto.ReturnUser) {
	q := utils.GetQuery()
	defer q.Close()
	user := dto.RegisterUser{}

	collection := utils.GetCollection(q, os.Getenv(envCollectionSpecifier))
	err := collection.FindOne(q.Ctx, bson.D{{"name", name}}).Decode(&user)
	returnUser := dto.ReturnUser{}
	returnUser.ReadRegisterUser(user)

	if err != nil {
		log.Panicln("user not found")
	}
	return bytes.Compare(user.EncryptedPassword, utils.Encrypt(password)) == 0, returnUser
}
