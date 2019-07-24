package main

import (
	"net/http"

	"github.com/gilperopiola/lyfe-companyon-backend/utils"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user *User
	c.BindJSON(&user)

	if user.Email == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, "email and password required")
		return
	}

	user.Password = hash(user.Email, user.Password)

	user, err := user.Create()
	if err != nil {
		c.JSON(http.StatusBadRequest, db.BeautifyError(err))
		return
	}

	c.JSON(http.StatusOK, user)
}

func GetUser(c *gin.Context) {
	user := &User{ID: utils.ToInt(c.Param("id_user"))}

	user, err := user.Get()
	if err != nil {
		c.JSON(http.StatusBadRequest, db.BeautifyError(err))
		return
	}

	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	var user *User
	c.BindJSON(&user)
	user.ID = utils.ToInt(c.Param("id_user"))

	if user.Email == "" {
		c.JSON(http.StatusBadRequest, "email required")
		return
	}

	user, err := user.Update()
	if err != nil {
		c.JSON(http.StatusBadRequest, db.BeautifyError(err))
		return
	}

	c.JSON(http.StatusOK, user)
}

func SearchUsers(c *gin.Context) {
	user := &User{}

	params := &SearchParameters{
		Filter:        c.Query("filter"),
		SortField:     c.Query("sortField"),
		SortDirection: c.Query("sortDirection"),
		Limit:         utils.ToInt(c.Query("limit")),
		Offset:        utils.ToInt(c.Query("offset")),
	}

	users, err := user.Search(params)
	if err != nil {
		c.JSON(http.StatusBadRequest, db.BeautifyError(err))
		return
	}

	c.JSON(http.StatusOK, users)
}
