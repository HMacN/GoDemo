package handlers

import (
	"net/http"
	"strconv"
)

func Snippet(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(strconv.Itoa(id)))
}
