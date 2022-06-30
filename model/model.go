package model

type Candidat struct {
	Id   int
	Name string
}

type Votant struct {
	Id   int
	Name string
	//Votes bool
	Vote ScrutinIndiv
}

type Diff struct {
	Ldiff int
	Rdiff int
}

type Pair struct {
	L Candidat
	R Candidat
}

type ResultPerCandidate struct {
	Cand Candidat
	Res  Result
}

type ResultList []ResultPerCandidate

type Result struct {
	DuelWon    int
	TotalPoint int
	WonAgainst []Candidat
}

type ScrutinIndiv []Candidat

func (s ScrutinIndiv) IndexOf(c Candidat) int {
	for pos, v := range s {
		if v == c {
			return pos
		}
	}
	return -1
}

type ResultDto struct {
	D Diff
	L Candidat
	R Candidat
}

func NewResultDto(d Diff, l Candidat, r Candidat) ResultDto {
	return ResultDto{
		D: d,
		L: l,
		R: r,
	}
}

func NewCandidat(id int, name string) Candidat {
	return Candidat{
		Id:   id,
		Name: name,
	}
}

func NewVotant(id int, name string) Votant {
	return Votant{
		Id:   id,
		Name: name,
	}
}
