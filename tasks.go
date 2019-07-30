package main

import (
	"net/http/httptest"
)

type Task struct {
	ID         int
	Name       string
	Importance int
	Status     TaskStatus
	Tags       []*Tag
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
