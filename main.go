package main

import (
	"os"
	// "pike/configs"
	"pike/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.SetTrustedProxies([]string{os.Getenv("TRUSTED_PROXY")})
	routes.PikeRoute(router)
	router.Run(":8080")
}
