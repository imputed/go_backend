package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"webapp/api/User"
	"webapp/api/dto"
	"webapp/api/service"
)

func RegisterUser(c *gin.Context) {
	registerUser := dto.RegisterUser{}
	err := c.Bind(&registerUser)
	if err != nil {
		log.Panic("registering user not possible")
	}
	res, err := User.CreateUser(registerUser)
	if err != nil {
		log.Panic(" error when registering user ")
	}
	registerUser.UserID = res.InsertedID
	res, err = service.Register(registerUser)
	if err != nil {
		log.Panic("error when inserting user credentials in registration")
	}
	c.JSON(http.StatusOK, gin.H{
		"res": res,
	})

}
