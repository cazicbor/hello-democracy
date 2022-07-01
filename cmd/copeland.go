package cmd

import (
	"fmt"
	"time"

	"github.com/cazicbor/hello-democracy/config"
	"github.com/cazicbor/hello-democracy/methods"
	"github.com/cazicbor/hello-democracy/mocks"
	"github.com/spf13/cobra"
)

var copelandCmd = &cobra.Command{
	Use:   "copeland",
	Short: "Use copeland's voting method",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("-------------")
		now := time.Now()
		candidates := mocks.CreateCandidatesByName([]string{"LeBron", "Steph", "KD", "Kyrie", "Joker", "Giannis", "Tatum", "Luka", "Embiid", "Butler"})

		// Generatation of random voters
		voters := mocks.AsyncGenerateRandomVoters(config.TotalVoters)
		// voters := mocks.GenerateRandomVoters(config.TotalVoters)
		fmt.Println("af mock voter : ", time.Since(now))
		now = time.Now()

		// Simulate vote
		mocks.SimulateRandomVotes(voters, candidates)
		fmt.Println("af mock votes : ", time.Since(now))
		now = time.Now()

		copelandWinner, copelandScore := methods.Copeland(voters, candidates)

		fmt.Println("-------------")
		fmt.Println("Copeland winner : ", copelandWinner)
		fmt.Println("Copeland score: ", copelandScore)
		fmt.Println("-------------")

		fmt.Println("af calc : ", time.Since(now))
	},
}
