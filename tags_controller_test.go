package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"testing"

	"github.com/gilperopiola/lyfe-companyon-backend/utils"
	"github.com/stretchr/testify/assert"
)

func TestCreateTagController(t *testing.T) {
	cfg.Setup("testing")
	db.Setup(cfg)
	defer db.Close()
	rtr.Setup()
	token := generateTestingToken()

	tag := &Tag{
		Name:           "name",
		PrimaryColor:   "primaryColor",
		SecondaryColor: "secondaryColor",
	}

	response := tag.GenerateTestRequest(token, "POST", "")
	json.Unmarshal(response.Body.Bytes(), &tag)
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "name", tag.Name)
}

func TestGetTagController(t *testing.T) {
	cfg.Setup("testing")
	db.Setup(cfg)
	defer db.Close()
	rtr.Setup()
	token := generateTestingToken()

	tag := &Tag{
		Name:           "name",
		PrimaryColor:   "primaryColor",
		SecondaryColor: "secondaryColor",
	}
	tag, _ = tag.Create()

	response := tag.GenerateTestRequest(token, "GET", "/"+strconv.Itoa(tag.ID))
	json.Unmarshal(response.Body.Bytes(), &tag)
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "name", tag.Name)
}

func TestUpdateTagController(t *testing.T) {
	cfg.Setup("testing")
	db.Setup(cfg)
	defer db.Close()
	rtr.Setup()
	token := generateTestingToken()

	tag := &Tag{
		Name:           "name",
		PrimaryColor:   "primaryColor",
		SecondaryColor: "secondaryColor",
	}
	tag, _ = tag.Create()

	tag.Name = "name2"
	tag.Enabled = !tag.Enabled

	response := tag.GenerateTestRequest(token, "PUT", "/"+strconv.Itoa(tag.ID))
	json.Unmarshal(response.Body.Bytes(), &tag)
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "name2", tag.Name)
	assert.False(t, tag.Enabled)
}

func TestSearchTagController(t *testing.T) {
	cfg.Setup("testing")
	db.Setup(cfg)
	defer db.Close()
	rtr.Setup()
	token := generateTestingToken()

	tag := &Tag{}
	for i := 15; i <= 25; i++ {
		tag = &Tag{
			Name:           "name" + utils.ToString(i),
			PrimaryColor:   "primaryColor" + utils.ToString(i),
			SecondaryColor: "secondaryColor" + utils.ToString(i),
		}
		tag, _ = tag.Create()
	}

	params := &SearchParameters{
		Filter:        "name1",
		SortField:     "id",
		SortDirection: "DESC",
		ShowPrivate:   true,
		Limit:         3,
		Offset:        1,
	}

	tags := make([]*Tag, 0)
	response := tag.GenerateTestRequest(token, "GET", getSearchURL(params))
	json.Unmarshal(response.Body.Bytes(), &tags)
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, params.Limit, len(tags))
	assert.NotZero(t, tags[0].ID)
	assert.Equal(t, "name18", tags[0].Name)
	assert.Equal(t, "primaryColor18", tags[0].PrimaryColor)
	assert.Equal(t, "secondaryColor18", tags[0].SecondaryColor)
	assert.True(t, tags[0].Enabled)
}
