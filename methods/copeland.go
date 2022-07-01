package methods

import "github.com/cazicbor/hello-democracy/model"

func Copeland(voters []model.Votant, candidates []model.Candidat) (model.Candidat, int) {
	numberOfDuels := len(candidates) - 1
	candidat, result := CondorcetVoteRound(voters, candidates)
	duelLost := numberOfDuels - result.DuelWon
	copelandScore := result.DuelWon - duelLost

	return candidat, copelandScore

}
