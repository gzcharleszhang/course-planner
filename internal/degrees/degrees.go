package degrees

import (
	"github.com/gzcharleszhang/course-planner/internal/courses"
	"github.com/gzcharleszhang/course-planner/internal/students"
)

type DegreeName string
type DegreeRequirements courses.CourseRequirementRules

type Degree interface {
	Plan(*courses.CoursePlan) bool
	IsCompleted(*students.Student, *courses.CoursePlan) bool
}

type RegularDegree struct {
	Name         DegreeName
	Requirements DegreeRequirements
}

func (deg RegularDegree) Plan(plan *courses.CoursePlan) bool {
	// TODO: implement
	return true
}

func (deg RegularDegree) IsCompleted(student *students.Student, plan *courses.CoursePlan) bool {
	for _, courseRequirement := range deg.Requirements {
		if !courseRequirement.IsSatisfied(&student.CompletedCourses) {
			return false
		}
	}
	return true
}
