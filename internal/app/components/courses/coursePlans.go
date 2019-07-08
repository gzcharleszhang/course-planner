package courses

import (
	"github.com/gzcharleszhang/course-planner/internal/app/components/terms"
)

type TermPlan struct {
	Term          terms.Term
	CourseRecords CourseRecords
}

type CoursePlanId string
type CoursePlanName string

type CoursePlan struct {
	Id        CoursePlanId
	Name      CoursePlanName
	TermPlans []*TermPlan
}

func (cp CoursePlan) Aggregate() *CourseRecords {
	records := CourseRecords{}
	for _, tp := range cp.TermPlans {
		for id, record := range tp.CourseRecords {
			records[id] = record
		}
	}
	return &records
}
