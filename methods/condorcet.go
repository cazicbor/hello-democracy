package methods

import (
	"fmt"
	"time"

	"github.com/cazicbor/hello-democracy/model"
)

func CondorcetVoteRound(votes []model.Voter, candidates []model.Candidate) (model.Candidate, model.Result) {
	now := time.Now()
	possiblePairDiff, resultList := model.GeneratePossiblePairDiffsAndResultList(candidates)
	// get score between pairs for each votant
	fmt.Println("---- before setPossiblePairDiffValue", time.Since(now))
	model.SetPossiblePairDiffValue(possiblePairDiff, votes)
	fmt.Println("--- after setPossiblePairDiffValue", time.Since(now))
	now = time.Now()
	// Set result values
	model.SetResultListValue(possiblePairDiff, resultList)
	now = time.Now()
	// Get winner
	var result model.Result = model.Result{}
	var candidat model.Candidate = model.Candidate{}
	for r, v := range resultList {
		if result.DuelsWon < v.DuelsWon {
			result = *v
			candidat = r
		}
	}
	return candidat, result
}
