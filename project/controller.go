package project

import (
	"encoding/json"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	results := FindAll()
	json.NewEncoder(w).Encode(results)
}
