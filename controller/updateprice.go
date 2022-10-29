package controller

import (
	f "fmt"
	"io"
	"net/http"
)

func Updateprice() string {
	url := "https://rest.coinapi.io/v1/exchangerate/BTC/USD"
	apiKey := "E47244C5-C111-4350-AB32-EA30F52551C8"

	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("X-CoinAPI-Key", apiKey)

	resp, err := client.Do(req)
	if err != nil {
		f.Println(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	f.Println("Response ", string(body))
	jsonString := string(body)
	return jsonString
}
