package plans

import (
	"github.com/gzcharleszhang/course-planner/internal/app/components/courses"
)

type Plan interface {
	IsCompleted(courseSelection *courses.CourseSelection) bool
}

type DegreeName string
type DegreeRequirements courses.CourseRequirementRule

type Degree struct {
	Name         DegreeName
	Requirements DegreeRequirements
}

func (deg Degree) IsCompleted(courseSelection *courses.CourseSelection) bool {
	records := courseSelection.Aggregate()
	return deg.Requirements.IsSatisfied(records)
}
