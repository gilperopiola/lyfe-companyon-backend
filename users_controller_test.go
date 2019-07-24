package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"testing"

	"github.com/gilperopiola/lyfe-companyon-backend/utils"
	"github.com/stretchr/testify/assert"
)

func TestCreateUserController(t *testing.T) {
	cfg.Setup("testing")
	db.Setup(cfg)
	defer db.Close()
	rtr.Setup()
	token := generateTestingToken()

	user := &User{
		Email:     "email",
		Password:  hash("email", "password"),
		FirstName: "firstName",
		LastName:  "lastName",
	}

	response := user.GenerateTestRequest(token, "POST", "")
	json.Unmarshal(response.Body.Bytes(), &user)
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "email", user.Email)
}

func TestGetUserController(t *testing.T) {
	cfg.Setup("testing")
	db.Setup(cfg)
	defer db.Close()
	rtr.Setup()
	token := generateTestingToken()

	user := &User{
		Email:     "email",
		Password:  hash("email", "password"),
		FirstName: "firstName",
		LastName:  "lastName",
	}
	user, _ = user.Create()

	response := user.GenerateTestRequest(token, "GET", "/"+strconv.Itoa(user.ID))
	json.Unmarshal(response.Body.Bytes(), &user)
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "email", user.Email)
}

func TestUpdateUserController(t *testing.T) {
	cfg.Setup("testing")
	db.Setup(cfg)
	defer db.Close()
	rtr.Setup()
	token := generateTestingToken()

	user := &User{
		Email:     "email",
		Password:  hash("email", "password"),
		FirstName: "firstName",
		LastName:  "lastName",
	}
	user, _ = user.Create()

	user.Email = "email2"
	user.Enabled = !user.Enabled

	response := user.GenerateTestRequest(token, "PUT", "/"+strconv.Itoa(user.ID))
	json.Unmarshal(response.Body.Bytes(), &user)
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "email2", user.Email)
	assert.False(t, user.Enabled)
}

func TestSearchUserController(t *testing.T) {
	cfg.Setup("testing")
	db.Setup(cfg)
	defer db.Close()
	rtr.Setup()
	token := generateTestingToken()

	user := &User{}
	for i := 15; i <= 25; i++ {
		user = &User{
			Email:     "email" + utils.ToString(i),
			Password:  hash("email"+utils.ToString(i), "password"),
			FirstName: "firstName" + utils.ToString(i),
			LastName:  "lastName" + utils.ToString(i),
		}
		user, _ = user.Create()
	}

	params := &SearchParameters{
		Filter:        "Name1",
		SortField:     "id",
		SortDirection: "DESC",
		Limit:         3,
		Offset:        1,
	}

	users := make([]*User, 0)
	response := user.GenerateTestRequest(token, "GET", getSearchURL(params))
	json.Unmarshal(response.Body.Bytes(), &users)
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, params.Limit, len(users))
	assert.NotZero(t, users[0].ID)
	assert.Equal(t, "email18", users[0].Email)
	assert.Equal(t, "firstName18", users[0].FirstName)
	assert.Equal(t, "lastName18", users[0].LastName)
	assert.True(t, users[0].Enabled)
	assert.NotZero(t, users[0].DateCreated)
}
