package service

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"time"
	"webapp/api/User"
	"webapp/api/dto"
)

//jwt service
type JWTService interface {
	GenerateToken(email string, isUser bool) string
	ValidateToken(token string) (*jwt.Token, error)
}
type authCustomClaims struct {
	Name string `json:"name"`
	User bool   `json:"user"`
	jwt.StandardClaims
}

type jwtServices struct {
	secretKey string
	issuer    string
}

func JWTAuthService() JWTService {
	return &jwtServices{
		secretKey: os.Getenv("JWT_SECRET"),
		issuer:    "Schafkopf4Friends",
	}
}

func (service *jwtServices) GenerateToken(email string, isUser bool) string {
	claims := &authCustomClaims{
		email,
		isUser,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    service.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func ValidateTokenAndReturnUser(c *gin.Context) {
	token := dto.TokenValidation{}
	err := c.BindJSON(&token)
	if err != nil {
		log.Panic(fmt.Sprintf("Error in parsing: %v", err))
	}
	claims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(token.Token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		log.Panic(fmt.Sprintf("Error in validation: %v", err))
	}

	user := dto.ReturnUser{}
	user = User.FilterUserByName(fmt.Sprintf("%v", claims["name"]))

	if token.Token != "" {
		c.JSON(http.StatusOK, gin.H{
			"token": token,
			"user":  user,
		})
	} else {
		c.JSON(http.StatusUnauthorized, nil)
	}
}

func (service *jwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])
		}
		return []byte(service.secretKey), nil
	})
}
