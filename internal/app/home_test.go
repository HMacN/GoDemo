package app

import (
	"GoDemo/internal/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHome(t *testing.T) {

}

func TestApplicationHome_CatchAll(t *testing.T) {
	responseRecorder := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, "/unauthorised/url", nil)
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	app := NewApp()
	app.Home(responseRecorder, request)
	result := responseRecorder.Result()

	assert.Equal(t, http.StatusNotFound, result.StatusCode)
}
