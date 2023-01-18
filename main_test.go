package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPingRouter(t *testing.T) {
	router := router()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, w.Body.String(), "{\"msg\":\"pong\"}")
}

func TestPs(t *testing.T) {
	router := router()
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	body := bytes.NewBufferString("{\"name\":\"foo\"}")
	c.Request, _ = http.NewRequest("POST", "/ps", body)
	router.ServeHTTP(w, c.Request)

	assert.JSONEq(t, w.Body.String(), "{\"msg\":{\"name\":\"foo\"}}")
	assert.Equal(t, w.Code, http.StatusOK)
}

func TestPs2(t *testing.T) {
	router := router()
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	body := bytes.NewBufferString("{\"name\":\"\"}")
	c.Request, _ = http.NewRequest("POST", "/ps", body)
	router.ServeHTTP(w, c.Request)

	assert.JSONEq(t, w.Body.String(), "{\"msg\":\"error\"}")
	assert.Equal(t, w.Code, 400)
}
