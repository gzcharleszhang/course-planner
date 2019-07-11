package plans

import (
	"github.com/gzcharleszhang/course-planner/internal/app/components/courses"
)

type Plan interface {
	IsCompleted(records *courses.CourseRecords) bool
	GetName() string
}
type Plans []Plan

type DegreeName string
type DegreeRequirements courses.CourseRequirementRule

type Degree struct {
	Name         DegreeName
	Requirements DegreeRequirements
}

func (deg Degree) IsCompleted(records *courses.CourseRecords) bool {
	return deg.Requirements.IsSatisfied(records)
}

func (deg Degree) GetName() string {
	return string(deg.Name)
}
