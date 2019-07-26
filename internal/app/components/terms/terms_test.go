package terms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTerm(t *testing.T) {
	var termName TermName
	var term Term

	termName = TermName("1A")
	term = NewTerm(termName, 1179)
	assert.Equal(t, Term{Name: termName, Id: TermId(1179), Season: FALL, Year: TermYear(2017)}, term)

	termName = TermName("1B")
	term = NewTerm(termName, 1181)
	assert.Equal(t, Term{Name: termName, Id: TermId(1181), Season: WINTER, Year: TermYear(2018)}, term)

	termName = TermName("a long time ago")
	term = NewTerm(termName, 105)
	assert.Equal(t, Term{Name: termName, Id: TermId(105), Season: SPRING, Year: TermYear(1910)}, term)
}
