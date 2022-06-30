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
	candidats := mocks.CreateCandidatsByName([]string{"Messi", "Cr7", "Zidane", "Xavi", "Iniesta", "De Bruyne", "Mahrez", "Salah", "Mane", "Firmino"})
	// generate random votants
	votants := mocks.AsyncGenerateRandomVotants(totalVotants)
	//votants := mocks.GenerateRandomVotants(totalVotants)
	fmt.Println("af mock voter : ", time.Since(now))
	now = time.Now()
	// simulate vote
	mocks.SimulateRandomVotes(votants, candidats)
	fmt.Println("af mock votes : ", time.Since(now))
	now = time.Now()
	// compare results
	result := methods.TwoRoundSystem(votants, candidats)
	//candidat, result := methods.CondorcetVoteRound(votants, candidats)
	fmt.Println("-------------")
	//fmt.Println("gagnant : ", candidat)
	fmt.Println("result : ", result)
	fmt.Println("-------------")
	fmt.Println("af calc : ", time.Since(now))
}
