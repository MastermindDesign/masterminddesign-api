package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/FR0NK3NST33N/masterminddesign-api/project"
	"github.com/FR0NK3NST33N/masterminddesign-api/utils"
	"github.com/felixge/httpsnoop"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index).Methods("GET")
	router.HandleFunc("/api", index).Methods("GET")
	router.HandleFunc("/api/projects", project.Index).Methods("GET")

	loggedRouter := logRequestHandler(router)

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})

	//router.HandleFunc("/api/contact", controllers.Contact)
	//log.Fatal(http.ListenAndServe(":8080", loggedRouter))
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(headersOk, originsOk, methodsOk)(loggedRouter)))

}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, friend...")
}

func logRequestHandler(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ri := &utils.HTTPReqInfo{
			Method:    r.Method,
			Uri:       r.URL.String(),
			Referer:   r.Header.Get("Referer"),
			UserAgent: r.Header.Get("User-Agent"),
		}

		ri.Ipaddr = utils.RequestGetRemoteAddress(r)

		m := httpsnoop.CaptureMetrics(h, w, r)

		ri.Code = m.Code
		ri.Size = m.Written
		ri.Duration = m.Duration

		e, err := json.Marshal(ri)

		if err != nil {
			fmt.Println(err)
			return
		}

		// uncomment to start logging http requests to elasticsearch
		// turned off due to aws pricing concerns
		//utils.LogToElasticsearch("mastermind_logs", string(e))

		fmt.Println(string(e))

	}

	// http.HandlerFunc wraps a function so that it
	// implements http.Handler interface
	return http.HandlerFunc(fn)
}
