package main

import (
	"log"
	"net/http"
	"path"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	p := path.Dir("index.html")

	w.Header().Set("Content-type", "text/html")
	http.ServeFile(w, r, p)
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {

	handleRequests()

}
