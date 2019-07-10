package courseSelections

import (
	"github.com/gzcharleszhang/course-planner/internal/app/components/courses"
	"github.com/gzcharleszhang/course-planner/internal/app/components/terms"
)

type TermSelection struct {
	Term          terms.Term
	CourseRecords courses.CourseRecords
}

func NewTermSelection(termName terms.TermName, uwTermId int) *TermSelection {
	return &TermSelection{
		Term:          terms.NewTerm(termName, uwTermId),
		CourseRecords: courses.CourseRecords{},
	}
}

// return the courses whose pre-requisites are not satisfied
func (ts TermSelection) InvalidCourses(pastRecords courses.CourseRecords) courses.CourseRecords {
	invalidRecords := courses.CourseRecords{}
	for _, record := range ts.CourseRecords {
		if !isPrereqSatisfied(record, &pastRecords) {
			invalidRecords = append(invalidRecords, record)
		}
	}
	return invalidRecords
}

func isPrereqSatisfied(record *courses.CourseRecord, pastRecords *courses.CourseRecords) bool {
	return record.Prereqs.IsSatisfied(pastRecords)
}
