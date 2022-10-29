package controller

import (
	"database/sql"
	"encoding/json"
	f "fmt"
	"go-simple-rest-application/model"
	"net/http"
)

func Lastprice(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		f.Println(r)
		resutlData, err := model.Getlastrecord(db)
		if err != nil {
			return
		}
		json.NewEncoder(w).Encode(resutlData)
	}
}

func Bytimestamp(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		searchDate := "2022-10-28 21:28:30"
		model.SearchByTimestamp(db, searchDate)
		type items struct {
			message string
		}
		i := items{"message"}
		json.NewEncoder(w).Encode(i)
	}
}

func Average(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startDate := "2022-10-28 21:29:33"
		endDate := "2022-10-28 21:32:33"
		model.GetAverage(db, startDate, endDate)
		type items struct {
			message string
		}
		i := items{"message"}
		json.NewEncoder(w).Encode(i)
	}
}
