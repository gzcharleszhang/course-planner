package newTimelineService

import (
	"context"
	"github.com/gzcharleszhang/course-planner/internal/app/components/timelines"
	"github.com/gzcharleszhang/course-planner/internal/app/components/users"
	"github.com/gzcharleszhang/course-planner/internal/app/models/timelineModel"
)

type Request struct {
	UserId users.UserId           `json:"user_id"`
	Name   timelines.TimelineName `json:"name"`
}

type Response struct {
	Timeline timelines.Timeline `json:"timeline"`
}

func Execute(ctx context.Context, req Request) (*Response, error) {
	tl, err := timelineModel.NewTimeline(ctx, req.Name, req.UserId)
	if err != nil {
		return nil, err
	}
	res := Response{Timeline: *tl}
	return &res, nil
}
