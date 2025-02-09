package main

import (
	"inkforge/setup"

	"github.com/gin-gonic/gin"
)

func init() {
	setup.LoadEnvVariables()
	setup.ConnectDatabase()
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"HELLO": "WORLD"})
	})
	r.Run()
}
