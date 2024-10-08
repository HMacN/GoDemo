package app

import (
	"GoDemo/internal/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestApplicationStatic_ServesFile(t *testing.T) {
	testUrl := "http://localhost:2000/static/testing/testing.txt"
	app := NewApp()
	mux := *app.Routes()

	responseRecorder := httptest.NewRecorder()
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
	assert.Equal(t, expected, result)
}

func TestApplicationStatic_StopsDirectoryTraversal(t *testing.T) {
	testUrl := "http://localhost:2000/static/../.."
	app := NewApp()
	mux := *app.Routes()

	responseRecorder := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, testUrl, nil)
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	mux.ServeHTTP(responseRecorder, request)
	result := responseRecorder.Result().StatusCode
	expected := http.StatusMovedPermanently
	assert.Equal(t, expected, result)
}
