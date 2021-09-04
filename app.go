package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type App struct {
	secretService SecretService
}

type CreateSecretDTO struct {
	PlainString string `json:"plain_string"`
}

type CreateSecretDTOResponse struct {
	Id string `json:"id"`
}

type GetSecretResponse struct {
	Data string `json:"data"`
}

func (app *App) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Wow, it works!"))
}

func (app *App) postSecretHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var createSecretDTO CreateSecretDTO
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&createSecretDTO); err != nil {
		response, _ := json.Marshal(map[string]string{
			"error": err.Error(),
		})
		w.WriteHeader(400)
		w.Write(response)
		return
	}
	secretId := app.secretService.SaveSecret(createSecretDTO.PlainString)
	response := CreateSecretDTOResponse{Id: secretId}
	responseBytes, _ := json.Marshal(response)
	w.WriteHeader(200)
	w.Write(responseBytes)
	return

}

func (app *App) getSecretHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	secret := app.secretService.LoadSecrets(id)
	response := GetSecretResponse{Data: secret}
	responseBytes, _ := json.Marshal(response)
	w.WriteHeader(200)
	w.Write(responseBytes)
	w.Header().Set("Content-Type", "application/json")
	return
}
