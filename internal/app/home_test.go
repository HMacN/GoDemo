package app

import (
	"GoDemo/internal/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHome_ReturnsOk(t *testing.T) {
	responseRecorder := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	app := NewApp()
	app.Home(responseRecorder, request)
	result := responseRecorder.Result()

	assert.Equal(t, http.StatusOK, result.StatusCode)
}

func TestApplicationHome_CatchAllUnauthorised(t *testing.T) {
	responseRecorder := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, "/unauthorised/url", nil)
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	app := NewApp()
	app.Home(responseRecorder, request)
	result := responseRecorder.Result()

	assert.Equal(t, http.StatusNotFound, result.StatusCode)
}
