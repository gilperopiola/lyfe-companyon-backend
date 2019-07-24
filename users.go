package main

import (
	"net/http/httptest"
	"time"
)

type User struct {
	ID          int
	Email       string
	Password    string
	FirstName   string
	LastName    string
	Enabled     bool
	DateCreated time.Time

	Token string
}

type UserActions interface {
	Login() (*User, error)

	Create() (*User, error)
	Get() (*User, error)
	Update() (*User, error)
	Search(params *SearchParameters) ([]*User, error)
}

type UserTestingActions interface {
	GenerateTestRequest(token, method, url string) *httptest.ResponseRecorder
	GenerateTestJSONBody() string
}
