package timelineModel

import (
	"context"
	"github.com/gzcharleszhang/course-planner/internal/app/components/plans"
	"github.com/gzcharleszhang/course-planner/internal/app/components/terms"
	"github.com/gzcharleszhang/course-planner/internal/app/components/timelines"
	"github.com/gzcharleszhang/course-planner/internal/app/components/users"
)

type TimelineModel struct {
	Id          timelines.TimelineId   `bson:"_id"`
	Name        timelines.TimelineName `bson:"name"`
	TermRecords []terms.TermRecordId   `bson:"term_records"`
	Plans       []plans.PlanId         `bson:"plans"`
	UserId      users.UserId           `bson:"user_id"`
}

func GetTimelinesByUserId(ctx context.Context, userId users.UserId) ([]*timelines.Timeline, error) {
	// TODO: implement
	return nil, nil
}
