package router

import (
	"github.com/OtavioPontes/gopportunities/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handler.ShowOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.GET("/openings", handler.ListOpeningsHandler)
	}

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "teste",
		})
	})
}
