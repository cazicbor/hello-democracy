package model

import (
	"fmt"
	"sync"

	"github.com/cazicbor/hello-democracy/config"
)

func SetPossiblePairDiffValue(possiblePairDiff map[Pair]Diff, votes []Votant) {
	var wg sync.WaitGroup
	ch := make(chan ResultDto, len(possiblePairDiff))
	// will create len(possibleCase) go routines
	for cPair := range possiblePairDiff {
		wg.Add(1)
		go func(cPair Pair, votes []Votant) {
			defer wg.Done()
			ch <- NewResultDto(GetDiff(cPair.L, cPair.R, votes))
		}(cPair, votes)
	}
	wg.Wait()
	close(ch)

	for v := range ch {
		possiblePairDiff[Pair{L: v.L, R: v.R}] = v.D
	}
	fmt.Println("total go routine on main :", len(possiblePairDiff))
}

func SetResultListValue(possiblePairDiff map[Pair]Diff, resultList map[Candidat]*Result) {
	for candidats, diff := range possiblePairDiff {
		if diff.Ldiff >= diff.Rdiff {
			resultList[candidats.L].DuelWon++
			resultList[candidats.L].TotalPoint += diff.Ldiff
			resultList[candidats.L].WonAgainst = append(resultList[candidats.L].WonAgainst, candidats.R)

		} else {
			resultList[candidats.R].DuelWon++
			resultList[candidats.R].TotalPoint += diff.Rdiff
			resultList[candidats.R].WonAgainst = append(resultList[candidats.R].WonAgainst, candidats.L)
		}
	}
}

func GetDiff(ref, adv Candidat, votes []Votant) (Diff, Candidat, Candidat) {
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
		go func(j int, subSlice []Votant) {
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
	// return pair with result for the channel
	return diff, ref, adv
}

func GeneratePossiblePairDiffsAndResultList(candidats []Candidat) (map[Pair]Diff, map[Candidat]*Result) {
	possiblePairDiff := make(map[Pair]Diff)
	duelsWonPerCandidate := make(map[Candidat]*Result)
	for i := range candidats {
		for j := i + 1; j < len(candidats); j++ {
			possiblePairDiff[Pair{L: candidats[i], R: candidats[j]}] = Diff{}
		}
		duelsWonPerCandidate[candidats[i]] = &Result{DuelWon: 0, TotalPoint: 0}
	}

	fmt.Println("cas possible", len(possiblePairDiff))
	return possiblePairDiff, duelsWonPerCandidate
}

// Useful funcs needed to sort maps

func (p ResultList) Len() int {
	return len(p)
}

func (p ResultList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p ResultList) Less(i, j int) bool {
	return p[i].Res.TotalPoint < p[j].Res.TotalPoint
}
