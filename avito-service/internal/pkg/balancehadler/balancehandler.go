package balancehadler

import (
	"encoding/json"
	"log"
	"net/http"

	"database/sql"
	_ "github.com/lib/pq"
)

const (
	host     = "host=db "
	port     = "port=5432 "
	user     = "user=handler "
	password = "password=pass "
	dbname   = "dbname=accounts "
	ssl_mode = "sslmode=disable"
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
		http.Error(w, "Database refused deposit\nDetails: "+err.Error(), http.StatusInternalServerError)
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
		http.Error(w, "Database refused withdrawal\nDetails: "+err.Error(), http.StatusInternalServerError)
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

func Get(w http.ResponseWriter, r *http.Request) {
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

	rows := DataBase.QueryRow(sqlstmt, account.ID)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Database refused get\nDetails: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var result userInfo
	result.ID = account.ID

	err = rows.Scan(&result.Balance)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Failed to parse result\nDetails: "+err.Error(), http.StatusInternalServerError)
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
