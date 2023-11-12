package router

import "github.com/gin-gonic/gin"

func Initialize() {
	//initialize router
	router := gin.Default()
	//initialize routes
	initializeRoutes(router)
	router.Run() // listen and serve on 0.0.0.0:8080
}
