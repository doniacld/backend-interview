package main

import (
	"net/http"

	"github.com/gustvision/backend-interview/pkg/account"
	"github.com/gustvision/backend-interview/pkg/user"
)

type handler struct {
	user    user.App
	account account.App
}

func (h handler) listen(host string) {
	http.HandleFunc("/user", h.GetUser)
	http.HandleFunc("/account", h.GetAccount)
	http.HandleFunc("/transaction", h.GetTransaction)

	http.ListenAndServe(host, nil)
}
