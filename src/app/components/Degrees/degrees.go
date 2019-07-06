package Degrees

import "app/components/Courses"

type DegreeName string
type DegreeRequirements Courses.CourseRequirementRules

type Degree interface {
	Plan(*Courses.CoursePlan) bool
	IsCompleted(*Courses.CoursePlan) bool
}

type RegularDegree struct {
	Name         DegreeName
	Requirements DegreeRequirements
}

func (deg RegularDegree) Plan(plan *Courses.CoursePlan) bool {
	return true
}

func (deg RegularDegree) IsCompleted(plan *Courses.CoursePlan) bool {
	return false
}
