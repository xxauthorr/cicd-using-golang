package main

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", Output)
	http.Handle("/", r)

	http.ListenAndServe(":8080", r)
}

func Output(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Random number is : %v", randomNumbers())
}

func randomNumbers() int {
	return rand.Intn(1000)
}
