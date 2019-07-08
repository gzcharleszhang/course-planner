package degrees

import "github.com/gzcharleszhang/course-planner/internal/courses"

type DegreeName string
type DegreeRequirements courses.CourseRequirementRules

type Degree interface {
	Plan(*courses.CoursePlan) bool
	IsCompleted(*courses.CoursePlan) bool
}

type RegularDegree struct {
	Name         DegreeName
	Requirements DegreeRequirements
}

func (deg RegularDegree) Plan(plan *courses.CoursePlan) bool {
	return true
}

func (deg RegularDegree) IsCompleted(plan *courses.CoursePlan) bool {
	return false
}
