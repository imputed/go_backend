package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"webapp/api/User"
	"webapp/api/controller"
	Game "webapp/api/game"
	"webapp/api/service"
	"webapp/api/utils"
)

func main() {
	const address = "0.0.0.0:3002"
	godotenv.Load(".env")

	var loginService service.LoginService = service.StaticLoginService()
	var jwtService service.JWTService = service.JWTAuthService()
	var loginController controller.LoginController = controller.LoginHandler(loginService, jwtService)

	r := gin.Default()
	r.Use(utils.CorsMiddleware())
	r.GET("/user", User.GetUser)

	r.GET("/user/:id", User.FilterUser)
	r.GET("/user/total/:name", Game.GetPlayerTotal)
	r.GET("/game/player/:p1/:p2/:p3/:p4", Game.FindGameByPlayers)
	r.POST("/game", Game.CreateGame)
	r.POST("/game/update", Game.UpdateGame)
	r.DELETE("game/:id", Game.DeleteGameByID)
	r.DELETE("game/", Game.DeleteAll)
	r.POST("/register", controller.RegisterUser)

	r.POST("/login", func(c *gin.Context) {
		token := loginController.Login(c)
		jwtService.ValidateToken(token)
		if token != "" {
			c.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			c.JSON(http.StatusUnauthorized, nil)
		}
	})
	err := r.Run(os.Getenv("ADDRESS")) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		log.Fatalf("Cannot run on %v", address)
	}
}
