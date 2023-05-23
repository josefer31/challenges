package server

import (
	"github.com/gin-gonic/gin"
	"polaris/cmd/polaris/boostrap"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	adController := boostrap.ProvideAdController()
	router.POST("/ads", adController.HandlerCreationAd)
	router.GET("/ads/:adId", adController.HandlerFindAd)
	return router
}
