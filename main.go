package main

import (
	"net/http"

	"github.com/bangaping27/go-restapi-mux/controllers/productcontroler"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/products", productcontroler.Index).Methods("GET")
	r.HandleFunc("/products/{id}", productcontroler.Show).Methods("GET")
	r.HandleFunc("/products", productcontroler.Create).Methods("POST")
	r.HandleFunc("/products/{id}", productcontroler.Update).Methods("PUT")
	r.HandleFunc("/products/{id}", productcontroler.Delete).Methods("DELETE")

	http.ListenAndServe(":3000", r)

}
