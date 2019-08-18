package main

import (
	"net/http/httptest"
	"time"
)

type Task struct {
	ID          int
	Name        string
	Description string
	Importance  int
	Status      TaskStatus
	Duration    TaskDuration
	DueDate     time.Time
	Tags        []*Tag
	DateCreated time.Time
}

type TaskActions interface {
	Create() (*Task, error)
	Get() (*Task, error)
	Update() (*Task, error)
	UpdateStatus() (*Task, error)
	Search(params *SearchParameters) ([]*Task, error)

	createTags() ([]*Tag, error)
	getTags() ([]*Tag, error)
	updateTags() ([]*Tag, error)
}

type TaskTestingActions interface {
	GenerateTestRequest(token, method, url string) *httptest.ResponseRecorder
	GenerateTestJSONBody() string
}

type TaskStatus int

const (
	Pending  TaskStatus = 1
	Doing    TaskStatus = 2
	Done     TaskStatus = 3
	Archived TaskStatus = 4
)

type TaskDuration int

const (
	ExtraSmall TaskDuration = 1
	Small      TaskDuration = 2
	Medium     TaskDuration = 3
	Large      TaskDuration = 4
	ExtraLarge TaskDuration = 5
)
