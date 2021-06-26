package service

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"
	"webapp/api/dto"
	"webapp/api/utils"
)

const collectionName = "credentials"

type LoginService interface {
	LoginUser(email string, password string) bool
	Register(registerUser dto.RegisterUser) (*mongo.InsertOneResult, error)
}

type loginInformation struct {
	email    string
	password string
}

func (info *loginInformation) Register(registerUser dto.RegisterUser) (*mongo.InsertOneResult, error) {
	q := utils.GetQuery()
	defer q.Close()
	collection := utils.GetCollection(q, collectionName)
	key := os.Getenv("AES_KEY")
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Println(err)
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println(err)
	}

	nonce := make([]byte, gcm.NonceSize())
	pw := gcm.Seal(nonce, nonce, []byte(registerUser.Password), nil)
	basicUser := dto.RegisterBasicUser{Name: registerUser.Name, Password: pw, UserID: registerUser.UserID}
	res, err := collection.InsertOne(q.Ctx, bson.M{"user": basicUser})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return res, nil
}

func StaticLoginService() LoginService {
	return &loginInformation{
		email:    "admin@wesionary.team",
		password: "admin",
	}
}
func (info *loginInformation) LoginUser(email string, password string) bool {
	return info.email == email && info.password == password
}
