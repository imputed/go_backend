package controller

import (
	"github.com/gin-gonic/gin"
	"webapp/api/dto"
	"webapp/api/service"
)

type LoginController interface {
	Login(ctx *gin.Context) string
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
	isUserAuthenticated := controller.loginService.LoginUser(credential.Name, credential.Password)
	if isUserAuthenticated {
		return controller.jWtService.GenerateToken(credential.Name, true)
	}
	return ""
}
