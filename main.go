package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	//inicialza o router usando configuraçoes default do gin
	r := gin.Default()
	//definindo uma rota
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
