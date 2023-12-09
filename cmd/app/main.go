package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ididie/ifidie_backend/internal/services/auth/login"
)

func main() {
	fmt.Println(login.Greet2())

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
