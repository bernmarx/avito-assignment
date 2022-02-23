package main

import (
	"log"
	"net/http"
)

const (
	port = ":8080"
)

func main() {
	http.HandleFunc("/deposit", deposit)

	log.Println("Starting server at " + port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func deposit(w http.ResponseWriter, r *http.Request) {

}
