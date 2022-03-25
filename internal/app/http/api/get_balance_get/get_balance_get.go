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

		var rd api.GetBalanceRequestData

		json.NewDecoder(r.Body).Decode(&rd)

		b := balance.NewBalance(strg, eR)

		balance, err := b.GetBalance(rd.Account_id, rd.Balance_id)
		if err != nil {
			err := err.(*errors.Error)

			log.Logger().WithError(err).Error(err.Error())
			http.Error(w, err.Msg, err.Code)
			return
		}

		query := r.URL.Query()
		currency, exists := query["currency"]

		if exists {
			rate, err := eR.GetExchangeRate(currency[0])
			if err != nil {
				log.Logger().WithError(err).Error(err.Error())
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			balance *= rate
		}

		response := api.Balance{
			Account_id: rd.Account_id,
			Balance_id: rd.Balance_id,
			Balance:    balance,
		}

		w.WriteHeader(http.StatusOK)

		j, err := json.Marshal(response)

		if err != nil {
			log.Logger().WithError(err).Error(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(j)
	}
}
