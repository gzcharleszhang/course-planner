package courses

import "github.com/gzcharleszhang/course-planner/internal/app/components/terms"

type TermSelection struct {
	Term          terms.Term
	CourseRecords CourseRecords
}

func NewTermSelection(termName terms.TermName, uwTermId int) *TermSelection {
	return &TermSelection{
		Term:          terms.NewTerm(termName, uwTermId),
		CourseRecords: CourseRecords{},
	}
}
