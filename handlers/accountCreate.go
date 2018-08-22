package handlers

import (
	"net/http"
	"../modules/storage"
	"../models"
	"../helpers"
	"strconv"
)

func CreateAccount(db storage.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		param,err := helpers.GetRequestParam(r, "id")
		id,err := strconv.Atoi(param)
		//id,err := strconv.Atoi(r.URL.Query().Get("id"));
		if id <=0 || err !=nil {
			http.Error(w, "id is empty" + err.Error(), http.StatusBadRequest)
			return
		}

		name,err := helpers.GetRequestParam(r, "name")
		//name := r.URL.Query().Get("name");
		if name == "" {
			http.Error(w, "name is empty", http.StatusBadRequest)

			return
		}

		amount,err := helpers.GetRequestParam(r, "amount")
		//amount := r.URL.Query().Get("amount")
		if amount == "" {
			http.Error(w, "amount is empty", http.StatusBadRequest)

			return
		}

		defer r.Body.Close()

		if _, ok := db.Get(id); ok == nil {
			http.Error(w, "db error: can not create account (account exists)", http.StatusInternalServerError)
			return
		}

		account := models.Account{}
		account.Init(id, name)
		account.SetAmount(amount)

		if ok := db.Set(account.GetId(), &account); !ok {
			http.Error(w, "db error: can not create account", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	})
}