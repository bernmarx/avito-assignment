package get_history_page_get

import (
	"encoding/json"
	standardErrors "errors"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/bernmarx/avito-assignment/internal/app/http/api"
	"github.com/bernmarx/avito-assignment/internal/domain/balance"
	"github.com/bernmarx/avito-assignment/internal/infrastructure/errors"
)

func Handler(strg balance.StorageAccess, eR balance.ExchangeRateGetter) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		var rd api.RequestData
		json.NewDecoder(r.Body).Decode(&rd)

		if rd.Sort == "" {
			rd.Sort = os.Getenv("DEFAULT_SORT")
		}

		variables := mux.Vars(r)

		page64, err := strconv.ParseInt(variables["page"], 10, 0)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		page := int(page64)

		b := balance.NewBalance(strg, eR)

		j, err := b.GetTransactionHistoryPage(rd.ID, rd.Sort, page)
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
