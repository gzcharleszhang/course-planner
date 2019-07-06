package Courses

import "app/components/Terms"

type TermCoursePlan struct {
	Term    Terms.Term
	Courses Courses
}

type CoursePlan []*TermCoursePlan
