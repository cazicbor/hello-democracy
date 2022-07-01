package mocks

import (
	"math/rand"
	"strconv"
	"sync"
	"time"

	"github.com/cazicbor/hello-democracy/config"

	"github.com/cazicbor/hello-democracy/model"
)

func GenerateRandomVoters(number int) []model.Voter {
	voters := make([]model.Voter, 0, number)
	for i := 1; i <= number; i++ {
		voters = append(voters, model.NewVoter(i, "Votant-"+strconv.Itoa(i)))
	}
	return voters
}

func AsyncGenerateRandomVoters(total int) []model.Voter {
	voters := make([]model.Voter, 0, total)
	var wg sync.WaitGroup
	ch := make(chan []model.Voter, (total/config.SubSliceSize)+1)
	for j, k := 0, config.SubSliceSize; j+config.SubSliceSize < total+config.SubSliceSize; j += config.SubSliceSize {
		wg.Add(1)
		go func(j, k int, ch chan []model.Voter) {
			defer wg.Done()
			diff := total - j
			currentSubSliceSize := config.SubSliceSize
			if diff < config.SubSliceSize {
				currentSubSliceSize = diff
			}

			subSlice := make([]model.Voter, 0, currentSubSliceSize)
			for i, l := j, 0; l < currentSubSliceSize; i++ {
				subSlice = append(subSlice, model.NewVoter(i, "Voter-"+strconv.Itoa(i)))
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
		voters = append(voters, v...)
	}
	return voters
}

func CreateCandidatesByName(candidatesNames []string) []model.Candidate {
	candidates := make([]model.Candidate, 0, len(candidatesNames))
	max := len(candidatesNames) - 1
	for i := max; i >= 0; i-- {
		candidates = append(candidates, model.NewCandidate(i, candidatesNames[i]))
	}
	return candidates
}

func SimulateRandomVotes(voters []model.Voter, candidates []model.Candidate) {
	possibleVotes := generateAllPermutations(candidates)
	rand.Seed(time.Now().UnixNano())
	max := len(voters) - 1
	for i := max; i >= 0; i-- {
		voters[i].Vote = possibleVotes[rand.Intn(len(possibleVotes))]
	}
}

func generateAllPermutations(candidates []model.Candidate) [][]model.Candidate {
	result := [][]model.Candidate{}
	perm(candidates, &result)
	return result
}

func perm(candidates []model.Candidate, result *[][]model.Candidate) {
	generatePermutations(candidates, result, 0)
}

func generatePermutations(candidates []model.Candidate, result *[][]model.Candidate, i int) {
	if i > len(candidates) {
		comb := []model.Candidate{}
		comb = append(comb, candidates...)
		*result = append(*result, comb)
		return
	}
	generatePermutations(candidates, result, i+1)
	for j := i + 1; j < len(candidates); j++ {
		candidates[i], candidates[j] = candidates[j], candidates[i]
		generatePermutations(candidates, result, i+1)
		candidates[i], candidates[j] = candidates[j], candidates[i]
	}
}
