package main

import (
	"net/http"

	"github.com/gilperopiola/lyfe-companyon-backend/utils"
	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	var task *Task
	c.BindJSON(&task)

	if task.Name == "" || task.Importance == 0 {
		c.JSON(http.StatusBadRequest, "name and importance fields required")
		return
	}

	task, err := task.Create()
	if err != nil {
		c.JSON(http.StatusBadRequest, db.BeautifyError(err))
		return
	}

	c.JSON(http.StatusOK, task)
}

func GetTask(c *gin.Context) {
	task := &Task{ID: utils.ToInt(c.Param("id_task"))}

	task, err := task.Get()
	if err != nil {
		c.JSON(http.StatusBadRequest, db.BeautifyError(err))
		return
	}

	c.JSON(http.StatusOK, task)
}

func UpdateTask(c *gin.Context) {
	var task *Task
	c.BindJSON(&task)
	task.ID = utils.ToInt(c.Param("id_task"))

	if task.Name == "" || task.Importance == 0 {
		c.JSON(http.StatusBadRequest, "name and importance required")
		return
	}

	task, err := task.Update()
	if err != nil {
		c.JSON(http.StatusBadRequest, db.BeautifyError(err))
		return
	}

	c.JSON(http.StatusOK, task)
}

func SearchTasks(c *gin.Context) {
	task := &Task{}

	params := &SearchParameters{
		Filter:        c.Query("filter"),
		SortField:     c.Query("sortField"),
		SortDirection: c.Query("sortDirection"),
		Limit:         utils.ToInt(c.Query("limit")),
		Offset:        utils.ToInt(c.Query("offset")),
	}

	tasks, err := task.Search(params)
	if err != nil {
		c.JSON(http.StatusBadRequest, db.BeautifyError(err))
		return
	}

	c.JSON(http.StatusOK, tasks)
}