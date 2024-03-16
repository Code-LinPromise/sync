package initializers

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitCors(router *gin.Engine) {
	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "PUT", "PATCH", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			if origin == "http://127.0.0.1:5173" || origin == "http://localhost:5173" || origin == "http://192.168.0.195:5173" {
				return true
			} else {
				log.Printf("%v is now allowed", origin)
				return false
			}
		},
		MaxAge: 12 * time.Hour,
	}))
}
