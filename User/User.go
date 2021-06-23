package User

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"

	"webapp/api/utils"
)

type User struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name string `bson:"name,omitempty"`
	Role string `bson:"role,omitempty"`
	Mail string `bson:"mail,omitempty"`
}

func GetUser(c *gin.Context) {
	q := utils.GetQuery()
	defer q.Close()

	collection := q.Client.Database("test").Collection("users")
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
	collection := q.Client.Database("test").Collection("users")

	json:= User{}
	c.BindJSON(&json)

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

	collection := q.Client.Database("test").Collection("users")
	res, err := collection.Find(q.Ctx, bson.M{"_id": id})
	var episodesFiltered []bson.M
	if err = res.All(q.Ctx, &episodesFiltered); err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(200, gin.H{
		"user": episodesFiltered,
	})
}
