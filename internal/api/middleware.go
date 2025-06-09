package api

import (
	"net/http"
	"strconv"
)

func ValidateTaskID(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		if _, err := strconv.ParseInt(id, 10, 64); err != nil {
			http.Error(w, "Invalid task ID: must be numeric", http.StatusBadRequest)
			return
		}
		next(w, r)
	}
}
