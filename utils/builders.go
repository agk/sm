package utils

import (
	"github.com/agk/specialist/c2/hw2/handlers"
	"github.com/gorilla/mux"
)

func BuildResource(router *mux.Router, prefix string) {
	
	router.HandleFunc(prefix, handlers.Grab).Methods("POST")

	router.HandleFunc(prefix, handlers.Solve).Methods("GET")
	
}