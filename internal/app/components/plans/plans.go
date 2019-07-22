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
type DegreeRequirements courses.CourseRequirementRules

type Degree struct {
	Name         DegreeName
	Requirements DegreeRequirements
}

func (deg Degree) IsCompleted(records *courses.CourseRecords) bool {
	for _, req := range deg.Requirements {
		if !req.IsSatisfied(records) {
			return false
		}
	}
	return true
}

func (deg Degree) GetName() string {
	return string(deg.Name)
}
