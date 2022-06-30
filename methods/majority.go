package methods

import (
	"sort"

	"github.com/cazicbor/hello-democracy/model"
)

func TwoRoundSystem(voters []model.Votant, c []model.Candidat) model.ResultList {
	// In this map, the key is the candidate
	// the value is the number of votes obtained by the candidate
	count := make(map[model.Candidat]int)
	iterationNumber := 1

	secondRoundList := make([]model.Candidat, 0)
	res1 := majorityProcedure(voters, c, count)

	if len(res1) == 1 {
		return res1
	} else {
		iterationNumber++
		for i := range res1 {
			secondRoundList = append(secondRoundList, res1[i].Cand)
			finalRes := TwoRoundSystem(voters, secondRoundList)
			return finalRes
		}
	}
	return model.ResultList{}
}

func majorityProcedure(voters []model.Votant, candidates []model.Candidat, res map[model.Candidat]int) model.ResultList {
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
			resPair = append(resPair, resultList[0], resultList[1])
			return resPair
		}
	}
	return resPair
}
