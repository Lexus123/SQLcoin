package server

import (
	"encoding/json"
	"html/template"
	"net/http"
	"time"

	"sqlcoin/services/blockreader"
	"sqlcoin/services/databaser"
	"sqlcoin/services/errorchecker"
	"sqlcoin/services/fileopener"

	"github.com/gorilla/mux"
)

/*
IndexData ...
*/
type IndexData struct {
	PageTitle string
}

/*
InputsData ...
*/
type InputsData struct {
	PageTitle string
	Inputs    []databaser.Input
}

/*
OutputsData ...
*/
type OutputsData struct {
	PageTitle string
	Outputs   []databaser.Output
}

/*
OutputData ...
*/
type OutputData struct {
	PageTitle string
	Output    databaser.Output
}

/*
TxData ...
*/
type TxData struct {
	PageTitle string
	Inputs    []databaser.Input
	Outputs   []databaser.Output
	TxHash    string
}

/*
TxsData ...
*/
type TxsData struct {
	PageTitle string
	Txs       []databaser.Tx
}

/*
Index ...
*/
func Index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	data := IndexData{
		PageTitle: "Homepage",
	}

	tmpl.Execute(w, data)
}

/*
Analyse ...
*/
func Analyse(w http.ResponseWriter, r *http.Request) {
	_, handler, err := r.FormFile("analysefile")
	errorchecker.CheckFileError(err)
	blkFile := fileopener.OpenFile(handler.Filename)
	blockreader.StartReading(blkFile)
}

// With templates

/*
MainIndex ...
*/
func MainIndex(w http.ResponseWriter, r *http.Request) {
	data := IndexData{
		PageTitle: "Main page",
	}

	tmpl := template.Must(template.ParseFiles("templates/main.html"))
	tmpl.Execute(w, data)
}

/*
InputsIndex ...
*/
func InputsIndex(w http.ResponseWriter, r *http.Request) {
	inputs := databaser.GetAllInputs()

	data := InputsData{
		PageTitle: "Inputs Page",
		Inputs:    inputs,
	}

	tmpl := template.Must(template.ParseFiles("templates/inputs.html"))
	tmpl.Execute(w, data)
}

/*
OutputsIndex ...
*/
func OutputsIndex(w http.ResponseWriter, r *http.Request) {
	outputs := databaser.GetAllOutputs()

	data := OutputsData{
		PageTitle: "Outputs Page",
		Outputs:   outputs,
	}

	tmpl := template.Must(template.ParseFiles("templates/outputs.html"))
	tmpl.Execute(w, data)
}

/*
TxsIndex ...
*/
func TxsIndex(w http.ResponseWriter, r *http.Request) {
	txs := databaser.GetAllTxs()

	data := TxsData{
		PageTitle: "Transactions Page",
		Txs:       txs,
	}

	tmpl := template.Must(template.ParseFiles("templates/txs.html"))
	tmpl.Execute(w, data)
}

/*
GetSingleOutput ...
*/
func GetSingleOutput(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	output := databaser.GetSingleOutput(vars["output"])

	data := OutputData{
		PageTitle: "Single Output Page",
		Output:    output,
	}

	var funcMap = template.FuncMap{
		"satsToWhole":     satsToWhole,
		"timestampToTime": timestampToTime,
	}

	tmpl := template.Must(template.New("single.html").Funcs(funcMap).ParseFiles("templates/output/single.html"))
	tmpl.Execute(w, data)
}

/*
GetSingleTx ...
*/
func GetSingleTx(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	txHash := databaser.GetSingleTx(vars["txHash"])

	data := TxData{
		PageTitle: "Single Tx Page",
		Inputs:    txHash.Inputs,
		Outputs:   txHash.Outputs,
		TxHash:    txHash.TxHash,
	}

	var funcMap = template.FuncMap{
		"satsToWhole":     satsToWhole,
		"timestampToTime": timestampToTime,
		"totalBtc":        totalBtc,
	}

	tmpl := template.Must(template.New("single.html").Funcs(funcMap).ParseFiles("templates/tx/single.html"))
	tmpl.Execute(w, data)
}

func satsToWhole(sats int64) float64 {
	return float64(sats) / 100000000
}

func timestampToTime(timestamp int64) time.Time {
	return time.Unix(timestamp, 0).UTC()
}

func totalBtc(arg interface{}) float64 {
	var totalBtc int64

	switch arg := arg.(type) {
	case []databaser.Output:
		for _, output := range arg {
			totalBtc += output.Amount
		}
	}

	return satsToWhole(totalBtc)
}

// API Handlers

/*
GetAllInputs ...
*/
func GetAllInputs(w http.ResponseWriter, r *http.Request) {
	inputs := databaser.GetAllInputs()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(inputs); err != nil {
		panic(err)
	}
}

/*
CountInputs ...
*/
func CountInputs(w http.ResponseWriter, r *http.Request) {
	countInputs := databaser.CountInputs()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(countInputs); err != nil {
		panic(err)
	}
}
