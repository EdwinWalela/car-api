package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

var PORT = "8000"

type Car struct {
	Id       uint64 `json:"id"`
	Category string `json:"category"`
	Make     string `json:"make"`
	Model    string `json:"model"`
	HP       uint16 `json:"HP`
	Year     uint16 `json:"year"`
}

type MultiResponse struct {
	Status string `json:"status"`
	Cars   []Car  `json:"cars",omitempty`
	Count  int    `json:"count"`
}

type SingleResponse struct {
	Status string `json:"status"`
	Car    Car    `json:"car,omitempty`
}

var AllCars = []Car{
	{1, "Sedan", "Toyota", "Vitz", 150, 2006},
	{2, "Sedan", "Mercedes-Benz", "s550", 740, 2018},
	{3, "SUV", "Lamborghini", "Urus", 650, 2019},
	{4, "SUV", "Tesla", "Model X", 450, 2019},
	{5, "SUV", "Range Rover", "Evoque", 237, 2019},
	{6, "Sedan", "BMW", "750i", 350, 2015},
	{7, "Coupe", "Buggati", "Chiron", 1500, 2018},
	{8, "Sedan", "Bentley", "Continental GT", 552, 2018}}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Create Response Object
	result := &MultiResponse{
		Status: "200",
		Count:  len(AllCars),
		Cars:   AllCars,
	}
	// Encode response to JSON
	res, _ := json.Marshal(result)
	// Send response
	w.Write([]byte(res))
}

func IndexOf(s []Car, v uint64) bool {
	for _, e := range s {
		if e.Id == v {
			return true
		}
	}
	return false
}

func idHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	result := &SingleResponse{}

	if IndexOf(AllCars, uint64(id)) {
		result.Status = "200"
		result.Car = AllCars[id-1]
	} else {
		result.Status = "404"
	}
	res, _ := json.Marshal(result)

	w.Write([]byte(res))
}

func categoryHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/cars", indexHandler)
	r.HandleFunc("/cars/{id}", idHandler)
	r.HandleFunc("/cars/category/{category}", categoryHandler)

	srv := http.Server{
		Addr: "0.0.0.0:" + PORT,
		//Set timeouts to prevent Slowloris attacks
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r}

	func() {
		if err := srv.ListenAndServe(); err != nil {
			panic(err)
		}
	}()

}
