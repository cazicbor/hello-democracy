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
	// set result values

	model.SetResultListValue(possiblePairDiff, resultList)
	now = time.Now()
	// get winner
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

// func CondorcetWinner(voters []Voter, candidates []Person) (Person, int) {
// 	countWins := make(map[Person]int)
// 	pairs := make([][]Person, 0)

// 	for i, a := range candidates {
// 		for j, b := 0; b < len(candidates[i+1:]); j++ {
// 			pairs[i] = append(pairs[i], a, b)
// 		}
// 	}
// 	for _, p := range pairs {
// 		for _, v := range voters {
// 			for i := range v.Prefs {
// 				if p[0] == v.Prefs[i] {
// 					countWins[p[0]] += 1
// 				}
// 			}
// 		}
// 	}
// 	for candid, score := range countWins {
// 		if countWins[candid] == len(candidates)-1 {
// 			return candid, score
// 		} else {
// 			fmt.Println("Pas de gagnant de Condorcet")
// 		}
// 	}
// 	return Person{}, -1
// }

//le i := 0 est inutile, et le dernier for aussi tu peux le kicker, vu que la taille de l'array sera 1 ou > 1 tu prends direct arr[0] et arr[1] si 2eme tour
