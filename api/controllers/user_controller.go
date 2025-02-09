package controllers

import (
	"net/http"

	"inkforge/database/models"
	"inkforge/pkg"
	"inkforge/setup"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	body := pkg.Credential{}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	pkg.ValidateSignup(body)
	hashPass, err := pkg.HashPassword([]byte(body.Password))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to hash password",
		})
		return
	}
	user := models.User{Email: body.Email, Username: body.Username, FirstName: body.FirstName, LastName: body.LastName, PasswordHash: hashPass, Role: "user"}
	setup.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{})
}

func SignIn(c *gin.Context) {
	body := pkg.Credential{}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
}
