package services

import "fmt"

// VoteCatOrDog receive vote and insert vote into database
func VoteCatOrDog(vote_id uint64) (error){
	fmt.Sprint("%d", vote_id)
	return nil
}