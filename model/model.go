package model

type Candidate struct {
	Id   int
	Name string
}

type Voter struct {
	Id   int
	Name string
	Vote ScrutinIndiv
}

type Diff struct {
	Ldiff int
	Rdiff int
}

type Pair struct {
	L Candidate
	R Candidate
}

type ResultPerCandidate struct {
	Cand Candidate
	Res  Result
}

type ResultList []ResultPerCandidate

type Result struct {
	DuelsWon    int
	TotalPoints int
	WonAgainst  []Candidate
}

type ScrutinIndiv []Candidate

type ResultDto struct {
	D Diff
	L Candidate
	R Candidate
}

func (s ScrutinIndiv) IndexOf(c Candidate) int {
	for pos, v := range s {
		if v == c {
			return pos
		}
	}
	return -1
}

func NewResultDto(d Diff, l Candidate, r Candidate) ResultDto {
	return ResultDto{
		D: d,
		L: l,
		R: r,
	}
}

func NewCandidate(id int, name string) Candidate {
	return Candidate{
		Id:   id,
		Name: name,
	}
}

func NewVoter(id int, name string) Voter {
	return Voter{
		Id:   id,
		Name: name,
	}
}
