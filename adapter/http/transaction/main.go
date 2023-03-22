package transaction

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/taisjjorge/sistema-planejamento-financeiro/model/transaction"
	"github.com/taisjjorge/sistema-planejamento-financeiro/util"
)

// GetTransactions ...
func GetTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-type", "application/json")

	var transactions = transaction.Transactions{
		transaction.Transaction{
			Title:     "Salário",
			Amount:    12000.0,
			Type:      0,
			CreatedAt: util.StringToTime("2022-03-22T20:24:00"),
		},
	}

	_ = json.NewEncoder(w).Encode(transactions)
}

// CreateATransaction ...
func CreateATransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var res = transaction.Transactions{}
	var body, _ = ioutil.ReadAll(r.Body)

	_ = json.Unmarshal(body, &res)

	fmt.Print(res)
}
