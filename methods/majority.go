package methods

import (
	"fmt"
	"sort"

	"github.com/cazicbor/hello-democracy/model"
)

func TwoRoundSystem(voters []model.Votant, c []model.Candidat) model.ResultList {
	fmt.Println(c)
	count := make(map[model.Candidat]int) //key: candidate, value: number of votes got by the candidate
	secondRoundList := make([]model.Candidat, 0)
	res1 := majorityProcedure(voters, c, count)
	if len(res1) == 1 {
		return res1
	} else {
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
	i := 0
	total := 0

	for i, v := range voters {
		// We first deal with the case where the voter does not want to vote
		// if !v.Votes {
		// 	fmt.Println("Cette personne a choisi de ne pas voter.")
		// 	continue
		// }
		for _, alternative := range candidates {
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
		i++
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
