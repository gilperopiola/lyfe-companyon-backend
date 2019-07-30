package main

import (
	"testing"

	"github.com/gilperopiola/lyfe-companyon-backend/utils"
	"github.com/stretchr/testify/assert"
)

func TestCreateTask(t *testing.T) {
	cfg.Setup("testing")
	db.Setup(cfg)
	defer db.Close()

	task := &Task{
		Name:        "name",
		Description: "description",
		Importance:  10,
		Tags:        createTestingTags(2),
	}

	task, err := task.Create()

	assert.NoError(t, err)
	assert.NotZero(t, task.ID)
	assert.Equal(t, "name", task.Name)
	assert.Equal(t, "description", task.Description)
	assert.Equal(t, 10, task.Importance)
	assert.Equal(t, Pending, task.Status)
	assert.Equal(t, 2, len(task.Tags))
	assert.NotZero(t, task.DateCreated)
}

func TestGetTask(t *testing.T) {
	cfg.Setup("testing")
	db.Setup(cfg)
	defer db.Close()

	task := &Task{
		Name:       "name",
		Importance: 10,
		Tags:       createTestingTags(2),
	}

	task, _ = task.Create()
	task, err := task.Get()

	assert.NoError(t, err)
	assert.NotZero(t, task.ID)
	assert.Equal(t, "name", task.Name)
	assert.Equal(t, 10, task.Importance)
	assert.Equal(t, Pending, task.Status)
	assert.Equal(t, 2, len(task.Tags))
}

func TestUpdateTask(t *testing.T) {
	cfg.Setup("testing")
	db.Setup(cfg)
	defer db.Close()

	tags := createTestingTags(2)

	task := &Task{
		Name:       "name",
		Importance: 10,
		Tags:       tags,
	}

	task, _ = task.Create()

	task.Name = "name2"
	task.Importance = 1
	task.Status = Doing
	task.Tags = []*Tag{tags[0]}

	task, err := task.Update()
	assert.NoError(t, err)
	assert.NotZero(t, task.ID)
	assert.Equal(t, "name2", task.Name)
	assert.Equal(t, 1, task.Importance)
	assert.Equal(t, Doing, task.Status)
	assert.Equal(t, 1, len(task.Tags))
}

func TestSearchTasks(t *testing.T) {
	cfg.Setup("testing")
	db.Setup(cfg)
	defer db.Close()

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

	tasks, err := task.Search(params)
	assert.NoError(t, err)

	assert.Equal(t, params.Limit, len(tasks))
	assert.NotZero(t, tasks[0].ID)
	assert.Equal(t, "name18", tasks[0].Name)
	assert.Equal(t, "name17", tasks[1].Name)
}
