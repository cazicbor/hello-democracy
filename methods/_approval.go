package methods

func ApprovalMethod(voters []Voter, candidates []Person) PairList {
	count := make(map[Person]int)
	secondRoundList := make([]Person, 0)
	for _, v := range voters {
		for _, alternative := range candidates {
			if v.Choice == alternative.ID {
				res[alternative] += 1
				total++
			}
		}
	}
}

func approvalProcedure(voters []Voter, candidates []Person, res map[Person]int)
