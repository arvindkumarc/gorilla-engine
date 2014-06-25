package main

import (
	"net/http"
	"log"
	"github.com/arvindkumarc/gorilla-engine/controllers"
	"github.com/gorilla/mux"
)

func main() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/hello", controllers.HelloWorldHandler).Methods("GET")

	http.Handle("/", rtr)

	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}
