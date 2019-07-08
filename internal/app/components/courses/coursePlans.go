package courses

import (
	"github.com/gzcharleszhang/course-planner/internal/app/components/terms"
)

type TermSelection struct {
	Term          terms.Term
	CourseRecords CourseRecords
}

type CourseSelectionId string
type CourseSelectionName string

type CourseSelection struct {
	Id             CourseSelectionId
	Name           CourseSelectionName
	TermSelections []*TermSelection
}

// Flattens TermPlans and aggregates it into one CourseRecords
func (cs CourseSelection) Aggregate() *CourseRecords {
	records := CourseRecords{}
	for _, ts := range cs.TermSelections {
		for id, record := range ts.CourseRecords {
			records[id] = record
		}
	}
	return &records
}
