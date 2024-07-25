package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"context"
	"encoding/json"
	"io"
	"net/http"
	_ "strconv"

	_ "github.com/gorilla/mux"
)

// CreateVote receive data voting and create - 1 for cat or 2 for dog
func CreateVote(w http.ResponseWriter, r *http.Request) {
	parameters, err := io.ReadAll(r.Body)
	if err != nil {
		responses.ErrorMessage(w, http.StatusUnprocessableEntity, err)
		return
	}

	var vote models.Vote

	if err = json.Unmarshal(parameters, &vote); err != nil {
		responses.ErrorMessage(w, http.StatusBadRequest, err)
		return
	}

	if err = vote.Validation(); err != nil {
		responses.ErrorMessage(w, http.StatusBadRequest, err)
		return
	}

	mongo_client, err := database.MongoConnect()

	if err != nil {
		responses.ErrorMessage(w, http.StatusInternalServerError, err)
		return
	}
	defer mongo_client.Disconnect(context.TODO())

	_, err = repositories.InsertVote(vote, mongo_client)
	if err != nil {
		responses.ErrorMessage(w, http.StatusInternalServerError, err)
		return
	}

	responses.SuccessMessage(w, http.StatusCreated, "vote created with success")
}
