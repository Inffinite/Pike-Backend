package routes

import (
	"pike/controllers"

	"github.com/gin-gonic/gin"
)

func PikeRoute(router *gin.Engine) {
	router.GET("/ping", controllers.Ping())
	router.GET("/shutdown", controllers.ShutdownDevice())
}
