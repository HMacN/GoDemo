package app

import (
	"GoDemo/internal/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestApplicationHome_ReturnsOk(t *testing.T) {
	// Arrange
	responseRecorder := httptest.NewRecorder()
	app := NewApp()
	srv := app.Routes()

	request, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	// Act
	srv.ServeHTTP(responseRecorder, request)

	// Assert
	result := responseRecorder.Result()
	assert.Equal(t, http.StatusOK, result.StatusCode)
}

func TestApplicationHome_ReturnsHomepage(t *testing.T) {
	// Arrange
	app := NewApp()
	responseRecorder := httptest.NewRecorder()

	request, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	// Act
	app.Home(responseRecorder, request)
	response := responseRecorder.Result()
	defer response.Body.Close()

	result, err := io.ReadAll(response.Body)
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	// Assert
	assert.True(t, strings.Contains(string(result), "<!doctype html>"))
}

func TestApplicationHome_CatchAllUnauthorised(t *testing.T) {
	// Arrange
	app := NewApp()
	responseRecorder := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, "/unauthorised/url", nil)
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	// Act
	app.Home(responseRecorder, request)
	result := responseRecorder.Result()

	// Assert
	assert.Equal(t, http.StatusNotFound, result.StatusCode)
}
