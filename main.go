package main

import (
	f "fmt"
	"go-simple-rest-application/controller"
	"go-simple-rest-application/dbconnection"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {

	var dbcon = dbconnection.Connector()
	// for t := range time.Tick(60 * time.Second) {
	// 	fmt.Printf("working %s \n", t)
	// 	var resultobject = controller.Updateprice()
	// 	result, err := model.Createrecord(dbcon, resultobject)
	// 	if err != nil {
	// 		log.Printf("Insert product failed with error %s", err)
	// 		return
	// 	}
	// 	f.Println("Result ", result)
	// }

	//

	port := ":8002"
	r := chi.NewRouter()

	//Get Last price
	r.Get("/getprice/last", controller.Lastprice(dbcon))

	//Get price by timestamp
	r.Get("/getprice/bytimestamp", controller.Bytimestamp(dbcon))

	//Get average price
	r.Get("/getprice/average", controller.Average(dbcon))
	f.Println("Serving on port: " + port)
	http.ListenAndServe(port, r)
}
