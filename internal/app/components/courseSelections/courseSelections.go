package courseSelections

import (
	"github.com/gzcharleszhang/course-planner/internal/app/components/courses"
	"github.com/gzcharleszhang/course-planner/internal/app/components/plans"
	"github.com/rs/xid"
)

type CourseSelectionId string
type CourseSelectionName string

type CourseSelection struct {
	Id             CourseSelectionId
	Name           CourseSelectionName
	TermSelections []*courses.TermSelection
	Plans          []plans.Plan
}

func newCourseSelectionId() CourseSelectionId {
	return CourseSelectionId(xid.New().String())
}

func NewCourseSelection(name CourseSelectionName) *CourseSelection {
	return &CourseSelection{
		Id:             newCourseSelectionId(),
		Name:           name,
		TermSelections: []*courses.TermSelection{},
		Plans:          []plans.Plan{},
	}
}

func (cs CourseSelection) AddTermSelection(ts *courses.TermSelection) {
	cs.TermSelections = append(cs.TermSelections, ts)
}

// Flattens TermSelections and aggregates it into one CourseRecords
func (cs CourseSelection) Aggregate() *courses.CourseRecords {
	records := courses.CourseRecords{}
	for _, ts := range cs.TermSelections {
		for id, record := range ts.CourseRecords {
			records[id] = record
		}
	}
	return &records
}

// TODO: make it return the plans that are not satisfied
// Checks if all declared plans in Plans are satisfied
func (cs CourseSelection) IsSatisfied() bool {
	records := cs.Aggregate()
	for _, plan := range cs.Plans {
		if !plan.IsCompleted(records) {
			return false
		}
	}
	return true
}
