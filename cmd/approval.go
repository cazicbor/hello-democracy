package cmd

import (
	"fmt"
	"time"

	"github.com/cazicbor/hello-democracy/config"
	"github.com/cazicbor/hello-democracy/methods"
	"github.com/cazicbor/hello-democracy/mocks"
	"github.com/spf13/cobra"
)

var approvalCmd = &cobra.Command{
	Use:   "approval",
	Short: "Use approval voting method",
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

		approvalResult := methods.ApprovalMethod(voters, candidates)

		fmt.Println("-------------")
		fmt.Println("Approval winner : ", approvalResult)
		fmt.Println("-------------")

		fmt.Println("af calc : ", time.Since(now))
	},
}
