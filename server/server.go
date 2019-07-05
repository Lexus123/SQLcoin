package server

import (
	"log"
	"net/http"
)

/*
Host ...
*/
func Host() {
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":12345", router))
}
