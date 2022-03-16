package transfer_post

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

		var data api.Transaction
		json.NewDecoder(r.Body).Decode(&data)

		b := balance.NewBalance(strg, eR)

		err := b.MakeTransfer(data.ID, data.Receiver, data.Amount)
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
		w.Write([]byte("Transfer was successful"))
	}
}
