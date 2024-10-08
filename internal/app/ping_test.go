package app

import (
	"GoDemo/internal/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestApplicationPing_ReturnsStatusOk(t *testing.T) {
	// Arrange
	app := NewApp()
	responseRecorder := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, "", nil)
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	// Act
	app.Ping(responseRecorder, request)
	result := responseRecorder.Result()

	// Assert
	assert.Equal(t, http.StatusOK, result.StatusCode)
}

func TestApplicationPing_ReturnsMessagePong(t *testing.T) {
	// Arrange
	app := NewApp()
	responseRecorder := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, "", nil)
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	// Act
	app.Ping(responseRecorder, request)
	result := responseRecorder.Result()
	defer result.Body.Close()
	body, err := io.ReadAll(result.Body)
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	// Assert
	assert.Equal(t, "pong", string(body))
}

func TestApplicationPing_OnlyAllowsGet(t *testing.T) {
	//Arrange
	tests := []struct {
		method string
		expect int
	}{
		{
			method: http.MethodGet,
			expect: http.StatusOK,
		},
		{
			method: http.MethodConnect,
			expect: http.StatusMethodNotAllowed,
		},
		{
			method: http.MethodDelete,
			expect: http.StatusMethodNotAllowed,
		},
		{
			method: http.MethodConnect,
			expect: http.StatusMethodNotAllowed,
		},
		{
			method: http.MethodPost,
			expect: http.StatusMethodNotAllowed,
		},
		{
			method: http.MethodPut,
			expect: http.StatusMethodNotAllowed,
		},
		{
			method: http.MethodOptions,
			expect: http.StatusMethodNotAllowed,
		},
	}

	// Act
	for _, test := range tests {
		request, err := http.NewRequest(test.method, "", nil)
		responseRecorder := httptest.NewRecorder()
		if err != nil {
			assert.Fail(t, err.Error())
			return
		}

		app := NewApp()
		app.Ping(responseRecorder, request)

		// Assert
		assert.Equal(t, test.expect, responseRecorder.Code)
	}
}

func TestApplicationPing_MethodNotAllowedText(t *testing.T) {
	// Arrange
	app := NewApp()
	responseRecorder := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodPost, "", nil)
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	// Act
	app.Ping(responseRecorder, request)
	result := responseRecorder.Result()
	defer result.Body.Close()
	body, err := io.ReadAll(result.Body)
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	// Assert
	assert.Equal(t, "Method Not Allowed", strings.Trim(string(body), "\n"))
}

func TestApplicationPing_MethodNotAllowedHeader(t *testing.T) {
	// Arrange
	app := NewApp()
	responseRecorder := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodPost, "", nil)
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	// Act
	app.Ping(responseRecorder, request)
	result := responseRecorder.Result().Header.Get("Allow")

	// Assert
	assert.Equal(t, http.MethodGet, result)
}
