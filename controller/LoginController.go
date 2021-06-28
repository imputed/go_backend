package controller

import (
	"github.com/gin-gonic/gin"
	"webapp/api/dto"
	"webapp/api/service"
)

type LoginController interface {
	Login(ctx *gin.Context) (string, dto.ReturnUser)
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

func (controller *loginController) Login(c *gin.Context) (string, dto.ReturnUser) {
	var credential dto.LoginCredentials
	c.BindJSON(&credential)
	isUserAuthenticated, user := controller.loginService.LoginUser(credential.Name, credential.Password)
	if isUserAuthenticated {
		return controller.jWtService.GenerateToken(credential.Name, true), user
	}
	return "", user
}
