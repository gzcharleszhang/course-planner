package getTimelineService

import (
	"context"
	"github.com/gzcharleszhang/course-planner/internal/app/components/timelines"
	"github.com/gzcharleszhang/course-planner/internal/app/models/timelineModel"
)

type Request struct {
	TimelineId timelines.TimelineId `json:"timeline_id"`
}

type Response struct {
	Timeline timelines.Timeline `json:"timeline"`
}

func Execute(ctx context.Context, req Request) (*Response, error) {
	tl, err := timelineModel.GetTimelineById(ctx, req.TimelineId)
	if err != nil {
		return nil, err
	}
	res := Response{Timeline: *tl}
	return &res, nil
}
