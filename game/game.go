package game

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
	Winner string `bson:"winner"`
}

type UpdateGameStruct struct {
	Winner string             `bson:"winner"`
	Game   primitive.ObjectID `bson:"game"`
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
	collection := utils.GetCollection(q, collectionName)
	round := Round{Winner: json.Winner}
	update := bson.M{
		"$push": bson.M{"rounds": round},
	}
	res := collection.FindOneAndUpdate(q.Ctx, bson.M{"_id": json.Game}, update)
	c.JSON(200, gin.H{
		"game": res,
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
