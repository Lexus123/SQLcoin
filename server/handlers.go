package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sqlcoin/services/databaser"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func BlockIndex(w http.ResponseWriter, r *http.Request) {
	blocks := databaser.GetAllBlocks()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(blocks); err != nil {
		panic(err)
	}
}

func BlockShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	block := databaser.GetSingleBlocks(vars["blockId"])

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(block); err != nil {
		panic(err)
	}
}
