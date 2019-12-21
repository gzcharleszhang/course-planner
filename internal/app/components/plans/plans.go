package plans

import (
	"github.com/gzcharleszhang/course-planner/internal/app/components/courses"
)

type Plan interface {
	IsCompleted(records *courses.CourseRecords) bool
	GetName() string
}
type Plans []*Plan
type PlanId string
type PlanType string
type PlanName string
type PlanRequirements courses.CourseRequirementRules

const PlanTypeDegree PlanType = "degree"

// TODO: if no minors or options differ from this structure, we can change Plan to this and remove the interface
type Degree struct {
	Id           PlanId           `json:"_id"`
	PlanType     PlanType         `json:"plan_type"`
	Name         PlanName         `json:"name"`
	Requirements PlanRequirements `json:"requirements"`
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
