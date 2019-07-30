package main

import (
	"net/http/httptest"
)

type Tag struct {
	ID             int
	Name           string
	PrimaryColor   string
	SecondaryColor string
	Enabled        bool
}

type TagActions interface {
	Create() (*Tag, error)
	Get() (*Tag, error)
	Update() (*Tag, error)
	Search(params *SearchParameters) ([]*Tag, error)

	GetTasks() (*[]Task, error)
}

type TagTestingActions interface {
	GenerateTestRequest(token, method, url string) *httptest.ResponseRecorder
	GenerateTestJSONBody() string
}
