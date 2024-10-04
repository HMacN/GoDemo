package webService

import (
	"GoDemo/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPing_ReturnsStatusOk(t *testing.T) {
	responseRecorder := httptest.NewRecorder()

	request, err := http.NewRequest(http.MethodGet, "/ping", nil)
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	handler(responseRecorder, request)
	result := responseRecorder.Result()

	assert.Equal(t, result.StatusCode, http.StatusOK)
}

func TestPing_ReturnsMessagePong(t *testing.T) {
	responseRecorder := httptest.NewRecorder()

	request, err := http.NewRequest(http.MethodGet, "/ping", nil)
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	handler(responseRecorder, request)
	result := responseRecorder.Result()
	defer result.Body.Close()
	body, err := io.ReadAll(result.Body)
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	assert.Equal(t, string(body), "pong")
}
