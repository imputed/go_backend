package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"webapp/api/User"
	Game "webapp/api/game"
	"webapp/api/utils"
)

func main() {
	const address = "0.0.0.0:3002"

	r := gin.Default()
	r.Use(utils.CorsMiddleware())
	r.GET("/user", User.GetUser)
	r.POST("/user", User.CreateUser)
	r.GET("/user/:id", User.FilterUser)
	r.GET("/game/player/:p1/:p2/:p3/:p4", Game.FindGameByPlayers)
	r.POST("/game", Game.CreateGame)
	r.POST("/game/update", Game.UpdateGame)

	err := r.Run("0.0.0.0:3002") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		log.Fatalf("Cannot run on %v", address)
	}
}
