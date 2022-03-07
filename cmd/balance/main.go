package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"github.com/bernmarx/avito-assignment/internal/balance"
	"github.com/bernmarx/avito-assignment/internal/http/api"
)

const (
	port = ":8080"
)

func connectToDB() (*sql.DB, error) {
	connData := "host=" + os.Getenv("DB_HOST") + " " + "port=" + os.Getenv("DB_PORT")
	connData = connData + " " + "user=" + os.Getenv("DB_USER") + " " + "password=" + os.Getenv("DB_PASSWORD")
	connData = connData + " " + "dbname=" + os.Getenv("DB_NAME") + " " + "sslmode=" + os.Getenv("DB_SSLMODE")
	log.Println(connData)

	db, err := sql.Open("postgres", connData)
	if err != nil {
		return db, err
	}

	err = db.Ping()

	return db, err
}

func main() {
	db, err := connectToDB()
	if err != nil {
		log.Fatalln(err)
	}

	s := balance.NewStorage(db)

	service := api.NewService()

	r := mux.NewRouter()

	r.HandleFunc("/deposit", service.GetDepositHandler(s)).Methods("POST")
	r.HandleFunc("/withdraw", service.GetWithdrawHandler(s)).Methods("POST")
	r.HandleFunc("/transfer", service.GetTransferHandler(s)).Methods("POST")
	r.HandleFunc("/balance", service.GetBalanceHandler(s)).Methods("GET")
	r.HandleFunc("/history", service.GetTransactionHistoryHandler(s)).Methods("GET")
	r.HandleFunc("/history/{page}", service.GetTransactionHistoryPageHandler(s)).Methods("GET")

	http.Handle("/", r)
	log.Println("Starting server at " + port)

	err = http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
