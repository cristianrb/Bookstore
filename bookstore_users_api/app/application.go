package app

import "github.com/gin-gonic/gin"

const PORT = "8080"

func StartApplication() {
	router := gin.Default()
	mapUrls(router)
	router.Run(":" + PORT)
}
