package methods

import "github.com/cazicbor/hello-democracy/model"

func Copeland(voters []model.Voter, candidates []model.Candidate) (model.Candidate, int) {
	numberOfDuels := len(candidates) - 1
	candidat, result := CondorcetVoteRound(voters, candidates)
	duelLost := numberOfDuels - result.DuelsWon

	// result.DuelsWon - duelLost corresponds to the Copeland score
	return candidat, result.DuelsWon - duelLost

}
