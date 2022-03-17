package get_balance_get

import (
	"encoding/json"
	"net/http"

	"github.com/bernmarx/avito-assignment/internal/app/http/api"
	"github.com/bernmarx/avito-assignment/internal/domain/balance"
	"github.com/bernmarx/avito-assignment/internal/infrastructure/errors"
	"github.com/bernmarx/avito-assignment/internal/infrastructure/log"
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
			err := err.(*errors.Error)

			log.Logger().WithError(err).Error(err.Error())
			http.Error(w, err.Msg, err.Code)
			return
		}

		query := r.URL.Query()
		currency, exists := query["currency"]

		//Returns error if there is more than 1 currency value
		if exists && (len(currency) != 1) {
			log.Logger().Info("invalid conversion query")
			http.Error(w, "Invalid conversion query", http.StatusBadRequest)
			return
		}

		if exists {
			rate, err := eR.GetExchangeRate(currency[0])
			if err != nil {
				log.Logger().WithError(err).Error(err.Error())
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			acc.Balance *= rate
		}

		acc.ID = rd.ID

		w.WriteHeader(http.StatusOK)
		j, err := acc.GetJSON()
		if err != nil {
			log.Logger().WithError(err).Error(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(j)
	}
}
