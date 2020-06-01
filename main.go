// all go executable need a main package
package main

import (
	"fmt"
	"log"      // for logging any errors
	"net/http" // for writing rest api

	"github.com/gorilla/mux"
)

func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "GET called"}`))
}

func post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "POST called"}`))
}

func put(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(`{"message": "PUT called"}`))
}

func delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "DELETE called"}`))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "not found :("}`))
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{"Type":"%s","message":"Hello Distributed Systems"}`, r.Method)))
}

func msg(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	// fmt.Println(query)
	if r.Method == "POST" {
		msgParam, ok := query["msg"] // gets string[] of length 1
		// if msg query param exists
		if ok {
			msg := msgParam[0] // get string from array
			// fmt.Println(msg)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(fmt.Sprintf(`{"Type":"%s","message":"%s"}`, r.Method, msg)))
			// no msg query, return 405
		} else {
			// fmt.Println("no msg")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(`This method is unsupported.`))
		}
	} else if r.Method == "GET" && len(query) == 0 {
		w.Write([]byte(fmt.Sprintf(`{"Type": "%s","message":"Get Message Received"}`, r.Method)))
	} else { // 405
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(`This method is unsupported.`))
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", get).Methods(http.MethodGet)
	r.HandleFunc("/", post).Methods(http.MethodPost)
	r.HandleFunc("/", put).Methods(http.MethodPut)
	r.HandleFunc("/", delete).Methods(http.MethodDelete)
	r.HandleFunc("/", notFound)

	r.HandleFunc("/distributed", hello).Methods(http.MethodGet)
	r.HandleFunc("/systems", msg).Methods(http.MethodGet, http.MethodPost)
	log.Fatal(http.ListenAndServe(":8085", r))
}
