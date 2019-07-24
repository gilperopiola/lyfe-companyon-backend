package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"

	"github.com/gilperopiola/lyfe-companyon-backend/utils"
)

func (task *Task) GenerateTestRequest(token, method, url string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	body := task.GetJSONBody()
	req, _ := http.NewRequest(method, "/Tasks"+url, bytes.NewReader([]byte(body)))
	req.Header.Set("Authorization", token)
	rtr.ServeHTTP(w, req)
	return w
}

func (task *Task) GetJSONBody() string {
	tagsString := ""
	for key, tag := range task.Tags {
		tagsString += `{"id": ` + utils.ToString(tag.ID) + `}`

		if key+1 != len(task.Tags) {
			tagsString += ", "
		}
	}

	body := `{
		"name": "` + task.Name + `",
		"importance": ` + utils.ToString(task.Importance) + `,
		"status": ` + utils.ToString(int(task.Status)) + `,
		"tags": [` + tagsString + `]
	}`
	return body
}
