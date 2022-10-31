package main

import (
	f "fmt"
	"go-simple-rest-application/controller"
	"go-simple-rest-application/dbconnection"
	"go-simple-rest-application/model"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

func main() {

	var dbcon = dbconnection.Connector()
	for t := range time.Tick(60 * time.Second) {
		f.Printf("working %s \n", t)
		var resultobject = controller.Updateprice()
		result, err := model.Createrecord(dbcon, resultobject)
		if err != nil {
			log.Printf("Insert product failed with error %s", err)
			return
		}
		f.Println("Result ", result)
	}

	//

	port := ":8080"
	r := chi.NewRouter()
	r.Route("/getprice", func(r chi.Router) {
		r.Get("/last", controller.Lastprice(dbcon))
		r.Get("/bytimestamp", controller.Bytimestamp(dbcon))
		r.Get("/average", controller.Average(dbcon))
	})
	f.Println("Serving on port: " + port)
	http.ListenAndServe(port, r)
}
