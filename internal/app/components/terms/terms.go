package terms

type TermName string
type TermSeason int
type TermYear int

const (
	WINTER TermSeason = 1
	SPRING TermSeason = 5
	FALL   TermSeason = 9
)

type Term struct {
	Name   TermName
	Season TermSeason
	Year   TermYear
}

func NewTerm(name TermName, uwTermId int) Term {
	year := TermYear(uwTermId/10 + 1900)
	season := TermSeason(uwTermId % 10)
	return Term{Name: name, Season: season, Year: year}
}
