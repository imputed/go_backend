package utils

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

const dbName = "test"

type Query struct {
	Ctx    context.Context
	Client *mongo.Client
	Close  context.CancelFunc
	Err    error
}

func GetQuery() Query {
	q := Query{}
	q.Ctx, q.Close = context.WithTimeout(context.Background(), 10*time.Second)
	q.Client, q.Err = mongo.Connect(q.Ctx, options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	return q
}

func GetCollection(q Query, collectionName string) *mongo.Collection {
	return q.Client.Database(dbName).Collection(collectionName)
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

func Encrypt(text string) []byte {
	key := os.Getenv("AES_KEY")
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Panic(err)
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		log.Panic(err)
	}

	nonce := make([]byte, gcm.NonceSize())
	return gcm.Seal(nonce, nonce, []byte(text), nil)
}
