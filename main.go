package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func GetAllData(w http.ResponseWriter, r *http.Request) {}
func GetData(w http.ResponseWriter, r *http.Request) {}
func CreateData(w http.ResponseWriter, r *http.Request) {}


// main function
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/people", GetAllData).Methods("GET")
	router.HandleFunc("/people/{id}", GetData).Methods("GET")
	router.HandleFunc("/people/{id}", CreateData).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", router))
}

type Data struct {
	ID        string   `json:"id,omitempty"`
	Data_type string   `json:"data_type,omitempty"`
	Hash string `json:"hash,omitempty"`
}
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person