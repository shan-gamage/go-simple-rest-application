package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	port := ":8000"
	r := chi.NewRouter()

	//Get Last price
	r.Get("/getprice/last", getprice.lastPrice{})

	//Get price by timestamp
	r.Get("/getprice/bytimestamp", func(w http.ResponseWriter, r *http.Request) {
		type items struct {
			message string
		}
		i := items{"message"}
		json.NewEncoder(w).Encode(i)
	})

	//Get average price
	r.Get("/getprice/average", func(w http.ResponseWriter, r *http.Request) {
		type items struct {
			message string
		}
		i := items{"message"}
		json.NewEncoder(w).Encode(i)
	})
	fmt.Println("Serving on port: " + port)
	http.ListenAndServe(port, r)
}
