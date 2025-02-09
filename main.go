package main

import (
	"inkforge/api/controllers"
	"inkforge/database/migrations"
	"inkforge/setup"

	"github.com/gin-gonic/gin"
)

func init() {
	setup.LoadEnvVariables()
	setup.ConnectDatabase()
	migrations.MigrateTables()
}

func main() {
	r := gin.Default()
	r.POST("/api/signup", controllers.SignUp)
	r.Run()
}
