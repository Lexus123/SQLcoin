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
		"BlockIndex",
		"GET",
		"/api/blocks",
		BlockIndex,
	},
	Route{
		"BlockShow",
		"GET",
		"/api/block/{blockId}",
		BlockShow,
	},
}
