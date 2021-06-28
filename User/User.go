package User

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"os"
	"webapp/api/dto"

	"webapp/api/utils"
)

const envCollectionSpecifier = "UserCollectionName"

type User struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"name,omitempty"`
	Role string             `bson:"role,omitempty"`
	Mail string             `bson:"mail,omitempty"`
}

func GetUser(c *gin.Context) {
	q := utils.GetQuery()
	defer q.Close()

	collection := utils.GetCollection(q, os.Getenv(envCollectionSpecifier))
	cur, err := collection.Find(q.Ctx, bson.M{})
	if err != nil {
		log.Println(err)
	} else {
		re := []dto.RegisterUser{}
		for cur.Next(q.Ctx) {
			result := dto.RegisterUser{}
			err := cur.Decode(&result)
			if err != nil {
				log.Fatal(err)
			} else {
				re = append(re, result)
			}
			if err := cur.Err(); err != nil {
				log.Fatal(err)
			}
		}
		c.JSON(200, gin.H{
			"user": re,
		})
	}
}

func CreateUser(registeruser dto.RegisterUser) (*mongo.InsertOneResult, error) {
	q := utils.GetQuery()
	defer q.Close()
	collection := utils.GetCollection(q, os.Getenv(envCollectionSpecifier))

	res, err := collection.InsertOne(q.Ctx, registeruser)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return res, nil
}

func FilterUserByID(c *gin.Context) {
	q := utils.GetQuery()
	defer q.Close()

	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	user := dto.RegisterUser{}
	collection := utils.GetCollection(q, os.Getenv(envCollectionSpecifier))

	err := collection.FindOne(q.Ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"user": user,
		})
	}
	c.JSON(200, gin.H{
		"user": user,
	})
}
func FilterUserByName(name string) dto.ReturnUser {
	q := utils.GetQuery()
	defer q.Close()

	user := dto.ReturnUser{}
	collection := utils.GetCollection(q, os.Getenv(envCollectionSpecifier))

	err := collection.FindOne(q.Ctx, bson.M{"name": name}).Decode(&user)
	if err != nil {
		log.Panic("User Not Found q")
	}
	return user
}
func Register(c *gin.Context) {
	q := utils.GetQuery()
	defer q.Close()
	collection := utils.GetCollection(q, os.Getenv(envCollectionSpecifier))
	registerUser := dto.RegisterUser{}
	err := c.Bind(&registerUser)
	if err != nil {
		log.Panic("registering user not possible")
	}
	registerUser.EncryptedPassword = utils.Encrypt(registerUser.Password)
	res, err := collection.InsertOne(q.Ctx, registerUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "An error occured during registration",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"res": res,
	})

}
