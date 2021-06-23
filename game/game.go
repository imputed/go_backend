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
	c.BindJSON(&game)
	res, err := collection.InsertOne(q.Ctx, game)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(200, gin.H{
		"game": res,
	})
}
func UpdateGame(c *gin.Context) {
	q := utils.GetQuery()
	defer q.Close()
	json := UpdateGameStruct{}
	c.Bind(&json)
	game := Game{}
	collection := utils.GetCollection(q, collectionName)
	round := Round{Value: json.Value, Cards: json.Cards}
	update := bson.M{
		"$push": bson.M{"rounds": round},
	}
	collection.FindOneAndUpdate(q.Ctx, bson.M{"_id": json.Game}, update).Decode(&game)
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
	q := utils.GetQuery()
	defer q.Close()
	collection := utils.GetCollection(q, collectionName)
	name := c.Param("name")
	query := bson.D{
		{
			"player", name},
	}
	var result int32
	cur, err := collection.Find(q.Ctx, query)
	pos := -1
	for cur.Next(q.Ctx) {

		// create a value into which the single document can be decoded
		game := Game{}
		err := cur.Decode(&game)
		if err != nil {
			log.Fatal(err)
		}
		if pos == -1 {
			for i := 0; i < len(game.Player); i++ {
				if game.Player[i] == name {
					pos = i
					break
				}
			}
		}
		for _, r := range game.Rounds {
			result += r.Value[pos]
		}
	}
	if err != nil {
		return
	}
	c.JSON(200, gin.H{
		"result": result,
	})

}
