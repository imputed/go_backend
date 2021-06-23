package utils

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Query struct {
	Ctx    context.Context
	Client *mongo.Client
	Close  context.CancelFunc
	Err    error
}

func GetQuery() Query {
	q := Query{}
	q.Ctx, q.Close = context.WithTimeout(context.Background(), 10*time.Second)
	q.Client, q.Err = mongo.Connect(q.Ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	return q
}

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

