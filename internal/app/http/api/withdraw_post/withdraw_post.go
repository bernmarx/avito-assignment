package withdraw_post

import (
	"encoding/json"
	"net/http"

	"github.com/bernmarx/avito-assignment/internal/app/http/api"
	"github.com/bernmarx/avito-assignment/internal/domain/balance"
	"github.com/bernmarx/avito-assignment/internal/infrastructure/errors"
	"github.com/bernmarx/avito-assignment/internal/infrastructure/log"
)

// Handler for withdraw POST method
func Handler(strg balance.StorageAccess, eR balance.ExchangeRateGetter) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		var rd api.DepositWithdrawRequestData

		json.NewDecoder(r.Body).Decode(&rd)

		b := balance.NewBalance(strg, eR)

		err := b.Storage.WithdrawMoney(rd.AccountID, rd.BalanceID, rd.Amount)
		if err != nil {
			err := err.(*errors.Error)

			log.Logger().WithError(err).Error(err.Error())
			http.Error(w, err.Msg, err.Code)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Withdrawal was successful"))
	}
}
