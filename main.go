package main

import (
	"fmt"
	"go-simple-rest-application/controller"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	port := ":8000"
	r := chi.NewRouter()

	//Get Last price
	r.Get("/getprice/last", controller.Lastprice())

	//Get price by timestamp
	r.Get("/getprice/bytimestamp", controller.Bytimestamp())

	//Get average price
	r.Get("/getprice/average", controller.Average())
	fmt.Println("Serving on port: " + port)
	http.ListenAndServe(port, r)
}
