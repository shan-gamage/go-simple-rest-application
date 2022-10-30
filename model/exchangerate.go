package model

import (
	"context"
	"database/sql"
	"encoding/json"
	f "fmt"
	"log"
	"time"
)

func Createrecord(db *sql.DB, rateobject string) (int64, error) {
	type Bird struct {
		Asset_id_base  string
		Asset_id_quote string
		Time           string
		Rate           float64
	}
	var bird Bird
	json.Unmarshal([]byte(rateobject), &bird)

	result, err := db.Exec("INSERT INTO exchange_rate(assest_type, exchange_type, rate) VALUES (?,?,?)", bird.Asset_id_base, bird.Asset_id_quote, bird.Rate)
	if err != nil {
		return 0, f.Errorf("AddAlbum: %v", err)
	}

	// Get the new album's generated ID for the client.
	id, err := result.LastInsertId()
	if err != nil {
		return 0, f.Errorf("AddAlbum: %v", err)
	}
	// Return the new album's ID.
	return id, nil

}

func Getlastrecord(db *sql.DB) (string, error) {
	query := "SELECT rate, created_at FROM exchange_rate ORDER BY id DESC LIMIT 1;"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	messsage := "SUCCESS"
	var rate float64
	var createdAt string
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		rate = 0.00
		createdAt = ""
		messsage = err.Error()
	}
	defer stmt.Close()
	type ResultData struct {
		Rate      float64
		CreatedAt string
		Message   string
	}
	row := stmt.QueryRowContext(ctx)
	if err := row.Scan(&rate, &createdAt); err != nil {
		f.Println("err", err)
		rate = 0.00
		createdAt = ""
		messsage = err.Error()
	}
	f.Println("rate", rate)
	f.Println("createdAt", createdAt)
	finalResult := ResultData{rate, createdAt, messsage}
	f.Println(finalResult)
	jsonData, err := json.Marshal(finalResult)
	f.Println((jsonData))
	return string(jsonData), nil
}

func SearchByTimestamp(db *sql.DB, searchdata string) (string, error) {
	messsage := "SUCCESS"
	var rate float64
	var nearestRate float64
	var createdAt string
	type ResultData struct {
		Rate        float64
		CreatedAt   string
		Message     string
		NearestRate float64
	}
	rows, err := db.Query("SELECT rate, created_at FROM exchange_rate WHERE created_at = ?", searchdata)
	if err != nil {
		f.Println("Searched value", err)
		rate = 0.00
		nearestRate = 0.00
		createdAt = ""
		messsage = err.Error()
	}
	defer rows.Close()
	if rows.Next() {
		if err := rows.Scan(&rate, &createdAt); err != nil {
			f.Println("ERROR finding value", err)
			rate = 0.00
			nearestRate = 0.00
			createdAt = ""
			messsage = err.Error()
		}
		f.Println("Exact match")
		f.Println("Rate : ", rate)
		f.Println("Created at: ", createdAt)
	} else {
		rows, err := db.Query("SELECT rate, created_at FROM exchange_rate ORDER BY ABS(UNIX_TIMESTAMP(created_at) - UNIX_TIMESTAMP(?))  LIMIT 1", searchdata)
		if err != nil {
			f.Println("Searched value", err)
			rate = 0.00
			nearestRate = 0.00
			createdAt = ""
			messsage = err.Error()
		}
		if rows.Next() {
			if err := rows.Scan(&nearestRate, &createdAt); err != nil {
				f.Println("ERROR finding value", err)
			}
			f.Println("Nearest match")
			f.Println("Rate : ", rate)
			f.Println("Created at: ", createdAt)
		}
	}
	finalResult := ResultData{rate, createdAt, messsage, nearestRate}
	jsonData, err := json.Marshal(finalResult)
	f.Println((jsonData))
	return string(jsonData), nil
}

func GetAverage(db *sql.DB, startDate string, endDate string) (string, error) {
	messsage := "SUCCESS"
	var rateValue float64
	type ResultData struct {
		Rate    float64
		Message string
	}
	rows, err := db.Query("SELECT AVG(rate) FROM exchange_rate WHERE created_at BETWEEN ?  AND ?", startDate, endDate)
	if err != nil {
		f.Println("Searched value", err)
		rateValue = 0.00
		messsage = err.Error()
	}
	defer rows.Close()
	if rows.Next() {
		if err := rows.Scan(&rateValue); err != nil {
			f.Println("ERROR finding value", err)
			rateValue = 0.00
			messsage = err.Error()
		}
		f.Println("Exact match")
		f.Println("Average Rate : ", rateValue)
	}
	finalResult := ResultData{rateValue, messsage}
	jsonData, err := json.Marshal(finalResult)
	f.Println((jsonData))
	return string(jsonData), nil
}
