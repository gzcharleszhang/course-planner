package courses

import "github.com/gzcharleszhang/course-planner/internal/terms"

type TermCoursePlan struct {
	Term    terms.Term
	Courses Courses
}

type CoursePlan struct {
	CompletedTermPlans []*TermCoursePlan
	FutureTermPlans    []*TermCoursePlan
}
