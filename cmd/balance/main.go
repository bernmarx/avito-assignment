package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"github.com/bernmarx/avito-assignment/internal/balance"
	"github.com/bernmarx/avito-assignment/internal/exchangerateapi"
	"github.com/bernmarx/avito-assignment/internal/http/api"
)

func connectToDB() (*sql.DB, error) {
	connData := "host=" + os.Getenv("DB_HOST") + " " + "port=" + os.Getenv("DB_PORT")
	connData = connData + " " + "user=" + os.Getenv("DB_USER") + " " + "password=" + os.Getenv("DB_PASSWORD")
	connData = connData + " " + "dbname=" + os.Getenv("DB_NAME") + " " + "sslmode=" + os.Getenv("DB_SSLMODE")

	db, err := sql.Open("postgres", connData)
	if err != nil {
		return db, err
	}

	err = db.Ping()

	return db, err
}

func main() {
	db, err := connectToDB()

	//If connection to Postgres failed, wait for 3 seconds and try again
	if err != nil {
		time.Sleep(time.Second * 3)

		//If connection still fails, stop service
		db, err = connectToDB()
		if err != nil {
			sentry.CaptureException(err)
			log.Fatalln(err)
		}
	}

	err = sentry.Init(sentry.ClientOptions{
		Dsn:         os.Getenv("SENTRY_DSN"),
		Environment: "",
		Release:     os.Getenv("SENTRY_RELEASE"),
		Debug:       true,
	})

	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

	s := balance.NewStorage(db)

	eRurl := os.Getenv("EXCHANGE_RATE_API_URL")

	eRcurPos, _ := strconv.ParseInt(os.Getenv("EXCHANGE_RATE_API_CUR_POS"), 10, 0)

	eR := exchangerateapi.NewExchangeRate(http.DefaultClient, eRurl, int(eRcurPos))

	service := api.NewService()

	r := mux.NewRouter()

	r.HandleFunc("/deposit", service.GetDepositHandler(s, eR)).Methods("POST")
	r.HandleFunc("/withdraw", service.GetWithdrawHandler(s, eR)).Methods("POST")
	r.HandleFunc("/transfer", service.GetTransferHandler(s, eR)).Methods("POST")
	r.HandleFunc("/balance", service.GetBalanceHandler(s, eR)).Methods("GET")
	r.HandleFunc("/history", service.GetTransactionHistoryHandler(s, eR)).Methods("GET")
	r.HandleFunc("/history/{page}", service.GetTransactionHistoryPageHandler(s, eR)).Methods("GET")

	http.Handle("/", r)

	err = http.ListenAndServe(":"+os.Getenv("SERVICE_PORT"), nil)
	if err != nil {
		log.Fatalln(err.Error())
	}

	log.Println("Server was started at " + os.Getenv("SERVICE_PORT") + " and is listening")
}
