package main

import (
	"testing"

	"github.com/gilperopiola/lyfe-companyon-backend/utils"
	"github.com/stretchr/testify/assert"
)

func TestLoginUser(t *testing.T) {
	cfg.Setup("testing")
	db.Setup(cfg)
	defer db.Close()

	user := &User{
		Email:     "email",
		Password:  utils.Hash("email", "password"),
		FirstName: "first_name",
		LastName:  "last_name",
	}

	user, _ = user.Create()
	user, err := user.Login()

	assert.NoError(t, err)
	assert.NotZero(t, user.ID)
	assert.Equal(t, "email", user.Email)
	assert.Equal(t, "first_name", user.FirstName)
	assert.Equal(t, "last_name", user.LastName)
	assert.Equal(t, true, user.Enabled)
	assert.NotZero(t, user.DateCreated)
	assert.NotZero(t, user.Token)
}

func TestCreateUser(t *testing.T) {
	cfg.Setup("testing")
	db.Setup(cfg)
	defer db.Close()

	user := &User{
		Email:     "email",
		Password:  utils.Hash("email", "password"),
		FirstName: "firstName",
		LastName:  "lastName",
	}

	user, err := user.Create()

	assert.NoError(t, err)
	assert.NotZero(t, user.ID)
	assert.Equal(t, "email", user.Email)
	assert.Equal(t, "firstName", user.FirstName)
	assert.Equal(t, "lastName", user.LastName)
	assert.Equal(t, true, user.Enabled)
	assert.NotZero(t, user.DateCreated)
}

func TestGetUser(t *testing.T) {
	cfg.Setup("testing")
	db.Setup(cfg)
	defer db.Close()

	user := &User{
		Email:     "email",
		Password:  utils.Hash("email", "password"),
		FirstName: "firstName",
		LastName:  "lastName",
	}

	user, _ = user.Create()
	user, err := user.Get()

	assert.NoError(t, err)
	assert.NotZero(t, user.ID)
	assert.Equal(t, "email", user.Email)
	assert.Equal(t, "firstName", user.FirstName)
	assert.Equal(t, "lastName", user.LastName)
	assert.Equal(t, true, user.Enabled)
	assert.NotZero(t, user.DateCreated)
}

func TestUpdateUser(t *testing.T) {
	cfg.Setup("testing")
	db.Setup(cfg)
	defer db.Close()

	user := &User{
		Email:     "email",
		Password:  utils.Hash("email", "password"),
		FirstName: "firstName",
		LastName:  "lastName",
	}
	user, _ = user.Create()

	user.Email = "email2"
	user.FirstName = "firstName2"
	user.LastName = "lastName2"
	user.Enabled = !user.Enabled

	user, err := user.Update()
	assert.NoError(t, err)
	assert.NotZero(t, user.ID)
	assert.Equal(t, "email2", user.Email)
	assert.Equal(t, "firstName2", user.FirstName)
	assert.Equal(t, "lastName2", user.LastName)
	assert.False(t, user.Enabled)
}

func TestSearchUsers(t *testing.T) {
	cfg.Setup("testing")
	db.Setup(cfg)
	defer db.Close()

	user := &User{}
	for i := 15; i <= 25; i++ {
		user = &User{
			Email:     "email" + utils.ToString(i),
			Password:  utils.Hash("email"+utils.ToString(i), "password"),
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

	users, err := user.Search(params)
	assert.NoError(t, err)

	assert.Equal(t, params.Limit, len(users))
	assert.NotZero(t, users[0].ID)
	assert.Equal(t, "email18", users[0].Email)
	assert.Equal(t, "firstName18", users[0].FirstName)
	assert.Equal(t, "lastName18", users[0].LastName)
	assert.True(t, users[0].Enabled)
	assert.NotZero(t, users[0].DateCreated)
}
