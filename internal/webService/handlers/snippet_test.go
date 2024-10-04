package handlers

import (
	"GoDemo/internal/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestSnippet_ReturnsSnippetId(t *testing.T) {
	testId := 1
	responseRecorder := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, "/?id="+strconv.Itoa(testId), nil)
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	Snippet(responseRecorder, request)
	result := responseRecorder.Result()
	defer result.Body.Close()
	body, err := io.ReadAll(result.Body)
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	assert.Equal(t, strconv.Itoa(testId), string(body))
}
