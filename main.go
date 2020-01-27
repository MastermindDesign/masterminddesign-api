package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/FR0NK3NST33N/masterminddesign-api/project"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index).Methods("GET")
	router.HandleFunc("/api", index).Methods("GET")
	router.HandleFunc("/api/projects", project.Index).Methods("GET")
	//router.HandleFunc("/api/contact", controllers.Contact)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, friend...")
}
