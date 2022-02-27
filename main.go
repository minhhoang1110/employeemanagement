package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/minhhoang1110/employeemanagement/api"
)

func main() {
	r := mux.NewRouter()
	api.Register(r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
