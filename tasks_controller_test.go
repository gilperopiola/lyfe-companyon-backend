package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"testing"

	"github.com/gilperopiola/lyfe-companyon-backend/utils"
	"github.com/stretchr/testify/assert"
)

func TestCreateTaskController(t *testing.T) {
	cfg.Setup("testing")
	db.Setup(cfg)
	defer db.Close()
	rtr.Setup()
	token := generateTestingToken()

	task := &Task{
		Name:       "name",
		Importance: 10,
		Tags:       createTestingTags(2),
	}

	response := task.GenerateTestRequest(token, "POST", "")
	json.Unmarshal(response.Body.Bytes(), &task)
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "name", task.Name)
}

func TestGetTaskController(t *testing.T) {
	cfg.Setup("testing")
	db.Setup(cfg)
	defer db.Close()
	rtr.Setup()
	token := generateTestingToken()

	task := &Task{
		Name:       "name",
		Importance: 10,
		Tags:       createTestingTags(2),
	}
	task, _ = task.Create()

	response := task.GenerateTestRequest(token, "GET", "/"+strconv.Itoa(task.ID))
	json.Unmarshal(response.Body.Bytes(), &task)
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "name", task.Name)
}

func TestUpdateTaskController(t *testing.T) {
	cfg.Setup("testing")
	db.Setup(cfg)
	defer db.Close()
	rtr.Setup()
	token := generateTestingToken()

	tags := createTestingTags(2)

	task := &Task{
		Name:       "name",
		Importance: 10,
		Duration:   Small,
		Tags:       tags,
	}
	task, _ = task.Create()

	task.Name = "name2"
	task.Status = Doing
	task.Duration = Medium
	task.Tags = []*Tag{tags[0]}

	response := task.GenerateTestRequest(token, "PUT", "/"+strconv.Itoa(task.ID))
	json.Unmarshal(response.Body.Bytes(), &task)
	log.Println(response.Body.String())
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "name2", task.Name)
	assert.Equal(t, Doing, task.Status)
	assert.Equal(t, Medium, task.Duration)
	assert.Equal(t, 1, len(task.Tags))
}

func TestSearchTaskController(t *testing.T) {
	cfg.Setup("testing")
	db.Setup(cfg)
	defer db.Close()
	rtr.Setup()
	token := generateTestingToken()

	tags := createTestingTags(2)

	task := &Task{}
	for i := 15; i <= 25; i++ {
		task = &Task{
			Name:       "name" + utils.ToString(i),
			Importance: 10,
			Tags:       tags,
		}
		task, _ = task.Create()
	}

	params := &SearchParameters{
		Filter:        "name1",
		SortField:     "id",
		SortDirection: "DESC",
		Limit:         3,
		Offset:        1,
	}

	tasks := make([]*Task, 0)
	response := task.GenerateTestRequest(token, "GET", getSearchURL(params))
	json.Unmarshal(response.Body.Bytes(), &tasks)
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, params.Limit, len(tasks))
	assert.NotZero(t, tasks[0].ID)
	assert.Equal(t, "name18", tasks[0].Name)
	assert.Equal(t, "name17", tasks[1].Name)
}
