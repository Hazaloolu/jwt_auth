package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hazaloolu/jwt_auth/initializers"
	"github.com/hazaloolu/jwt_auth/models"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.context) {
	// get the email/password off request body
	var body struct {
		Email    string
		password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// Hash the password

	hash, err := bcrypt.GenerateFromPassword([]byte(body.password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})

		return
	}

	// create the user
	user := models.User{Email: body.Email, Password: string(hash)}
	result := initializers.DB.create(&user)

	if result.Error != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})

	}
	// respond
	c.JSON(http.Status, gin.H{})

}
