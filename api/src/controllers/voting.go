package controllers

import (
	"fmt"
	"net/http"
)

// CreateVote receive data voting and create
func CreateVote(w http.ResponseWriter, r *http.Request){
	fmt.Printf("API Starting on port 5000")
}