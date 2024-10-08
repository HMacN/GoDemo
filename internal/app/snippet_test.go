package app

import (
	"GoDemo/internal/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestSnippet_ReturnsSnippetId(t *testing.T) {
	// Arrange
	app := NewApp()
	testId := 1
	responseRecorder := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, "/?id="+strconv.Itoa(testId), nil)
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	// Act
	app.Snippet(responseRecorder, request)
	result := responseRecorder.Result()
	defer result.Body.Close()
	body, err := io.ReadAll(result.Body)
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	// Assert
	assert.Equal(t, strconv.Itoa(testId), string(body))
}
