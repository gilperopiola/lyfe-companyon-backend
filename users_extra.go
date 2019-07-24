package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"

	"github.com/gilperopiola/lyfe-companyon-backend/utils"
)

func (user *User) GenerateTestRequest(token, method, url string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	body := user.GetJSONBody()
	req, _ := http.NewRequest(method, "/Users"+url, bytes.NewReader([]byte(body)))
	req.Header.Set("Authorization", token)
	rtr.ServeHTTP(w, req)
	return w
}

func (user *User) GetJSONBody() string {
	body := `{
		"email": "` + user.Email + `",
		"password": "` + user.Password + `",
		"firstName": "` + user.FirstName + `",
		"lastName": "` + user.LastName + `",
		"enabled": ` + utils.BoolToString(user.Enabled) + `
	}`
	return body
}
