package main

import (
	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.String(200, "Hello World!")
}

func main() {
	router := gin.Default()
	router.GET("/", Home)
	router.Run(":5000")
}
