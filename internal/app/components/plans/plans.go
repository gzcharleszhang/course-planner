package plans

import (
	"github.com/gzcharleszhang/course-planner/internal/app/components/courses"
)

type Plan interface {
	IsCompleted(coursePlan *courses.CoursePlan) bool
}

type DegreeName string
type DegreeRequirements courses.CourseRequirementRule

type Degree struct {
	Name         DegreeName
	Requirements DegreeRequirements
}

func (deg Degree) IsCompleted(coursePlan *courses.CoursePlan) bool {
	records := coursePlan.Aggregate()
	return deg.Requirements.IsSatisfied(records)
}
