package controllers

import (
	"log"
	"net/http"
	"os/exec"
	"pike/responses"

	"github.com/gin-gonic/gin"
)

func Ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, responses.PikeResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"message": "pong"}})
	}
}

func ShutdownDevice() gin.HandlerFunc {
	return func(c *gin.Context) {
		cmd := exec.Command("shutdown", "-h", "1")
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, responses.PikeResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"message": "pong"}})
	}
}
