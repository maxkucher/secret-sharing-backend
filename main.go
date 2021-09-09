package main

import (
	"github.com/maxkucher/secret-sharing-backend/app"
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
	secretService := app.FSSecretService{
		FilePath: secretsFilePath,
	}
	server := app.App{SecretService: &secretService}
	router := mux.NewRouter()
	router.HandleFunc("/healthcheck", server.HealthCheckHandler).Methods("GET", "POST")
	router.HandleFunc("/secretHandler", server.PostSecretHandler).Methods("POST")
	router.HandleFunc("/secretHandler/{id}", server.GetSecretHandler).Methods("GET")
	http.ListenAndServe(":8080", router)
}
