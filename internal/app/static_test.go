package app

import (
	"GoDemo/internal/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestApplicationStatic_ServesFile(t *testing.T) {
	// Arrange
	testUrl := "http://localhost:2000/static/testing/testing.txt"
	app := NewApp()
	mux := *app.Routes()
	responseRecorder := httptest.NewRecorder()

	// Act
	request, err := http.NewRequest(http.MethodGet, testUrl, nil)
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	testFileContent, err := os.ReadFile(app.StaticPath + "testing/testing.txt")
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	expected := string(testFileContent)
	mux.ServeHTTP(responseRecorder, request)
	result := responseRecorder.Body.String()

	// Assert
	assert.Equal(t, expected, result)
}

func TestApplicationStatic_StopsDirectoryTraversal(t *testing.T) {
	// Arrange
	testUrl := "http://localhost:2000/static/../.."
	app := NewApp()
	mux := *app.Routes()
	responseRecorder := httptest.NewRecorder()

	// Act
	request, err := http.NewRequest(http.MethodGet, testUrl, nil)
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	mux.ServeHTTP(responseRecorder, request)
	result := responseRecorder.Result().StatusCode
	expected := http.StatusMovedPermanently

	// Assert
	assert.Equal(t, expected, result)
}

func TestApplicationStatic_EmptyDirReturnsNotFound(t *testing.T) {
	// Arrange
	testUrl := "http://localhost:2000/static/testing/empty/"
	app := NewApp()
	mux := *app.Routes()
	responseRecorder := httptest.NewRecorder()

	// Act
	request, err := http.NewRequest(http.MethodGet, testUrl, nil)
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	mux.ServeHTTP(responseRecorder, request)
	result := responseRecorder.Result().StatusCode
	expected := http.StatusNotFound

	// Assert
	assert.Equal(t, expected, result)
}
