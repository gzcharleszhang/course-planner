package getTimelinesService

import (
	"context"
	"github.com/gzcharleszhang/course-planner/internal/app/components/timelines"
	"github.com/gzcharleszhang/course-planner/internal/app/components/users"
	"github.com/gzcharleszhang/course-planner/internal/app/models/timelineModel"
)

type Request struct {
	UserId users.UserId `json:"user_id"`
}

type Response struct {
	Timelines timelines.Timelines `json:"timelines"`
}

func Execute(ctx context.Context, req Request) (*Response, error) {
	tls, err := timelineModel.GetTimelinesByUserId(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	res := Response{Timelines: tls}
	return &res, nil
}
