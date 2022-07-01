package methods

import "github.com/cazicbor/hello-democracy/model"

func Copeland(voters []model.Voter, candidates []model.Candidate) (model.Candidate, int) {
	numberOfDuels := len(candidates) - 1
	candidat, result := CondorcetVoteRound(voters, candidates)
	duelLost := numberOfDuels - result.DuelsWon
	copelandScore := result.DuelsWon - duelLost

	return candidat, copelandScore

}
