package helpers

import (
	"net/http"
)

func SetHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}
