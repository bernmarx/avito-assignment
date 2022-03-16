package get_balance_get

import (
	"encoding/json"
	standardErrors "errors"
	"log"
	"net/http"

	"github.com/bernmarx/avito-assignment/internal/app/http/api"
	"github.com/bernmarx/avito-assignment/internal/domain/balance"
	"github.com/bernmarx/avito-assignment/internal/infrastructure/errors"
)

func Handler(strg balance.StorageAccess, eR balance.ExchangeRateGetter) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		var rd api.RequestData
		var acc api.Account
		json.NewDecoder(r.Body).Decode(&rd)

		b := balance.NewBalance(strg, eR)

		var err error
		acc.Balance, err = b.GetBalance(rd.ID)
		if err != nil {
			var sErr *errors.Error

			if standardErrors.As(err, &sErr) {
				log.Println(sErr.Msg)
				http.Error(w, sErr.Msg, sErr.Code)
				return
			}

			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		query := r.URL.Query()
		currency, exists := query["currency"]

		//Returns error if there is more than 1 currency value
		if exists && (len(currency) != 1) {
			log.Println("Invalid conversion query")
			http.Error(w, "Invalid conversion query", http.StatusBadRequest)
			return
		}

		if exists {
			rate, err := eR.GetExchangeRate(currency[0])
			if err != nil {
				log.Println(err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			acc.Balance *= rate
		}

		acc.ID = rd.ID

		w.WriteHeader(http.StatusOK)
		j, err := acc.GetJSON()
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(j)
	}
}
