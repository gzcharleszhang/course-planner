package courses

import (
	"github.com/gzcharleszhang/course-planner/internal/terms"
)

type TermPlan struct {
	Term          terms.Term
	CourseRecords CourseRecords
}

type CoursePlan struct {
	TermPlans []*TermPlan
}
