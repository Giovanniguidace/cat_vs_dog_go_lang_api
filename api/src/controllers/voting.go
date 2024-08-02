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
)

// CreateVote receive data voting and create - 1 for cat or 2 for dog
func CreateVote(w http.ResponseWriter, r *http.Request) {
	parameters, err := io.ReadAll(r.Body)
	if err != nil {
		responses.ErrorMessage(w, r, http.StatusUnprocessableEntity, err)
		return
	}

	var vote models.Vote

	if err = json.Unmarshal(parameters, &vote); err != nil {
		responses.ErrorMessage(w, r, http.StatusBadRequest, err)
		return
	}

	if err = vote.Validation(); err != nil {
		responses.ErrorMessage(w, r, http.StatusBadRequest, err)
		return
	}

	mongo_client, err := database.MongoConnect()

	if err != nil {
		responses.ErrorMessage(w, r, http.StatusInternalServerError, err)
		return
	}
	defer mongo_client.Disconnect(context.TODO())

	_, err = repositories.InsertVote(vote, mongo_client)
	if err != nil {
		responses.ErrorMessage(w, r, http.StatusInternalServerError, err)
		return
	}

	responses.SuccessMessage(w, r, http.StatusCreated, "vote created with success")
}
