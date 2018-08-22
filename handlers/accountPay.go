package handlers

import (
	"net/http"
	"strconv"
	"../modules/storage"
	"../modules/billing"
)

func PayToAccount(db storage.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" {
			http.Error(w, "wrong http method. PUT required", http.StatusBadRequest)
			return
		}

		sendAccountId,err := strconv.Atoi(r.URL.Query().Get("sendAccountId"))
		if sendAccountId <=0 || err !=nil {
			http.Error(w, "sendAccountId is empty", http.StatusBadRequest)
			return
		}

		receiveAccountId,err := strconv.Atoi(r.URL.Query().Get("receiveAccountId"))
		if receiveAccountId <=0 || err !=nil {
			http.Error(w, "receiveAccountId is empty", http.StatusBadRequest)
			return
		}

		summ := r.URL.Query().Get("summ")
		if summ == "" {
			http.Error(w, "summ is empty", http.StatusBadRequest)

			return
		}

		defer r.Body.Close()

		accountFrom,ok := db.Get(sendAccountId)
		if  ok != nil {
			http.Error(w, "db error: " + ok.Error(), http.StatusInternalServerError)
			return
		}

		accountTo, ok := db.Get(receiveAccountId)
		if  ok != nil {
			http.Error(w, "db error: " + ok.Error(), http.StatusInternalServerError)
			return
		}

		status := billing.PayToAccount(accountFrom, accountTo, summ)

		if !status {
			http.Error(w, "billing error: payment failed", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	})
}