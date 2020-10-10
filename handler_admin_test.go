package main

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// /
func TestIndex(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/", nil)

	res, err := app.Test(req, -1)
	assert.Nil(t, err)
	assert.Equal(t, res.StatusCode, 200)

	body, err := ioutil.ReadAll(res.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello World!", string(body))
}

// /ping
func TestPing(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/ping", nil)

	res, err := app.Test(req, -1)
	assert.Nil(t, err)
	assert.Equal(t, res.StatusCode, 200)

	body, err := ioutil.ReadAll(res.Body)
	assert.Nil(t, err)
	assert.Equal(t, "pong", string(body))
}

// /pid
func TestPid(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/admin/pid", nil)

	res, err := app.Test(req, -1)
	assert.Nil(t, err)
	assert.Equal(t, res.StatusCode, 200)

	body, err := ioutil.ReadAll(res.Body)
	assert.Nil(t, err)
	assert.Contains(t, string(body), "Worker #")
}

// /session
func TestSession(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/admin/session", nil)

	res, err := app.Test(req, -1)
	assert.Nil(t, err)
	assert.Equal(t, res.StatusCode, 200)

	body, err := ioutil.ReadAll(res.Body)
	assert.Nil(t, err)
	assert.Contains(t, string(body), "Session id: ")
}
