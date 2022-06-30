package methods

import (
	"fmt"
	"sort"

	"github.com/cazicbor/hello-democracy/model"
)

func ApprovalMethod(voters []model.Votant, candidates []model.Candidat) model.ResultList {
	count := make(map[model.Candidat]int)

	secondRoundList := make([]model.Candidat, 0)

	res1 := approvalProcedure(voters, candidates, count)
	// Case where a candidate is elected in the first round
	if len(res1) == 1 {
		return res1
	} else {
		for i := range res1 {
			secondRoundList = append(secondRoundList, res1[i].Cand)
			finalRes := ApprovalMethod(voters, secondRoundList)
			return finalRes
		}
	}
	return model.ResultList{}
}

func approvalProcedure(voters []model.Votant, candidates []model.Candidat, res map[model.Candidat]int) model.ResultList {
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
				TotalPoint: v,
			},
		})
	}
	fmt.Println("result per candidate: ", resultList)

	sort.Sort(sort.Reverse(resultList))

	for cand := range resultList {
		percentage := (resultList[cand].Res.TotalPoint * 100) / total
		if percentage > 50 {
			resPair = append(resPair, model.ResultPerCandidate{
				Cand: resultList[0].Cand,
				Res: model.Result{
					TotalPoint: resultList[0].Res.TotalPoint,
				},
			})
			return resPair
		} else {
			for i := 0; i < len(candidates)-1; i++ {
				resPair = append(resPair, resultList[i])
				return resPair
			}
		}
	}
	return resPair
}
