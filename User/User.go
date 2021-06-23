package User

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"

	"webapp/api/utils"
)

const collectionName = "users"

type User struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"name,omitempty"`
	Role string             `bson:"role,omitempty"`
	Mail string             `bson:"mail,omitempty"`
}

func GetUser(c *gin.Context) {
	q := utils.GetQuery()
	defer q.Close()

	collection := utils.GetCollection(q, collectionName)
	cur, err := collection.Find(q.Ctx, bson.M{})
	if err != nil {
		log.Println(err)
	} else {
		re := []bson.D{}
		for cur.Next(q.Ctx) {
			var result bson.D
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
			"results": re,
		})
	}
}

func CreateUser(c *gin.Context) {
	q := utils.GetQuery()
	defer q.Close()
	collection := utils.GetCollection(q, collectionName)

	json := User{}
	err := c.BindJSON(&json)
	if err != nil {
		log.Println("An Error Occured while marshaling")
	}
	res, err := collection.InsertOne(q.Ctx, bson.M{"user": json})
	if err != nil {
		log.Println(err)
	} else {
		c.JSON(200, gin.H{
			"user": res,
		})
	}

}

func FilterUser(c *gin.Context) {
	q := utils.GetQuery()
	defer q.Close()

	id, _ := primitive.ObjectIDFromHex(c.Param("id"))

	collection := utils.GetCollection(q, collectionName)

	res, err := collection.Find(q.Ctx, bson.M{"_id": id})

	var users []bson.M
	if err = res.All(q.Ctx, &users); err != nil {
		log.Fatal(err)
	}
	c.JSON(200, gin.H{
		"user": users,
	})
}
