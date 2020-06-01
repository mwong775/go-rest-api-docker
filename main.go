// all go executable need a main package
package main

import (
	"fmt"
	"log"      // for logging any errors
	"net/http" // for writing rest api

	"github.com/gorilla/mux"
)

type event struct {
	ID          string `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

type allEvents []event

var events = allEvents{
	{
		ID:          "1",
		Title:       "Introduction to Golang",
		Description: "Learn how golang works and try it out!",
	},
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page!\n")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "response to homepage request - Hello World!"}`))
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
	r.HandleFunc("/", home)
	r.HandleFunc("/", post).Methods(http.MethodPost)
	r.HandleFunc("/", put).Methods(http.MethodPut)
	r.HandleFunc("/", delete).Methods(http.MethodDelete)
	r.HandleFunc("/", notFound)

	log.Fatal(http.ListenAndServe(":8085", r))
}
