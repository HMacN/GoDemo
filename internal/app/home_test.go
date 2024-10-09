package app

import (
	"GoDemo/internal/assert"
	"bytes"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApplicationHome_ReturnsOk(t *testing.T) {
	// Arrange
	responseRecorder := httptest.NewRecorder()
	app := NewApp()

	// Act
	request, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	app.Home(responseRecorder, request)
	result := responseRecorder.Result()

	// Assert
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

	filePaths := []string{
		app.TemplateBasePath,
		app.PartialsNavPath,
		app.HomePagePath,
	}

	templateSet, err := template.ParseFiles(filePaths...)
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	buffer := new(bytes.Buffer)
	err = templateSet.ExecuteTemplate(buffer, "base", nil)
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	expected := buffer.String()

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
	assert.Equal(t, string(expected), string(result))
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
