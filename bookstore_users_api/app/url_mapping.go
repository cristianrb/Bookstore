package app

import (
	"cristianrb/controllers/ping"
	"cristianrb/controllers/users"
	"github.com/gin-gonic/gin"
)

func mapUrls(router *gin.Engine) {
	router.GET("/ping", ping.Ping)

	router.POST("/users", users.Create)
	router.GET("/users/:user_id", users.Get)
	router.PUT("/users/:user_id", users.Update)
	router.PATCH("/users/:user_id", users.Update)
	router.DELETE("/users/:user_id", users.Delete)
	router.GET("/internal/users/search", users.Search)
	router.POST("/users/login", users.Login)
}
