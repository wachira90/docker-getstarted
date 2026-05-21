package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func api(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Api!!")
	fmt.Println("Endpoint Hit: Api")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/api", api)
	fmt.Println("REST STARTING .... ")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}