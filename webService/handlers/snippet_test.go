package handlers

import (
	"GoDemo/assert"
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
	result := responseRecorder.Result().Header.Get("Id")

	assert.Equal(t, strconv.Itoa(testId), result)
}
