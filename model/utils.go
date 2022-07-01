package model

import (
	"fmt"
	"sync"

	"github.com/cazicbor/hello-democracy/config"
)

func SetPossiblePairDiffValue(possiblePairDiff map[Pair]Diff, votes []Voter) {
	var wg sync.WaitGroup
	ch := make(chan ResultDto, len(possiblePairDiff))

	for cPair := range possiblePairDiff {
		wg.Add(1)
		go func(cPair Pair, votes []Voter) {
			defer wg.Done()
			ch <- NewResultDto(GetDiff(cPair.L, cPair.R, votes))
		}(cPair, votes)
	}
	wg.Wait()
	close(ch)

	for v := range ch {
		possiblePairDiff[Pair{L: v.L, R: v.R}] = v.D
	}
	fmt.Println("total goroutines on main:", len(possiblePairDiff))
}

func SetResultListValue(possiblePairDiff map[Pair]Diff, resultList map[Candidate]*Result) {
	for candidates, diff := range possiblePairDiff {
		if diff.Ldiff >= diff.Rdiff {
			resultList[candidates.L].DuelsWon++
			resultList[candidates.L].TotalPoints += diff.Ldiff
			resultList[candidates.L].WonAgainst = append(resultList[candidates.L].WonAgainst, candidates.R)

		} else {
			resultList[candidates.R].DuelsWon++
			resultList[candidates.R].TotalPoints += diff.Rdiff
			resultList[candidates.R].WonAgainst = append(resultList[candidates.R].WonAgainst, candidates.L)
		}
	}
}

func GetDiff(ref, adv Candidate, votes []Voter) (Diff, Candidate, Candidate) {
	diff := Diff{}
	var diffWg sync.WaitGroup
	total := len(votes)
	diffCh := make(chan Diff, total)
	stackSizePerGoRoutine := config.SubSliceSize / 10
	// will create (len(votes) / config.SubSliiceSize) go routines
	for j := total; j >= 0; j -= stackSizePerGoRoutine { // in bce style to gain precious millis
		limit := j - stackSizePerGoRoutine
		if limit < 0 {
			limit = 0
		}
		subSlice := votes[limit:j]
		diffWg.Add(1)
		go func(j int, subSlice []Voter) {
			defer diffWg.Done()
			currentDiff := Diff{}
			subSliceMax := len(subSlice) - 1
			for i := subSliceMax; i >= 0; i-- {
				if subSlice[i].Vote.IndexOf(ref) < subSlice[i].Vote.IndexOf(adv) {
					currentDiff.Ldiff += 1
				} else {
					currentDiff.Rdiff += 1
				}
			}
			diffCh <- currentDiff
		}(j, subSlice)
	}

	diffWg.Wait()
	close(diffCh)

	for v := range diffCh {
		diff.Ldiff += v.Ldiff
		diff.Rdiff += v.Rdiff
	}
	// Returns a pair with result of the channel
	return diff, ref, adv
}

func GeneratePossiblePairDiffsAndResultList(candidats []Candidate) (map[Pair]Diff, map[Candidate]*Result) {
	possiblePairDiff := make(map[Pair]Diff)
	duelsWonPerCandidate := make(map[Candidate]*Result)
	for i := range candidats {
		for j := i + 1; j < len(candidats); j++ {
			possiblePairDiff[Pair{L: candidats[i], R: candidats[j]}] = Diff{}
		}
		duelsWonPerCandidate[candidats[i]] = &Result{DuelsWon: 0, TotalPoints: 0}
	}

	fmt.Println("possible cases: ", len(possiblePairDiff))
	return possiblePairDiff, duelsWonPerCandidate
}

// Useful functions needed to sort maps
func (p ResultList) Len() int {
	return len(p)
}

func (p ResultList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p ResultList) Less(i, j int) bool {
	return p[i].Res.TotalPoints < p[j].Res.TotalPoints
}

func Ballot(votingMethod string, resultList ResultList, total int, resPair ResultList, candidates []Candidate) (ResultList, bool, ResultList) {
	for cand := range resultList {
		percentage := (resultList[cand].Res.TotalPoints * 100) / total
		if percentage > 50 {
			resPair = append(resPair, ResultPerCandidate{
				Cand: resultList[0].Cand,
				Res: Result{
					TotalPoints: resultList[0].Res.TotalPoints,
				},
			})
			return nil, true, resPair
		} else {
			for i := 0; i < len(candidates)-1; i++ {
				// Case disjunction: majority run-off and approval method
				if votingMethod == "approval" {
					resPair = append(resPair, resultList[i])
				} else if votingMethod == "majority" {
					resPair = append(resPair, resultList[0], resultList[1])
				}
				return nil, true, resPair
			}
		}
	}
	return resPair, false, nil
}
