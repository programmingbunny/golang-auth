package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type person struct {
	Firstname string
}

func main() {

	http.HandleFunc("/encode", sliceOfValues)
	http.HandleFunc("/decode", decodingValues)

	http.ListenAndServe(":3000", nil)

}

// encode a slice of values
func sliceOfValues(w http.ResponseWriter, r *http.Request) {
	a := person{
		Firstname: "Bunny",
	}
	b := person{
		Firstname: "Jenny",
	}
	c := person{
		Firstname: "Programs",
	}

	people := []person{a, b, c}

	err := json.NewEncoder(w).Encode(people)
	if err != nil {
		log.Panic("issue with encoding", err)
	}

}

func decodingValues(w http.ResponseWriter, r *http.Request) {
	decodingPerson := []person{}

	err := json.NewDecoder(r.Body).Decode(&decodingPerson)
	if err != nil {
		log.Panic("issue with decoding", err)
	}
	fmt.Println(decodingPerson)
}
