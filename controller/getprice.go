package controller

import (
	"encoding/json"
	"net/http"
)

func Lastprice() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type items struct {
			message string
		}
		i := items{"message"}
		json.NewEncoder(w).Encode(i)
	}
}

func Bytimestamp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type items struct {
			message string
		}
		i := items{"message"}
		json.NewEncoder(w).Encode(i)
	}
}

func Average() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type items struct {
			message string
		}
		i := items{"message"}
		json.NewEncoder(w).Encode(i)
	}
}
