package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stavagg/petGoApi/internal/config" // Измени на своё имя модуля, если нужно
)

func main() {
	cfg := config.Load()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello from upgraded server!",
			"port":    cfg.Port, // Для теста
		})
	})

	r.Run(cfg.Port)
}
