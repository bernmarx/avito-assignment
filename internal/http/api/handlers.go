package api

import (
	"encoding/json"
	"github.com/bernmarx/avito-assignment/internal/domain/balance"
	"github.com/bernmarx/avito-assignment/internal/domain/conversion"
	"github.com/bernmarx/avito-assignment/internal/exchangerateapi"
	"log"
	"net/http"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetDepositHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		var data Transaction
		json.NewDecoder(r.Body).Decode(&data)

		b := balance.NewBalance()

		err := b.Deposit(data.ID, data.Amount)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Deposit was successful"))
	}
}

func (s *Service) GetWithdrawHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		var data Transaction
		json.NewDecoder(r.Body).Decode(&data)

		b := balance.NewBalance()

		err := b.Withdraw(data.ID, data.Amount)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Withdrawal was successful"))
	}
}

func (s *Service) GetTransferHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		var data Transaction
		json.NewDecoder(r.Body).Decode(&data)

		b := balance.NewBalance()

		err := b.Transfer(data.ID, data.Receiver, data.Amount)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Transfer was successful"))
	}
}

func (s *Service) GetBalanceHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		var data UserBalance
		json.NewDecoder(r.Body).Decode(&data)

		b := balance.NewBalance()

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
			c := conversion.NewConverter(exchangerateapi.NewExchangeRate())
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
