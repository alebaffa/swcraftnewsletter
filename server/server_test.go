package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alebaffa/newsletter-web/routing"
	"github.com/stretchr/testify/assert"
)

func TestHomePageHandler(t *testing.T) {
	assert := assert.New(t)

	mux := routing.NewMux()

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	mux.ServeHTTP(res, req)

	assert.Equal(res.Code, 200)
}
