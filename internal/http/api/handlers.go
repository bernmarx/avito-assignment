package api

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/bernmarx/avito-assignment/internal/balance"
	"github.com/gorilla/mux"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetDepositHandler(strg balance.StorageAccess, eR balance.ExchangeRateGetter) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		var data Transaction
		json.NewDecoder(r.Body).Decode(&data)

		b := balance.NewBalance(strg, eR)

		err := b.MakeDeposit(data.ID, data.Amount)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Deposit was successful"))
	}
}

func (s *Service) GetWithdrawHandler(strg balance.StorageAccess, eR balance.ExchangeRateGetter) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		var data Transaction
		json.NewDecoder(r.Body).Decode(&data)

		b := balance.NewBalance(strg, eR)

		err := b.MakeWithdraw(data.ID, data.Amount)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Withdrawal was successful"))
	}
}

func (s *Service) GetTransferHandler(strg balance.StorageAccess, eR balance.ExchangeRateGetter) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		var data Transaction
		json.NewDecoder(r.Body).Decode(&data)

		b := balance.NewBalance(strg, eR)

		err := b.MakeTransfer(data.ID, data.Receiver, data.Amount)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Transfer was successful"))
	}
}

func (s *Service) GetBalanceHandler(strg balance.StorageAccess, eR balance.ExchangeRateGetter) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		var rd RequestData
		var acc Account
		json.NewDecoder(r.Body).Decode(&rd)

		b := balance.NewBalance(strg, eR)

		var err error
		acc.Balance, err = b.GetBalance(rd.ID)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
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
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Write(j)
	}
}

func (s *Service) GetTransactionHistoryHandler(strg balance.StorageAccess, eR balance.ExchangeRateGetter) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		var data RequestData
		json.NewDecoder(r.Body).Decode(&data)

		b := balance.NewBalance(strg, eR)

		j, err := b.GetTransactionHistory(data.ID)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(j)
	}
}

func (s *Service) GetTransactionHistoryPageHandler(strg balance.StorageAccess, eR balance.ExchangeRateGetter) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		var rd RequestData
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
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(j)
	}
}
