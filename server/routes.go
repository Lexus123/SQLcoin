package server

import (
	"net/http"
)

/*
Route ...
*/
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

/*
Routes ...
*/
type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"Analyse",
		"POST",
		"/analyse",
		Analyse,
	},
	Route{
		"MainIndex",
		"GET",
		"/main",
		MainIndex,
	},
	Route{
		"InputsIndex",
		"GET",
		"/inputs",
		InputsIndex,
	},
	Route{
		"OutputsIndex",
		"GET",
		"/outputs",
		OutputsIndex,
	},
	Route{
		"TxsIndex",
		"GET",
		"/txs",
		TxsIndex,
	},
	Route{
		"GetSingleOutput",
		"GET",
		"/output/{output}",
		GetSingleOutput,
	},
	Route{
		"GetSingleTx",
		"GET",
		"/tx/{txHash}",
		GetSingleTx,
	},

	// API routes from here on
	Route{
		"CountInputs",
		"GET",
		"/api/inputs/count",
		CountInputs,
	},
	Route{
		"GetAllInputs",
		"GET",
		"/api/inputs",
		GetAllInputs,
	},
}
