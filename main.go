package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type CitiesResponse struct {
	Cities []string `json:"cities"`
}

func CityHandler(res http.ResponseWriter, req *http.Request) {
	citiesResponse := &CitiesResponse{
		Cities: []string{
			"San Francisco",
			"Amsterdam",
			"Berlin",
			"New York",
		},
	}
	data, _ := json.MarshalIndent(citiesResponse, "", "  ")
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.Write(data)
}

func main() {
	log.Println("Listening on this host: http://localhost:5000")

	http.HandleFunc("/cities.json", CityHandler)
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal("Unable to listen on :5000: ", err)
	}
}
