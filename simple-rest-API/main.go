package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Sample struct {
	id   int
	name string
	desc string
}

type Samples []Sample

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page Endpoint Hit")
	fmt.Println("Home Page Endpoint is being Accessed")
}

func handleReqs() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/sample", returnSamples)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleReqs()
}

func returnSamples(w http.ResponseWriter, _ *http.Request) {
	samples := Samples{
		Sample{id: 1, name: "suhas", desc: "Go King"},
		Sample{id: 2, name: "Go", desc: "Web Application Development Language"},
	}

	fmt.Println("Samples Endpoint Hit: Serving Samples...!")
	json.NewEncoder(w).Encode(samples)
}
