package get_history_get

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

		var data api.RequestData
		json.NewDecoder(r.Body).Decode(&data)

		b := balance.NewBalance(strg, eR)

		j, err := b.GetTransactionHistory(data.ID)
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

		w.WriteHeader(http.StatusOK)
		w.Write(j)
	}
}
