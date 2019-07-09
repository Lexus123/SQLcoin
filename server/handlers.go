package server

import (
	"encoding/json"
	"html/template"
	"net/http"
	"sqlcoin/services/blockreader"
	"sqlcoin/services/databaser"
	"sqlcoin/services/errorchecker"
	"sqlcoin/services/fileopener"

	"github.com/gorilla/mux"
)

type IndexData struct {
	PageTitle string
}

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	data := IndexData{
		PageTitle: "Homepage",
	}

	tmpl.Execute(w, data)
}

func Analyse(w http.ResponseWriter, r *http.Request) {
	_, handler, err := r.FormFile("analysefile")
	errorchecker.CheckFileError(err)
	blkFile := fileopener.OpenFile(handler.Filename)
	blockreader.StartReading(blkFile)
}

func MainIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/main.html"))

	data := IndexData{
		PageTitle: "Main page",
	}

	tmpl.Execute(w, data)
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
