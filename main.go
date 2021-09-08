package main

import (
	"log"
	"net/http"
	"os"
)
import "github.com/gorilla/mux"

func main() {
	secretsFilePath := os.Getenv("DATA_FILE_PATH")
	file, err := os.Create(secretsFilePath)
	if err != nil {
		log.Fatalf("Cannot initialize secrets file %s", err)
	}
	_, err = file.WriteString("{}")
	if err != nil {
		log.Fatalf("Error initializing secrets file with an empty map %s", err)
	}
	secretService := FSSecretService{
		FilePath: secretsFilePath,
	}
	app := App{secretService: &secretService}
	router := mux.NewRouter()
	router.HandleFunc("/healthcheck", app.healthCheckHandler).Methods("GET", "POST")
	router.HandleFunc("/secretHandler", app.postSecretHandler).Methods("POST")
	router.HandleFunc("/secretHandler/{id}", app.getSecretHandler).Methods("GET")
	http.ListenAndServe(":8080", router)
}
