package terms

import (
	"github.com/gzcharleszhang/course-planner/internal/app/components/courses"
)

type TermRecord struct {
	Term          Term
	CourseRecords courses.CourseRecords
}

func NewTermRecord(termName TermName, uwTermId int) *TermRecord {
	return &TermRecord{
		Term:          NewTerm(termName, uwTermId),
		CourseRecords: courses.CourseRecords{},
	}
}

// Returns the courses whose pre-requisites are not satisfied
func (tr TermRecord) InvalidCourses(pastRecords courses.CourseRecords) courses.CourseRecords {
	invalidRecords := courses.CourseRecords{}
	for _, record := range tr.CourseRecords {
		if !record.IsPrereqSatisfied(&pastRecords) {
			invalidRecords = append(invalidRecords, record)
		}
	}
	return invalidRecords
}
