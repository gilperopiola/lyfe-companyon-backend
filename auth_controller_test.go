package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSignupController(t *testing.T) {
	cfg.Setup("testing")
	db.Setup(cfg)
	defer db.Close()
	rtr.Setup()

	w := httptest.NewRecorder()
	body := `{"email": "email", "password": "password", "repeatPassword": "password"}`
	req, _ := http.NewRequest("POST", "/Signup", bytes.NewReader([]byte(body)))
	rtr.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestLoginController(t *testing.T) {
	cfg.Setup("testing")
	db.Setup(cfg)
	defer db.Close()
	rtr.Setup()

	user := &User{
		Email:    "email",
		Password: hash("email", "password"),
	}
	user, _ = user.Create()

	w := httptest.NewRecorder()
	body := `{"email": "email", "password": "password"}`
	req, _ := http.NewRequest("POST", "/Login", bytes.NewReader([]byte(body)))
	rtr.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
