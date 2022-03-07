package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/bernmarx/avito-assignment/internal/balance"
	"github.com/bernmarx/avito-assignment/internal/converter"
	"github.com/bernmarx/avito-assignment/internal/exchangerateapi"
	"github.com/gorilla/mux"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetDepositHandler(strg balance.StorageAccess) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		var data Transaction
		json.NewDecoder(r.Body).Decode(&data)

		b := balance.NewBalance(strg)

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

func (s *Service) GetWithdrawHandler(strg balance.StorageAccess) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		var data Transaction
		json.NewDecoder(r.Body).Decode(&data)

		b := balance.NewBalance(strg)

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

func (s *Service) GetTransferHandler(strg balance.StorageAccess) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		var data Transaction
		json.NewDecoder(r.Body).Decode(&data)

		b := balance.NewBalance(strg)

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

func (s *Service) GetBalanceHandler(strg balance.StorageAccess) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		var data RequestData
		json.NewDecoder(r.Body).Decode(&data)

		b := balance.NewBalance(strg)

		var err error
		data.Balance, err = b.GetBalance(data.ID)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		query := r.URL.Query()
		currency, exists := query["currency"]

		if exists && len(currency) == 1 {
			c := converter.NewConverter(exchangerateapi.NewExchangeRate())
			data.Balance, err = c.ConvertCurrency(data.Balance, currency[0])
			if err != nil {
				log.Println(err.Error())
				http.Error(w, "Error while converting currency\nDetails: "+err.Error(), http.StatusInternalServerError)
				return
			}
		}
		if len(currency) > 1 {
			log.Println("Invalid conversion query")
			http.Error(w, "Invalid conversion query", http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		j, err := data.GetJSON()
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Write(j)
	}
}

func (s *Service) GetTransactionHistoryHandler(strg balance.StorageAccess) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		var data RequestData
		json.NewDecoder(r.Body).Decode(&data)

		b := balance.NewBalance(strg)

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

func (s *Service) GetTransactionHistoryPageHandler(strg balance.StorageAccess) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		var data RequestData
		json.NewDecoder(r.Body).Decode(&data)

		variables := mux.Vars(r)
		page64, err := strconv.ParseInt(variables["page"], 10, 0)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		page := int(page64)

		b := balance.NewBalance(strg)

		j, err := b.GetTransactionHistoryPage(data.ID, data.Sort, page)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(j)
	}
}
