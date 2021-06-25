package game

import "C"
import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"webapp/api/utils"
)

const collectionName = "games"

type Game struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Rounds []Round            `bson:"rounds,omitempty"`
	Player []string           `bson:"player,omitempty"`
}

type Round struct {
	Value []int32   `bson:"value,omitempty"`
	Cards [][]uint8 `bson:"cards,omitempty"`
}

type UpdateGameStruct struct {
	Value []int32            `bson:"value,omitempty"`
	Cards [][]uint8          `bson:"cards,omitempty"`
	Game  primitive.ObjectID `bson:"game"`
}
type FindPlayersStruct struct {
	Player []string `bson:"player"`
}

func CreateGame(c *gin.Context) {
	q := utils.GetQuery()
	defer q.Close()

	collection := utils.GetCollection(q, collectionName)
	game := Game{}
	err := c.BindJSON(&game)
	if err != nil {
		log.Panic("no values transmitted")
	}
	res, err := collection.InsertOne(q.Ctx, game)
	if err != nil {
		log.Panic(err)
	}
	c.JSON(200, gin.H{
		"game": res,
	})
}
func UpdateGame(c *gin.Context) {
	q := utils.GetQuery()
	defer q.Close()
	json := UpdateGameStruct{}
	err := c.Bind(&json)
	if err != nil {
		log.Panic("an error occurred during binding")
	}
	game := Game{}
	collection := utils.GetCollection(q, collectionName)
	round := Round{Value: json.Value, Cards: json.Cards}
	update := bson.M{
		"$push": bson.M{"rounds": round},
	}
	err = collection.FindOneAndUpdate(q.Ctx, bson.M{"_id": json.Game}, update).Decode(&game)
	if err != nil {
		log.Panic("error occurred in database connection")
	}
	c.JSON(200, gin.H{
		"game": game,
	})
}

func FindGameByPlayers(c *gin.Context) {
	q := utils.GetQuery()
	defer q.Close()
	collection := utils.GetCollection(q, collectionName)

	var b []string
	b = append(b, c.Param("p1"))
	b = append(b, c.Param("p2"))
	b = append(b, c.Param("p3"))
	b = append(b, c.Param("p4"))
	query := bson.D{
		{
			"player",
			bson.D{
				{
					"$all",
					bson.A{b},
				},
			},
		},
	}

	game := Game{}
	err := collection.FindOne(q.Ctx, query).Decode(&game)
	if err != nil {
		return
	}

	c.JSON(200, gin.H{
		"game": game,
	})

}

func GetPlayerTotal(c *gin.Context) {
	var (
		pos    int
		result int32
	)
	nameParameter := c.Param("name")
	q := utils.GetQuery()
	defer q.Close()
	collection := utils.GetCollection(q, collectionName)
	cur, err := collection.Find(q.Ctx, bson.D{{"player", nameParameter}})
	if err != nil {
		log.Panic(err)
	}

	for cur.Next(q.Ctx) {
		game := Game{}
		err := cur.Decode(&game)
		if err != nil {
			log.Panic(err)
		}
		for i := 0; i < len(game.Player); i++ {
			if game.Player[i] == nameParameter {
				pos = i
				break
			}
		}
		for i := range game.Rounds {
			result += game.Rounds[i].Value[pos]
		}
	}

	c.JSON(200, gin.H{
		"result": result,
	})
}

func DeleteGameByID(c *gin.Context) {
	q := utils.GetQuery()
	defer q.Close()
	collection := utils.GetCollection(q, collectionName)

	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	query := bson.M{
		"_id": id,
	}

	obj, err := collection.DeleteOne(q.Ctx, query)
	if err != nil {
		return
	}
	c.JSON(200, gin.H{
		"DeletedGameCount": obj.DeletedCount,
	})
}

func DeleteAll(c *gin.Context) {
	q := utils.GetQuery()
	defer q.Close()
	collection := utils.GetCollection(q, collectionName)
	obj, err := collection.DeleteMany(q.Ctx, bson.M{})
	if err != nil {
		return
	}
	c.JSON(200, gin.H{
		"DeletedGameCount": obj.DeletedCount,
	})
}
