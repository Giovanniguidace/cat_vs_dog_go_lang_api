package models

import (
	"errors"
	"time"

)

type Vote struct{
	VoteId uint64 `json:"vote_id,omitempty"`
	DateCreation time.Time `json:"date_creation,omitempty"`
}

type Results struct{
	Cat uint64 `json:"dog,omitempty"`
	Dog uint64 `json:"dog,omitempty"`
}


// Validation check and validate returned values.
func (vote *Vote) Validation() error {

	if vote.VoteId != 1 && vote.VoteId != 2 {
		return errors.New("invalid field name or value")
	}

	return nil
}
