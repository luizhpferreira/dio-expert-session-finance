package http

import (
	"net/http"

	"github.com/luizhpferreira/dio-expert-session-finance/model/transaction/adapter/http/actuator"
	"github.com/luizhpferreira/dio-expert-session-finance/model/transaction/adapter/http/transaction"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Init() {
	http.HandleFunc("/transactions", transaction.GetTransactions)
	http.HandleFunc("/transactions/create", transaction.CreateTransaction)

	http.Handle("/metrics", promhttp.Handler())

	http.HandleFunc("/health", actuator.Health)

	http.ListenAndServe(":8080", nil)
}
