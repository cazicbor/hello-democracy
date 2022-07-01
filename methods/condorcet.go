package methods

import (
	"fmt"
	"time"

	"github.com/cazicbor/hello-democracy/model"
)

func CondorcetVoteRound(votes []model.Votant, candidats []model.Candidat) (model.Candidat, model.Result) {
	now := time.Now()
	possiblePairDiff, resultList := model.GeneratePossiblePairDiffsAndResultList(candidats)
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
	var candidat model.Candidat = model.Candidat{}
	for r, v := range resultList {
		if result.DuelWon < v.DuelWon {
			result = *v
			candidat = r
		}
	}
	return candidat, result
}
