//go get github.com/shopspring/decimal
//go get github.com/gorilla/mux
package main

import (
	"./modules/storage"
	"./modules/billing"
	"./handlers"
	"net/http"
	"log"
	"github.com/gorilla/mux"
)

func main()  {
	accountStorage := storage.NewInMemoryDB()
	payChan := make(chan *billing.GoPay)
	payExitChan := make(chan bool)
	defer close(payExitChan)
	defer close(payChan)

	go billing.GoPayToAccount(payChan, payExitChan)

	muxRouter := mux.NewRouter().StrictSlash(true)
	muxRouter.Handle("/account/create", handlers.CreateAccount(accountStorage)).Methods("POST")
	muxRouter.Handle("/account/get/{id}", handlers.GetAccount(accountStorage)).Methods("GET")
	muxRouter.Handle("/account/get/{id}/amount", handlers.GetAmount(accountStorage)).Methods("GET")
	muxRouter.Handle("/account/pay", handlers.PayToAccount(accountStorage)).Methods("PUT")
	muxRouter.Handle("/account/pay2", handlers.PayToAccountChan(accountStorage, payChan)).Methods("PUT")

	err := http.ListenAndServe(":7002", muxRouter)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}