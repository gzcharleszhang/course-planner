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
	TermSelections []*TermSelection
	Plans          plans.Plans
}

func newCourseSelectionId() CourseSelectionId {
	return CourseSelectionId(xid.New().String())
}

func NewCourseSelection(name CourseSelectionName) *CourseSelection {
	return &CourseSelection{
		Id:             newCourseSelectionId(),
		Name:           name,
		TermSelections: []*TermSelection{},
		Plans:          plans.Plans{},
	}
}

func (cs CourseSelection) AddTermSelection(ts *TermSelection) {
	cs.TermSelections = append(cs.TermSelections, ts)
}

// Flattens TermSelections and aggregates it into one CourseRecords
func (cs CourseSelection) Aggregate() *courses.CourseRecords {
	records := courses.CourseRecords{}
	for _, ts := range cs.TermSelections {
		for _, record := range ts.CourseRecords {
			records = append(records, record)
		}
	}
	return &records
}

// Checks if all declared plans in Plans are satisfied
func (cs CourseSelection) IncompletePlans() plans.Plans {
	records := cs.Aggregate()
	var incompletePlans plans.Plans
	for _, plan := range cs.Plans {
		if !plan.IsCompleted(records) {
			incompletePlans = append(incompletePlans, plan)
		}
	}
	return incompletePlans
}

func (cs CourseSelection) InvalidCourses() courses.CourseRecords {
	pastRecords, invalidRecords := courses.CourseRecords{}, courses.CourseRecords{}
	for _, ts := range cs.TermSelections {
		// we keep accumulating invalid courses for each term
		invalidCourses := ts.InvalidCourses(pastRecords)
		invalidRecords = invalidRecords.Merge(invalidCourses)
		// we keep accumulating past records that are valid
		pastRecords = pastRecords.Merge(ts.CourseRecords.Exclude(invalidRecords))
	}
	return invalidRecords
}
