package main

import (
	"net/http"
	"tictactoewithapi/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/board", controller.Get_board)
	r.POST("/reset", controller.Reset_game)
	r.POST("/move", controller.Apply_move)
	r.POST("/start", controller.Start_game)
	r.Run(":3000")
}
