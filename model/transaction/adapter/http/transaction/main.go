package transaction

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/luizhpferreira/dio-expert-session-finance/model/transaction"
	"github.com/luizhpferreira/dio-expert-session-finance/util"
)

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var transactions = transaction.Transactions{
		transaction.Transaction{
			Title:     "salario",
			Amount:    1200.0,
			Type:      0,
			CreatedAt: util.StringToTime("2020-01-02T10:04:05"),
		},
	}

	_ = json.NewEncoder(w).Encode(transactions)
}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var res = transaction.Transactions{}
	var body, _ = io.ReadAll(r.Body)

	json.Unmarshal(body, &res)
	fmt.Print(res)
}
