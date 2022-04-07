package main

import (
	"database/sql"
	"errors"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"github.com/bernmarx/avito-assignment/internal/app/http/api/deposit_post"
	"github.com/bernmarx/avito-assignment/internal/app/http/api/get_balance_get"
	"github.com/bernmarx/avito-assignment/internal/app/http/api/get_balance_history_get"
	"github.com/bernmarx/avito-assignment/internal/app/http/api/get_balance_history_page_get"
	"github.com/bernmarx/avito-assignment/internal/app/http/api/transfer_post"
	"github.com/bernmarx/avito-assignment/internal/app/http/api/withdraw_post"
	"github.com/bernmarx/avito-assignment/internal/domain/balance"
	"github.com/bernmarx/avito-assignment/internal/infrastructure/exchangerateclient"
	"github.com/bernmarx/avito-assignment/internal/infrastructure/log"
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

	//If connection to DB failed, wait for 3 seconds and try again
	if err != nil {
		time.Sleep(time.Second * 3)

		//If connection still fails, stop service
		db, err = connectToDB()
		if err != nil {
			sentry.CaptureException(err)
			log.Logger().WithError(err).Fatal(err.Error())
		}
	}

	err = sentry.Init(sentry.ClientOptions{
		Dsn:     os.Getenv("SENTRY_DSN"),
		Release: os.Getenv("SENTRY_RELEASE"),
	})

	sentry.CaptureException(errors.New("my error"))

	defer sentry.Flush(time.Second * 5)

	if err != nil {
		log.Logger().WithError(err).Fatal(err.Error())
	}

	s := balance.NewStorage(db)

	eRurl := os.Getenv("EXCHANGE_RATE_API_URL")

	eRcurPos, _ := strconv.ParseInt(os.Getenv("EXCHANGE_RATE_API_CUR_POS"), 10, 0)

	eR := exchangerateclient.NewExchangeRate(http.DefaultClient, eRurl, int(eRcurPos))

	r := mux.NewRouter()

	r.HandleFunc("/deposit", deposit_post.Handler(s)).Methods("POST")
	r.HandleFunc("/withdraw", withdraw_post.Handler(s)).Methods("POST")
	r.HandleFunc("/transfer", transfer_post.Handler(s)).Methods("POST")
	r.HandleFunc("/balance", get_balance_get.Handler(s, eR)).Methods("GET")
	r.HandleFunc("/history", get_balance_history_get.Handler(s)).Methods("GET")
	r.HandleFunc("/history/{page}", get_balance_history_page_get.Handler(s)).Methods("GET")

	http.Handle("/", r)

	log.Logger().Info("Starting server on " + os.Getenv("SERVICE_PORT"))

	err = http.ListenAndServe(":"+os.Getenv("SERVICE_PORT"), nil)
	if err != nil {
		log.Logger().WithError(err).Fatal(err.Error())
	}
}
