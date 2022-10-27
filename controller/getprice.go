package controller

import {
	"net/http"
	"encoding/json"
}


func lastPrice() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type items struct {
			message string
		json.NewEncoder(w).Encode(i)
	}
}