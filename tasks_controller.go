package main

import (
	"log"
	"net/http"
	"time"

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

	if task.Name == "" || task.Importance == 0 || task.Duration == 0 {
		log.Printf("%v", task)
		c.JSON(http.StatusBadRequest, "all fields required")
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
		Filter:           c.Query("filter"),
		FilterTagID:      utils.ToInt(c.Query("filterTagID")),
		FilterImportance: utils.ToInt(c.Query("filterImportance")),
		SortField:        c.Query("sortField"),
		SortDirection:    c.Query("sortDirection"),
		ShowPrivate:      utils.ToBool(c.Query("showPrivate")),
		Limit:            utils.ToInt(c.Query("limit")),
		Offset:           utils.ToInt(c.Query("offset")),
	}

	tasks, err := task.Search(params)
	if err != nil {
		c.JSON(http.StatusBadRequest, db.BeautifyError(err))
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func GetWeeklyDoneAndArchivedTasks(c *gin.Context) {
	task := &Task{}

	tasks, err := task.GetDoneAndArchivedSince(time.Now().Add(-24 * 7 * time.Hour))
	if err != nil {
		c.JSON(http.StatusBadRequest, db.BeautifyError(err))
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func GetMonthlyDoneAndArchivedTasks(c *gin.Context) {
	task := &Task{}

	tasks, err := task.GetDoneAndArchivedSince(time.Now().Add(-24 * 31 * time.Hour))
	if err != nil {
		c.JSON(http.StatusBadRequest, db.BeautifyError(err))
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func SetTaskDoing(c *gin.Context) {
	task := &Task{ID: utils.ToInt(c.Param("id_task")), Status: Doing}

	task, err := task.UpdateStatus()
	if err != nil {
		c.JSON(http.StatusBadRequest, db.BeautifyError(err))
		return
	}

	c.JSON(http.StatusOK, task)
}

func CompleteTask(c *gin.Context) {
	task := &Task{ID: utils.ToInt(c.Param("id_task")), Status: Done}

	task, err := task.UpdateStatus()
	if err != nil {
		c.JSON(http.StatusBadRequest, db.BeautifyError(err))
		return
	}

	c.JSON(http.StatusOK, task)
}

func ArchiveTask(c *gin.Context) {
	task := &Task{ID: utils.ToInt(c.Param("id_task")), Status: Archived}

	task, err := task.UpdateStatus()
	if err != nil {
		c.JSON(http.StatusBadRequest, db.BeautifyError(err))
		return
	}

	c.JSON(http.StatusOK, task)
}
