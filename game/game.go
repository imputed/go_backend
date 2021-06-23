package game

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"webapp/api/utils"
)

type Game struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Rounds []Round            `bson:"rounds"`
	Player []string           `bson:"player"`
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
	collection := q.Client.Database("test").Collection("games")
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
	collection := q.Client.Database("test").Collection("games")
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

	json := FindPlayersStruct{}
	c.Bind(&json)

	game := Game{}

	collection := q.Client.Database("test").Collection("games")
	query :=bson.D{{
		"player",
		bson.D{{
			"$in",
			bson.D{{"$all", bson.A{json.Player}},
			},
			},
		},
	}}

	collection.FindOne(q.Ctx, query).Decode(&game)
	c.JSON(200, gin.H{
		"game": game,
	})
}
