package controller

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"webapp/api/dto"
	"webapp/api/service"
)

type LoginController interface {
	Login(ctx *gin.Context) string
	Register(user dto.RegisterUser) (*mongo.InsertOneResult, error)
}

type loginController struct {
	loginService service.LoginService
	jWtService   service.JWTService
}

func LoginHandler(loginService service.LoginService, jWtService service.JWTService) *loginController {
	return &loginController{
		loginService: loginService,
		jWtService:   jWtService,
	}
}

func (controller *loginController) Login(c *gin.Context) string {
	var credential dto.LoginCredentials
	c.Bind(&credential)
	isUserAuthenticated := controller.loginService.LoginUser(credential.Email, credential.Password)
	if isUserAuthenticated {
		return controller.jWtService.GenerateToken(credential.Email, true)
	}
	return ""
}

func (controller *loginController) Register(registerUser dto.RegisterUser) (*mongo.InsertOneResult, error) {
	res, _ := controller.loginService.Register(registerUser)
	return res, nil
}
