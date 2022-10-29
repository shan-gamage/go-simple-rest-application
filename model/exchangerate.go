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
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return "", err
	}
	defer stmt.Close()
	type ResultData struct {
		Rate      float64
		CreatedAt string
	}

	var rate float64
	var createdAt string
	row := stmt.QueryRowContext(ctx)
	if err := row.Scan(&rate, &createdAt); err != nil {
		f.Println("err", err)
		return "", err
	}
	f.Println("rate", rate)
	f.Println("createdAt", createdAt)
	finalResult := ResultData{rate, createdAt}
	json_data, err := json.Marshal(finalResult)
	f.Println((json_data))
	return string(json_data), nil
}

func SearchByTimestamp(db *sql.DB, searchdata string) error {
	// var price float64
	// var createdAt string
	rows, err := db.Query("SELECT id, rate, created_at FROM exchange_rate WHERE created_at = ?", searchdata)
	if err != nil {
		f.Println("Searched value", err)
		return err
	}
	defer rows.Close()
	var price float64
	var id int
	var createdDate string
	if rows.Next() {
		if err := rows.Scan(&id, &price, &createdDate); err != nil {
			f.Println("ERROR finding value", err)
		}
		f.Println("Exact match")
		f.Println("ID : ", id)
		f.Println("Rate : ", price)
		f.Println("Created at: ", createdDate)
	} else {
		rows, err := db.Query("SELECT id, rate, created_at FROM exchange_rate ORDER BY ABS(UNIX_TIMESTAMP(created_at) - UNIX_TIMESTAMP(?))  LIMIT 1", searchdata)
		if err != nil {
			f.Println("Searched value", err)
			return err
		}
		if rows.Next() {
			if err := rows.Scan(&id, &price, &createdDate); err != nil {
				f.Println("ERROR finding value", err)
			}
			f.Println("Nearest match")
			f.Println("ID : ", id)
			f.Println("Rate : ", price)
			f.Println("Created at: ", createdDate)
		}
	}
	return nil
}

func GetAverage(db *sql.DB, startDate string, endDate string) error {
	rows, err := db.Query("SELECT AVG(rate) FROM exchange_rate WHERE created_at BETWEEN ?  AND ?", startDate, endDate)
	if err != nil {
		f.Println("Searched value", err)
		return err
	}
	defer rows.Close()
	var rateValue float64
	if rows.Next() {
		if err := rows.Scan(&rateValue); err != nil {
			f.Println("ERROR finding value", err)
		}
		f.Println("Exact match")
		f.Println("Average Rate : ", rateValue)
	}
	return nil
}
