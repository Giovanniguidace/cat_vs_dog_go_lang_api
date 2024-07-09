package controllers

import (
	"api/src/models"
	_ "api/src/services"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	_ "strconv"

	_ "github.com/gorilla/mux"
)

// CreateVote receive data voting and create - 1 for cat or 2 for dog
func CreateVote(w http.ResponseWriter, r *http.Request){
	parameters, err := io.ReadAll(r.Body)

	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)

	if err != nil {
		json.NewEncoder(w).Encode(err)
	}

	var vote models.Vote

	if err = json.Unmarshal(parameters, &vote); err != nil{
		json.NewEncoder(w).Encode(err)
		return
	}

	if err = vote.Validation(); err != nil {
		fmt.Printf("%d", err)
		json.NewEncoder(w).Encode(err)
		return
	}

	fmt.Println(vote.VoteId)
}