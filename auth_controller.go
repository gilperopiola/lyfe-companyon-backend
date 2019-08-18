package main

import (
	"net/http"

	"github.com/gilperopiola/lyfe-companyon-backend/utils"
	"github.com/gin-gonic/gin"
)

//SignUp takes {email, password, repeatPassword}. It creates a user and returns it.
func SignUp(c *gin.Context) {
	var auth Auth
	c.BindJSON(&auth)

	if auth.Email == "" || auth.Password == "" || auth.RepeatPassword == "" {
		c.JSON(http.StatusBadRequest, "all fields required")
		return
	}

	if auth.Password != auth.RepeatPassword {
		c.JSON(http.StatusBadRequest, "passwords don't match")
		return
	}

	hashedPassword := utils.Hash(auth.Email, auth.Password)

	user := &User{
		Email:    auth.Email,
		Password: hashedPassword,
	}

	user, err := user.Create()
	if err != nil {
		c.JSON(http.StatusBadRequest, db.BeautifyError(err))
		return
	}

	c.JSON(http.StatusOK, user)
}

//Login takes {username, password}, checks if the user exists and returns it
func Login(c *gin.Context) {
	var auth Auth
	c.BindJSON(&auth)

	if auth.Email == "" || auth.Password == "" {
		c.JSON(http.StatusBadRequest, "all fields required")
		return
	}

	hashedPassword := utils.Hash(auth.Email, auth.Password)

	user := &User{
		Email:    auth.Email,
		Password: hashedPassword,
	}

	user, err := user.Login()
	if err != nil {
		c.JSON(http.StatusBadRequest, db.BeautifyError(err))
		return
	}

	c.JSON(http.StatusOK, user)
}
