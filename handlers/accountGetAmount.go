package handlers

import (
	"net/http"
	"strconv"
	"../modules/storage"
	"encoding/json"
	"github.com/gorilla/mux"
)

func GetAmount(db storage.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "wrong http method. GET required", http.StatusBadRequest)
			return
		}

		vars := mux.Vars(r)

		id,err := strconv.Atoi(vars["id"])
		if id <=0 || err !=nil {
			http.Error(w, "id is empty", http.StatusBadRequest)
			return
		}

		defer r.Body.Close()

		account,ok := db.Get(id)
		if  ok != nil {
			http.Error(w, "db error: " + ok.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(account.Amount.String())
	})
}