package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alebaffa/swcraftnewsletter/mail"
	"github.com/alebaffa/swcraftnewsletter/routing"
	"github.com/stretchr/testify/assert"
)

func TestHomePageHandler(t *testing.T) {
	assert := assert.New(t)
	res := httptest.NewRecorder()
	mux := routing.NewMux()

	req, _ := http.NewRequest("GET", "/", nil)
	mux.ServeHTTP(res, req)

	assert.Equal(res.Code, 200)
}

func ThankyouPageHandler(t *testing.T) {
	assert := assert.New(t)
	res := httptest.NewRecorder()
	mux := routing.NewMux()

	req, _ := http.NewRequest("GET", "/thankyou", nil)
	mux.ServeHTTP(res, req)

	assert.Equal(res.Code, 200)
}

func TestConfig(t *testing.T) {
	error := mail.Send("test")
	if error != nil {
		assert.Fail(t, "Error")
	}
}
