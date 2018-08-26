package handlers

import (
	"net/http"
	"../modules/storage"
	"../modules/billing"
	"io/ioutil"
	"fmt"
	"encoding/json"
)

type PayRequest struct {
	SendAccountId int
	ReceiveAccountId int
	Summ string
}

func PayToAccount(db storage.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("Error reading the body: %v\n", err)
			return
		}
		defer r.Body.Close()

		var requestData PayRequest
		if err = json.Unmarshal(body, &requestData); err != nil {
			panic(err)
		}

		if requestData.SendAccountId <=0 || err !=nil {
			http.Error(w, "sendAccountId is empty", http.StatusBadRequest)
			return
		}

		if requestData.ReceiveAccountId <=0 || err !=nil {
			http.Error(w, "receiveAccountId is empty", http.StatusBadRequest)
			return
		}

		if requestData.Summ == "" {
			http.Error(w, "summ is empty", http.StatusBadRequest)

			return
		}

		defer r.Body.Close()

		accountFrom,ok := db.Get(requestData.SendAccountId)
		if  ok != nil {
			http.Error(w, "db error: " + ok.Error(), http.StatusInternalServerError)
			return
		}

		accountTo, ok := db.Get(requestData.ReceiveAccountId)
		if  ok != nil {
			http.Error(w, "db error: " + ok.Error(), http.StatusInternalServerError)
			return
		}

		status := billing.PayToAccount(accountFrom, accountTo, requestData.Summ)

		if !status {
			http.Error(w, "billing error: payment failed", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	})
}