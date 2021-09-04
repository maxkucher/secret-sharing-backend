package main

import (
	"net/http"
)
import "github.com/gorilla/mux"

func main() {
	app := App{secretService: FSSecretService{}}
	router := mux.NewRouter()
	router.HandleFunc("/healthcheck", app.healthCheckHandler).Methods("GET", "POST")
	router.HandleFunc("/secretHandler", app.postSecretHandler).Methods("POST")
	router.HandleFunc("/secretHandler/{id}", app.getSecretHandler).Methods("GET")
	http.ListenAndServe(":8080", router)
}
