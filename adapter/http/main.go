package http

import (
	"net/http"

	"github.com/taisjjorge/sistema-planejamento-financeiro/adapter/http/actuator"
	"github.com/taisjjorge/sistema-planejamento-financeiro/adapter/http/transaction"
)

// Init ...
func Init(){
	http.HandleFunc("/transactions", transaction.GetTransactions)
	http.HandleFunc("/transaction", transaction.CreateATransaction)

	http.HandleFunc("/health", actuator.Health)

	http.ListenAndServe(":8080", nil)
}