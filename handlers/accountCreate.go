package handlers

import (
	"net/http"
	"../modules/storage"
	"../models"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type CreateRequest struct {
	Id int
	Name string
	Amount string
}


func CreateAccount(db storage.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("Error reading the body: %v\n", err)
			return
		}
		defer r.Body.Close()

		var requestData CreateRequest
		if err = json.Unmarshal(body, &requestData); err != nil {
			panic(err)
		}

		if requestData.Id <= 0 {
			http.Error(w, "id is empty" + err.Error(), http.StatusBadRequest)
			return
		}

		if requestData.Name == "" {
			http.Error(w, "name is empty", http.StatusBadRequest)

			return
		}

		if requestData.Amount == "" {
			http.Error(w, "amount is empty", http.StatusBadRequest)

			return
		}

		if _, ok := db.Get(requestData.Id); ok == nil {
			http.Error(w, "db error: can not create account (account exists)", http.StatusInternalServerError)
			return
		}

		account := models.Account{}
		account.Init(requestData.Id, requestData.Name)
		account.SetAmount(requestData.Amount)

		if ok := db.Set(account.GetId(), &account); !ok {
			http.Error(w, "db error: can not create account", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	})
}