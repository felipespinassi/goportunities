package router

import (
	"github.com/felipespinassi/goportunities/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handler.ShowOppeningHandler)
		v1.POST("/opening", handler.CreateOppeningHandler)
		v1.DELETE("/opening", handler.DeleteOppeningHandler)
		v1.PUT("/opening", handler.UpdateOppeningHandler)
		v1.GET("/openings", handler.ListOppeningHandler)
	}

}
