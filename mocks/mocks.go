package mocks

import (
	"math/rand"
	"strconv"
	"sync"
	"time"

	"github.com/cazicbor/hello-democracy/config"

	"github.com/cazicbor/hello-democracy/model"
)

func GenerateRandomVotants(number int) []model.Votant {
	votants := make([]model.Votant, 0, number)
	for i := 1; i <= number; i++ {
		votants = append(votants, model.NewVotant(i, "Votant-"+strconv.Itoa(i)))
	}
	return votants
}

func AsyncGenerateRandomVotants(total int) []model.Votant {
	votants := make([]model.Votant, 0, total)
	var wg sync.WaitGroup
	ch := make(chan []model.Votant, (total/config.SubSliceSize)+1)
	for j, k := 0, config.SubSliceSize; j+config.SubSliceSize < total+config.SubSliceSize; j += config.SubSliceSize {
		wg.Add(1)
		go func(j, k int, ch chan []model.Votant) {
			defer wg.Done()
			diff := total - j
			currentSubSliceSize := config.SubSliceSize
			if diff < config.SubSliceSize {
				currentSubSliceSize = diff
			}

			subSlice := make([]model.Votant, 0, currentSubSliceSize)
			for i, l := j, 0; l < currentSubSliceSize; i++ {
				subSlice = append(subSlice, model.NewVotant(i, "Votant-"+strconv.Itoa(i)))
				l++
			}

			ch <- subSlice
		}(j, k, ch)
		k = k + config.SubSliceSize
		if k > total {
			k = total
		}
	}
	wg.Wait()
	close(ch)

	for v := range ch {
		votants = append(votants, v...)
	}
	return votants
}

func CreateCandidatsByName(candidateNames []string) []model.Candidat {
	candidats := make([]model.Candidat, 0, len(candidateNames))
	max := len(candidateNames) - 1
	for i := max; i >= 0; i-- {
		candidats = append(candidats, model.NewCandidat(i, candidateNames[i]))
	}
	return candidats
}

func SimulateRandomVotes(votants []model.Votant, candidats []model.Candidat) {
	possibleVotes := generateAllPermutations(candidats)
	rand.Seed(time.Now().UnixNano())
	max := len(votants) - 1
	for i := max; i >= 0; i-- {
		votants[i].Vote = possibleVotes[rand.Intn(len(possibleVotes))]
	}
}

func generateAllPermutations(candidats []model.Candidat) [][]model.Candidat {
	result := [][]model.Candidat{}
	perm(candidats, &result)
	return result
}

func perm(a []model.Candidat, result *[][]model.Candidat) {
	generatePermutations(a, result, 0)
}

func generatePermutations(a []model.Candidat, result *[][]model.Candidat, i int) {
	if i > len(a) {
		comb := []model.Candidat{}
		comb = append(comb, a...)
		*result = append(*result, comb)
		return
	}
	generatePermutations(a, result, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		generatePermutations(a, result, i+1)
		a[i], a[j] = a[j], a[i]
	}
}
