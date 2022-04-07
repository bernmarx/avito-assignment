package get_balance_history_get

import (
	"encoding/json"
	"net/http"

	"github.com/bernmarx/avito-assignment/internal/app/http/api"
	"github.com/bernmarx/avito-assignment/internal/domain/balance"
	"github.com/bernmarx/avito-assignment/internal/infrastructure/errors"
	"github.com/bernmarx/avito-assignment/internal/infrastructure/log"
)

// Handler for get_balance_history GET method
func Handler(strg balance.StorageAccess) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		var rd api.GetBalanceHistoryRequestData

		json.NewDecoder(r.Body).Decode(&rd)

		b := balance.NewBalance(strg)

		j, err := b.Storage.GetBalanceHistory(rd.AccountID, rd.BalanceID, rd.Sort, 0)

		if err != nil {
			err := err.(*errors.Error)

			log.Logger().WithError(err).Error(err.Error())
			http.Error(w, err.Msg, err.Code)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(j)
	}
}
