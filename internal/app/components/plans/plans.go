package plans

import (
	"github.com/gzcharleszhang/course-planner/internal/app/components/courses"
)

type Plan interface {
	IsCompleted(records *courses.CourseRecords) bool
}

type DegreeName string
type DegreeRequirements courses.CourseRequirementRule

type Degree struct {
	Name         DegreeName
	Requirements DegreeRequirements
}

func (deg Degree) IsCompleted(records *courses.CourseRecords) bool {
	return deg.Requirements.IsSatisfied(records)
}
