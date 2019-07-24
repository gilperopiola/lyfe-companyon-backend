package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"

	"github.com/gilperopiola/lyfe-companyon-backend/utils"
)

func (tag *Tag) GenerateTestRequest(token, method, url string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	body := tag.GetJSONBody()
	req, _ := http.NewRequest(method, "/Tags"+url, bytes.NewReader([]byte(body)))
	req.Header.Set("Authorization", token)
	rtr.ServeHTTP(w, req)
	return w
}

func (tag *Tag) GetJSONBody() string {
	body := `{
		"name": "` + tag.Name + `",
		"primaryColor": "` + tag.PrimaryColor + `",
		"secondaryColor": "` + tag.SecondaryColor + `",
		"enabled": ` + utils.BoolToString(tag.Enabled) + `
	}`
	return body
}
