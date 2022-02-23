package main

import (
	"log"
	"net/http"

	"github.com/bernmarx/avito-assignment/avito-service/internal/pkg/balancehadler"
)

const (
	port = ":8080"
)

func main() {
	http.HandleFunc("/deposit", balancehadler.Deposit)
	http.HandleFunc("/withdraw", balancehadler.Withdraw)
	http.HandleFunc("/transfer", balancehadler.Transfer)
	http.HandleFunc("/get", balancehadler.Get)

	log.Println("Starting server at " + port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
