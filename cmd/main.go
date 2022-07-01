package main

import (
	"fmt"
	"time"

	"github.com/cazicbor/hello-democracy/methods"
	"github.com/cazicbor/hello-democracy/mocks"
)

func main() {
	fmt.Println("-------------")
	now := time.Now()
	totalVotants := 10000000
	// generate candidats, TO DO : associate voter and vote to gain more time
	candidats := mocks.CreateCandidatsByName([]string{"LeBron", "Steph", "KD", "Kyrie", "Joker", "Giannis", "Tatum", "Luka", "Embiid", "Butler"})
	// generate random votants
	votants := mocks.AsyncGenerateRandomVotants(totalVotants)
	// votants := mocks.GenerateRandomVotants(totalVotants)
	fmt.Println("af mock voter : ", time.Since(now))
	now = time.Now()
	// simulate vote
	mocks.SimulateRandomVotes(votants, candidats)
	fmt.Println("af mock votes : ", time.Since(now))
	now = time.Now()
	// compare results
	resultMajority := methods.TwoRoundSystem(votants, candidats)
	winnerCondorcet, resultCondorcet := methods.CondorcetVoteRound(votants, candidats)
	resultApproval := methods.ApprovalMethod(votants, candidats)
	winnerCopeland, scoreCopeland := methods.Copeland(votants, candidats)
	fmt.Println("-------------")
	fmt.Println("vainqueur majorité à 2 tours : ", resultMajority)
	fmt.Println("-------------")
	fmt.Println("vainqueur de condorcet : ", winnerCondorcet)
	fmt.Println("résultat condorcet : ", resultCondorcet)
	fmt.Println("-------------")
	fmt.Println("gagnant par approbation : ", resultApproval)
	fmt.Println("-------------")
	fmt.Println("gagnant copeland : ", winnerCopeland)
	fmt.Println("score de copeland : ", scoreCopeland)
	fmt.Println("af calc : ", time.Since(now))
}
