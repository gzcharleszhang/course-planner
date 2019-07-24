package terms

type TermName string
type TermSeason int
type TermYear int
type TermId int

const (
	WINTER TermSeason = 1
	SPRING TermSeason = 5
	FALL   TermSeason = 9
)

type Term struct {
	Name   TermName
	Id 	   TermId
	Season TermSeason
	Year   TermYear
}

func NewTerm(name TermName, uwTermId int) Term {
	year := TermYear(uwTermId/10 + 1900)
	season := TermSeason(uwTermId % 10)
	id := TermId(uwTermId)
	return Term{Name: name, Id: id, Season: season, Year: year}
}
