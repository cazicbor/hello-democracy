package methods

import (
	"sort"

	"github.com/cazicbor/hello-democracy/model"
)

func ApprovalMethod(voters []model.Voter, candidates []model.Candidate) model.ResultList {
	count := make(map[model.Candidate]int)

	secondRoundList := make([]model.Candidate, 0)

	firstRoundResult := approvalProcedure(voters, candidates, count)
	// Case where a candidate is elected in the first round
	if len(firstRoundResult) == 1 {
		return firstRoundResult
	} else {
		for i := range firstRoundResult {
			secondRoundList = append(secondRoundList, firstRoundResult[i].Cand)
			// Same model as for the majority run off: recursive call
			finalResult := ApprovalMethod(voters, secondRoundList)
			return finalResult
		}
	}
	return model.ResultList{}
}

func approvalProcedure(voters []model.Voter, candidates []model.Candidate, res map[model.Candidate]int) model.ResultList {
	resultList := make(model.ResultList, len(res))
	resPair := make(model.ResultList, 0)

	total := 0

	for _, v := range voters {
		for i, alternative := range candidates {
			if v.Vote.IndexOf(candidates[i]) == alternative.Id {
				res[alternative] += 1
				total++
			}
		}
	}
	for k, v := range res {
		resultList = append(resultList, model.ResultPerCandidate{
			Cand: k,
			Res: model.Result{
				TotalPoints: v,
			},
		})
	}

	sort.Sort(sort.Reverse(resultList))

	resPair, shouldReturn, returnValue := model.Ballot("approval", resultList, total, resPair, candidates)
	if shouldReturn {
		return returnValue
	}
	return resPair
}
