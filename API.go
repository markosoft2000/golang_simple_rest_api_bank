//go get github.com/shopspring/decimal
//go get github.com/gorilla/mux
package main

import (
	"./modules/storage"
	"./handlers"
	"net/http"
	"log"
	"github.com/gorilla/mux"
)

func main()  {
	accountStorage := storage.NewInMemoryDB()

	muxRouter := mux.NewRouter().StrictSlash(true)
	muxRouter.Handle("/account/create", handlers.CreateAccount(accountStorage)).Methods("POST")
	muxRouter.Handle("/account/get/{id}", handlers.GetAccount(accountStorage)).Methods("GET")
	muxRouter.Handle("/account/get/{id}/amount", handlers.GetAmount(accountStorage)).Methods("GET")
	muxRouter.Handle("/account/pay", handlers.PayToAccount(accountStorage)).Methods("PUT")

	err := http.ListenAndServe(":7000", muxRouter)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}