package balancehadler

import (
	"database/sql"
	"encoding/json"
	"errors"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

const (
	host     = "host=localhost "
	port     = "port=5432 "
	user     = "user=handler "
	password = "password=pass "
	dbname   = "dbname=test2 "
	ssl_mode = "sslmode=disable"

	exchangeAPIkey = "58ba13c0e8dd60f3f3a39804"
	baseCurrency   = "RUB"
)

type dwRequest struct {
	ID     int     `json:"id"`
	Amount float32 `json:"amount"`
}
type transferRequest struct {
	Sender   int     `json:"sender"`
	Reciever int     `json:"reciever"`
	Amount   float32 `json:"amount"`
}
type userInfo struct {
	ID      int     `json:"id"`
	Balance float32 `json:"balance"`
}
type depositInfo struct {
	Time   string  `json:"time"`
	Amount float32 `json:"amount"`
}
type withdrawalInfo struct {
	Time   string  `json:"time"`
	Amount float32 `json:"amount"`
}
type transferInfo struct {
	Time     string  `json:"time"`
	Reciever int     `json:"reciever"`
	Amount   float32 `json:"amount"`
}
type transactions struct {
	ID         int              `json:"id"`
	Deposit    []depositInfo    `json:"deposit"`
	Withdrawal []withdrawalInfo `json:"withdrawal"`
	Transfer   []transferInfo   `json:"transfer"`
}

func Deposit(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Wrong method", http.StatusMethodNotAllowed)
		return
	}

	var deposit dwRequest

	err := json.NewDecoder(r.Body).Decode(&deposit)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Bad JSON", http.StatusBadRequest)
		return
	}

	if deposit.Amount <= 0 {
		http.Error(w, "Can't deposit non-positive sum", http.StatusBadRequest)
		return
	}
	if deposit.ID <= 0 {
		http.Error(w, "Invalid ID or JSON", http.StatusBadRequest)
		return
	}

	DataBase, err := sql.Open("postgres", host+port+user+password+dbname+ssl_mode)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		return
	}

	sqlstmt := `call balance_deposit($1, $2)`

	_, err = DataBase.Exec(sqlstmt, deposit.ID, deposit.Amount)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Database refused deposit\nDetails: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successfully made a deposit!"))
}

func Withdraw(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Wrong method", http.StatusMethodNotAllowed)
		return
	}

	var withdraw dwRequest

	err := json.NewDecoder(r.Body).Decode(&withdraw)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Bad JSON", http.StatusBadRequest)
		return
	}

	if withdraw.Amount <= 0 {
		http.Error(w, "Can't withdraw non-positive sum", http.StatusBadRequest)
		return
	}
	if withdraw.ID <= 0 {
		http.Error(w, "Invalid ID or JSON", http.StatusBadRequest)
		return
	}

	DataBase, err := sql.Open("postgres", host+port+user+password+dbname+ssl_mode)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		return
	}

	sqlstmt := `call balance_withdraw($1, $2)`

	_, err = DataBase.Exec(sqlstmt, withdraw.ID, withdraw.Amount)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Database refused withdrawal\nDetails: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successfully made a withdraw!"))
}

func Transfer(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Wrong method", http.StatusMethodNotAllowed)
		return
	}

	var transfer transferRequest

	err := json.NewDecoder(r.Body).Decode(&transfer)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Bad JSON", http.StatusBadRequest)
		return
	}

	if transfer.Amount <= 0 {
		http.Error(w, "Can't transfer non-positive sum", http.StatusBadRequest)
		return
	}
	if transfer.Sender <= 0 || transfer.Reciever <= 0 {
		http.Error(w, "Invalid ID or JSON", http.StatusBadRequest)
		return
	}

	DataBase, err := sql.Open("postgres", host+port+user+password+dbname+ssl_mode)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		return
	}

	sqlstmt := `call balance_transfer($1, $2, $3)`

	_, err = DataBase.Exec(sqlstmt, transfer.Sender, transfer.Reciever, transfer.Amount)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Database refused transfer\nDetails: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successfully made a transfer!"))
}

func GetBalance(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Wrong method", http.StatusMethodNotAllowed)
		return
	}

	var account userInfo

	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Bad JSON", http.StatusBadRequest)
		return
	}

	if account.ID <= 0 {
		http.Error(w, "Invalid ID or JSON", http.StatusBadRequest)
		return
	}

	DataBase, err := sql.Open("postgres", host+port+user+password+dbname+ssl_mode)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		return
	}

	sqlstmt := `SELECT balance_get($1)`

	row := DataBase.QueryRow(sqlstmt, account.ID)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Database refused get\nDetails: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var result userInfo
	result.ID = account.ID

	err = row.Scan(&result.Balance)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Failed to parse result\nDetails: "+err.Error(), http.StatusInternalServerError)
		return
	}

	query := r.URL.Query()
	currency, exists := query["currency"]

	if exists && len(currency) == 1 {
		result.Balance, err = convertToCurrency(result.Balance, currency[0])
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Error while converting currency\nDetails: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}
	if len(currency) > 1 {
		log.Println("Invalid currency query")
		http.Error(w, "Invalid currency query", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&result)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Failed to convert info into JSON\nDetails: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Wrong method", http.StatusMethodNotAllowed)
		return
	}

	var account userInfo

	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Bad JSON", http.StatusBadRequest)
		return
	}

	if account.ID <= 0 {
		http.Error(w, "Invalid ID or JSON", http.StatusBadRequest)
		return
	}

	DataBase, err := sql.Open("postgres", host+port+user+password+dbname+ssl_mode)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		return
	}

	getDeposit := `SELECT deposit_time, amount::numeric::FLOAT8 FROM deposit_journal WHERE account_id = $1`
	getWithdrawal := `SELECT withdraw_time, amount::numeric::FLOAT8 FROM withdraw_journal WHERE account_id = $1`
	getTransfer := `SELECT transfer_time, reciever_id, amount::numeric::FLOAT8 FROM transfer_journal WHERE sender_id = $1`

	var deposit []depositInfo
	var withdrawal []withdrawalInfo
	var transfer []transferInfo

	deposits, err := DataBase.Query(getDeposit, account.ID)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Database refused get deposit\nDetails: "+err.Error(), http.StatusInternalServerError)
		return
	}

	for deposits.Next() {
		var time string
		var amount float32
		err := deposits.Scan(&time, &amount)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Failed to parse data\nDetails: "+err.Error(), http.StatusInternalServerError)
			return
		}

		deposit = append(deposit, depositInfo{time, amount})
	}

	withdrawals, err := DataBase.Query(getWithdrawal, account.ID)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Database refused get withdrawal\nDetails: "+err.Error(), http.StatusInternalServerError)
		return
	}

	for withdrawals.Next() {
		var time string
		var amount float32
		err := withdrawals.Scan(&time, &amount)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Failed to parse data\nDetails: "+err.Error(), http.StatusInternalServerError)
			return
		}

		withdrawal = append(withdrawal, withdrawalInfo{time, amount})
	}

	transfers, err := DataBase.Query(getTransfer, account.ID)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Failed to parse data from database\nDetails: "+err.Error(), http.StatusInternalServerError)
		return
	}

	for transfers.Next() {
		var time string
		var reciever int
		var amount float32
		err := transfers.Scan(&time, &reciever, &amount)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Failed to parse data\nDetails: "+err.Error(), http.StatusInternalServerError)
			return
		}

		transfer = append(transfer, transferInfo{time, reciever, amount})
	}

	if deposit == nil && withdrawal == nil && transfer == nil {
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}

	transactionsInfo := transactions{
		account.ID,
		deposit,
		withdrawal,
		transfer,
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&transactionsInfo)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Failed to convert info into JSON\nDetails: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func convertToCurrency(rubles float32, currency string) (float32, error) {
	resp, err := http.Get("https://v6.exchangerate-api.com/v6/" + exchangeAPIkey + "/pair/" +
		baseCurrency + "/" + currency)
	if err != nil {
		return 0.0, err
	}
	defer resp.Body.Close()

	type conversion struct {
		Rate float32 `json:"conversion_rate"`
	}
	var cr conversion

	err = json.NewDecoder(resp.Body).Decode(&cr)
	if err != nil {
		return 0.0, err
	}
	if cr.Rate == 0.0 {
		return rubles, errors.New("Wrong currency name")
	}

	return rubles * cr.Rate, nil
}
