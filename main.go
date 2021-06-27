package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	ENV_PORT := os.Getenv("PORT")

	v1 := router.Group("/api/v1")
	{
		v1.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "chenfangting 哈哈111asd哈",
			})
		})
	}

	router.Run(":" + string(ENV_PORT)) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	fmt.Println("server running on port", ENV_PORT)
}
