package handlers

import (
	"../modules/storage"
	"../models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

		account := models.Account{}
		account.Init(requestData.Id, requestData.Name)

		if ok := account.SetAmount(requestData.Amount); ok != nil {
			http.Error(w, ok.Error(), http.StatusInternalServerError)
			return
		}

		if ok := db.Create(requestData.Id, &account); ok != nil {
			http.Error(w, ok.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	})
}