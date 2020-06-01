// all go executable need a main package
package main

import (
	"log"      // for logging any errors
	"net/http" // for writing rest api

	"github.com/gorilla/mux"
)

func get(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Welcome to the home page!\n")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Successfully retrieved - Welcome to the home page!"}`))
}

func post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "Successfully created!"}`))
}

func put(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(`{"message": "Successfully updated!"}`))
}

func delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Successfully deleted!"}`))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "not found :("}`))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", get).Methods("GET")
	r.HandleFunc("/", post).Methods("POST")
	r.HandleFunc("/", put).Methods("PUT")
	r.HandleFunc("/", delete).Methods("DELETE")
	r.HandleFunc("/", notFound)

	log.Fatal(http.ListenAndServe(":8085", r))
}
