package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	f "fmt"
	"go-simple-rest-application/model"
	"net/http"
	"strconv"
	"time"
)

func Lastprice(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		f.Println(r)
		if r.URL.Query().Get("type") != "" {
			f.Println("FOUND TYPE", r.URL.Query().Get("type"))
		}
		resutlData, err := model.Getlastrecord(db)
		if err != nil {
			return
		}
		f.Println("result: ", resutlData)
		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(resutlData)
		}
		w.WriteHeader(http.StatusOK)
		s, _ := json.MarshalIndent(resutlData, "", "\t")
		json.NewEncoder(w).Encode(s)
	}
}

func Bytimestamp(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("timestamp") != "" {
			f.Println("FOUND TIMESTAMP", r.URL.Query().Get("timestamp"))
			var value = r.URL.Query().Get("timestamp")
			timeStamp, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
			// searchDate := "2022-10-28 21:28:30"
			timestampValue := time.Unix(timeStamp, 0).Format("2006-01-02 15:04:05")
			fmt.Println("timestampValue", timestampValue)
			resutlData, err := model.SearchByTimestamp(db, string(timestampValue))
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
			f.Println("result: ", resutlData)
			w.Header().Set("Content-Type", "application/json")
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(resutlData)
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(resutlData)
		} else {
			// w.WriteHeader(http.StatusInternalServerError)
			// json.NewEncoder(w).Encode(resutlData)
		}
	}
}

func Average(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("starttime") != "" && r.URL.Query().Get("endtime") != "" {
			f.Println("FOUND TIMESTAMP starttime", r.URL.Query().Get("starttime"))
			f.Println("FOUND TIMESTAMP endtime", r.URL.Query().Get("endtime"))
			var valueStarttime = r.URL.Query().Get("starttime")
			var valueEndtime = r.URL.Query().Get("endtime")
			timeStampStart, err := strconv.ParseInt(valueStarttime, 10, 64)
			timeStampEnd, err := strconv.ParseInt(valueEndtime, 10, 64)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
			startDate := string(time.Unix(timeStampStart, 0).Format("2006-01-02 15:04:05"))
			endDate := string(time.Unix(timeStampEnd, 0).Format("2006-01-02 15:04:05"))
			fmt.Println("startDate", startDate)
			fmt.Println("endDate", endDate)
			resutlData, err := model.GetAverage(db, startDate, endDate)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
			f.Println("result: ", resutlData)
			w.Header().Set("Content-Type", "application/json")
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(resutlData)
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(resutlData)
		}
	}
}
