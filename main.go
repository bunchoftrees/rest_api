package main

import (
	"crypto/sha1"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)
type Data_block struct {
	ID    string   `json:"id,omitempty"`
	Block string   `json:"data_block,omitempty"`
}

var data_blocks []Data_block

func GetData(w http.ResponseWriter, r *http.Request) {params := mux.Vars(r)
	for _, item := range data_blocks {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Data_block{})}

func CreateData(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var block Data_block
	_ = json.NewDecoder(r.Body).Decode(&block)
	block.ID = params["id"]
	data_blocks = append(data_blocks, block)
	json.NewEncoder(w).Encode(data_blocks)
	return
	//this doesn't work as expected
	HashThis(data_blocks)


}
//hash function untested
func HashThis(arr []Data_block) [20]byte {
	arrBytes := []byte{}
	for _, item := range arr {
		jsonBytes, _ := json.Marshal(item)
		arrBytes = append(arrBytes, jsonBytes...)
	}
	return sha1.Sum(arrBytes)
}


// main function
func main() {
	router := mux.NewRouter()
	data_blocks = append(data_blocks, Data_block{ID: "1", Block: "Forrest"})
	router.HandleFunc("/data/{id}", GetData).Methods("GET")
	router.HandleFunc("/data/{id}", CreateData).Methods("POST")
	log.Fatal(http.ListenAndServe(":8082", router))
}
