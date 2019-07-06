package Courses

import "app/components/Terms"

type TermCoursePlan struct {
	Term    Terms.Term
	Courses Courses
}

type CoursePlan struct {
	CompletedTermPlans []*TermCoursePlan
	FutureTermPlans    []*TermCoursePlan
}
