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
		"BlockIndex",
		"GET",
		"/blocks",
		BlockIndex,
	},
	Route{
		"BlockShow",
		"GET",
		"/block/{blockId}",
		BlockShow,
	},
}
