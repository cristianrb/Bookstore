package app

import (
	"cristianrb/controllers/ping"
	"cristianrb/controllers/users"
	"github.com/gin-gonic/gin"
)

func mapUrls(router *gin.Engine) {
	router.GET("/ping", ping.Ping)

	router.POST("/users", users.CreateUser)
	router.GET("/users/:user_id", users.GetUser)
	router.GET("/users", users.SearchUser)
}
