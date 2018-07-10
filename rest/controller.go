package rest

import (
	"net/http"
	"go-hello-skeleton/models"
	"encoding/json"
)


func sendBadRequest(w http.ResponseWriter){

}

func PutPet(w http.ResponseWriter,r *http.Request){
	var req models.PutPetRequest

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&req)
	if err != nil {
		sendBadRequest(w)
		return
	}

}

func GetPet(w http.ResponseWriter,r *http.Request){

}