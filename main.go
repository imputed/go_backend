package main

import (
	"github.com/gin-gonic/gin"
	"webapp/api/User"
	Game "webapp/api/game"
	"webapp/api/utils"
)





func main() {
	r := gin.Default()
	r.Use(utils.CorsMiddleware())
	r.GET("/user", User.GetUser)
	r.POST("/user", User.CreateUser)
	r.GET("/user/:id", User.FilterUser)
	r.GET("/game/player", Game.FindGameByPlayers)
	r.POST("/game", Game.CreateGame)
	r.POST("/game/update", Game.UpdateGame)

	r.Run("0.0.0.0:3002") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}


