package terms

type TermId string
type TermName string
type TermYear int

type Term struct {
	Id   TermId
	Name TermName
	Year TermYear
}
