package app

import (
	"cristianrb/logger"
	"github.com/gin-gonic/gin"
)

const PORT = "8080"

func StartApplication() {
	router := gin.Default()
	mapUrls(router)
	logger.Info("about to start the application...")
	router.Run(":" + PORT)
}
