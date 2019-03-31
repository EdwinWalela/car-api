package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var PORT = "8000"

type Car struct {
	id       uint64
	category string
	make     string
	model    string
	HP       uint64
	year     uint64
}

var cars = []Car{
	{1, "Sedan", "Toyota", "Vitz", 150, 2006},
	{2, "Sedan", "Mercedes-Benz", "s550", 740, 2018},
	{3, "SUV", "Lamborghini", "Urus", 650, 2019},
	{4, "SUV", "Tesla", "Model X", 450, 2019},
	{5, "SUV", "Range Rover", "Evoque", 237, 2019},
	{6, "Sedan", "BMW", "750i", 350, 2015},
	{7, "Coupe", "Buggati", "Chiron", 1500, 2018},
	{8, "Sedan", "Bentley", "Continental GT", 552, 2018}}

func IndexHandler(w http.ResponseWriter, r *http.Request) {

}

func IdHandler(w http.ResponseWriter, r *http.Request) {

}

func CategoryHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler)
	r.HandleFunc("cars/{id}", IdHandler)
	r.HandleFunc("cars/category/{category}", CategoryHandler)

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
		fmt.Println("Listening to requests on port:", PORT)
	}()

}
