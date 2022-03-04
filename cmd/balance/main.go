package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/bernmarx/avito-assignment/internal/http/api"
)

const (
	port = ":8080"
)

func main() {
	s := api.NewService()

	r := mux.NewRouter()

	r.HandleFunc("/deposit", s.GetDepositHandler()).Methods("POST")
	r.HandleFunc("/withdraw", s.GetWithdrawHandler()).Methods("POST")
	r.HandleFunc("/transfer", s.GetTransferHandler()).Methods("POST")
	r.HandleFunc("/getbalance", s.GetBalanceHandler()).Methods("GET")
	r.HandleFunc("/gethistory", s.GetTransactionHistoryHandler()).Methods("GET")
	r.HandleFunc("/gethistory/{page}", s.GetTransactionHistoryPageHandler()).Methods("GET")

	http.Handle("/", r)
	log.Println("Starting server at " + port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
