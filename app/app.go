package app

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/maxkucher/secret-sharing-backend/public"
	"net/http"
)

type App struct {
	SecretService SecretService
}

func (app *App) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Wow, it works!"))
}

func (app *App) PostSecretHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var createSecretDTO public.CreateSecretDTO
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&createSecretDTO); err != nil {
		response, _ := json.Marshal(map[string]string{
			"error": err.Error(),
		})
		w.WriteHeader(400)
		w.Write(response)
		return
	}
	secretId, err := app.SecretService.SaveSecret(createSecretDTO.PlainString)
	if err != nil {
		response, _ := json.Marshal(map[string]string{
			"error": err.Error(),
		})
		w.WriteHeader(500)
		w.Write(response)
		return
	}
	response := public.CreateSecretDTOResponse{Id: secretId}
	responseBytes, _ := json.Marshal(response)
	w.WriteHeader(200)
	w.Write(responseBytes)
	return

}

func (app *App) GetSecretHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	secret, err := app.SecretService.LoadSecrets(id)
	if err != nil {
		response, _ := json.Marshal(map[string]string{
			"error": err.Error(),
		})
		w.WriteHeader(500)
		w.Write(response)
		return
	}
	if secret == "" {
		response, _ := json.Marshal(map[string]string{
			"error": "Secret not found",
		})
		w.WriteHeader(404)
		w.Write(response)
		return
	}
	response := public.GetSecretResponse{Data: secret}
	responseBytes, _ := json.Marshal(response)
	w.WriteHeader(200)
	w.Write(responseBytes)
	w.Header().Set("Content-Type", "application/json")
	return
}
