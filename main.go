package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"
)

func main() {
	server := http.Server{
		Addr:              ":8081",
		WriteTimeout:      time.Second * 2, // write request не больше 2 секунд
	}
	http.HandleFunc("/api/v1/", GetHandler)

	log.Println("Started listen and serve at :8081")
	log.Fatalln(server.ListenAndServe())
}


func GetHandler(w http.ResponseWriter, r *http.Request) {
	defer func(t time.Time) { log.Printf("Done request %s in %s", r.URL.Path, time.Since(t)) } (time.Now())

	w.WriteHeader(200)
	log.Println("Got a request:", r.URL.Path)

	hostname, err := os.Hostname()
	if err != nil {
		log.Println("error getting hostname:", err)
		ReturnErrorResponse(&w, http.StatusInternalServerError, err)
		return
	}

	weatherInfo, err := GetWeather()
	if err != nil {
		log.Println("error getting weather:", err)
		ReturnErrorResponse(&w, http.StatusInternalServerError, err)
		return
	}

	ReturnBaseResponse(&w, http.StatusOK, map[string]interface{}{
		"hostname" : hostname,
		"cpu_amount": runtime.NumCPU(),
		"go_version": runtime.Version(),
		"moscow_weather_today": weatherInfo,
	})
}

func GetWeather() (map[string]interface{}, error){
	response, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=Moscow&APPID=b3927aa371c03b7b3a4bf7f5e39853b3")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var weatherInfo map[string]interface{}
	if err = json.Unmarshal(body, &weatherInfo); err != nil {
		return nil, err
	}

	return weatherInfo, nil
}
