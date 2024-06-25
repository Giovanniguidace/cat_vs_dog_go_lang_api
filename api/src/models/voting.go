package models

import (
	"errors"
	"time"
)

type Vote struct{
	VoteId uint64 `json:"vote_id,omitempty"`
	DateCreation time.Time `json:"date_creation,omitempty"`
}


func (vote *Vote) Validation() error {

	if vote.VoteId != 1 && vote.VoteId != 2 {
		return errors.New("Vote Id must be 1 to vote in Cat or 2 to vote in Dog")
	}

	return nil
}