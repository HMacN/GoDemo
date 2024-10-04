package handlers

import (
	"GoDemo/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSnippet_ReturnsSnippetId(t *testing.T) {
	responseRecorder := httptest.NewRecorder()

	request, err := http.NewRequest(http.MethodGet, "/view?id=1", nil)
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	Snippet(responseRecorder, request)
}
