package timelines

import (
	"context"
	"github.com/gzcharleszhang/course-planner/internal/app/components/courses"
	"github.com/gzcharleszhang/course-planner/internal/app/components/plans"
	"github.com/gzcharleszhang/course-planner/internal/app/components/terms"
	"github.com/gzcharleszhang/course-planner/internal/app/components/users"
	"github.com/rs/xid"
)

type TimelineId string
type TimelineName string

type Timeline struct {
	Id          TimelineId
	Name        TimelineName
	TermRecords []*terms.TermRecord
	Plans       plans.Plans
}

func newTimelineId() TimelineId {
	return TimelineId(xid.New().String())
}

// TODO: use cr parameter when creating Timeline
func NewTimeline(name TimelineName, cr *courses.CourseRecords) *Timeline {
	return &Timeline{
		Id:          newTimelineId(),
		Name:        name,
		TermRecords: []*terms.TermRecord{},
		Plans:       plans.Plans{},
	}
}

func GetTimelineById(ctx context.Context, timelineId TimelineId) (*Timeline, error) {
	// TODO: implement
	return nil, nil
}

func GetTimelinesByUserId(ctx context.Context, userId users.UserId) ([]*Timeline, error) {
	// TODO: implement
	return nil, nil
}

func (t Timeline) AddTermRecord(tr *terms.TermRecord) {
	t.TermRecords = append(t.TermRecords, tr)
}

// Flattens TermRecords and aggregates it into one CourseRecords
func (t Timeline) Aggregate() *courses.CourseRecords {
	records := courses.CourseRecords{}
	for _, tr := range t.TermRecords {
		for _, record := range tr.CourseRecords {
			records = append(records, record)
		}
	}
	return &records
}

// Checks if all declared plans in Plans are satisfied
func (t Timeline) IncompletePlans() plans.Plans {
	records := t.Aggregate()
	var incompletePlans plans.Plans
	for _, plan := range t.Plans {
		if !plan.IsCompleted(records) {
			incompletePlans = append(incompletePlans, plan)
		}
	}
	return incompletePlans
}

// Returns the courses whose pre-requisites are not satisfied
func (t Timeline) InvalidCourses() courses.CourseRecords {
	pastRecords, invalidRecords := courses.CourseRecords{}, courses.CourseRecords{}
	for _, tr := range t.TermRecords {
		// we keep accumulating invalid courses for each term
		invalidCourses := tr.InvalidCourses(pastRecords)
		invalidRecords = invalidRecords.Merge(invalidCourses)
		// we keep accumulating past records that are valid
		pastRecords = pastRecords.Merge(tr.CourseRecords.Exclude(invalidRecords))
	}
	return invalidRecords
}
