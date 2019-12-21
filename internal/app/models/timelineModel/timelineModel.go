package timelineModel

import (
	"context"
	"github.com/gzcharleszhang/course-planner/internal/app/components/plans"
	"github.com/gzcharleszhang/course-planner/internal/app/components/terms"
	"github.com/gzcharleszhang/course-planner/internal/app/components/timelines"
	"github.com/gzcharleszhang/course-planner/internal/app/components/users"
	"github.com/gzcharleszhang/course-planner/internal/app/db"
	"github.com/gzcharleszhang/course-planner/internal/app/models/planModel"
	"github.com/gzcharleszhang/course-planner/internal/app/models/termRecordModel"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
)

type TimelineModel struct {
	Id            timelines.TimelineId   `bson:"_id"`
	Name          timelines.TimelineName `bson:"name"`
	TermRecordIds []terms.TermRecordId   `bson:"term_record_ids"`
	PlanIds       []plans.PlanId         `bson:"plan_ids"`
	UserId        users.UserId           `bson:"user_id"`
}

func (tlm TimelineModel) ToTimeline(ctx context.Context) (*timelines.Timeline, error) {
	trs, err := termRecordModel.GetTermRecordsByIds(ctx, tlm.TermRecordIds)
	if err != nil {
		return nil, err
	}
	pls, err := planModel.GetPlansByIds(ctx, tlm.PlanIds)
	if err != nil {
		return nil, err
	}
	tl := &timelines.Timeline{
		Id:          tlm.Id,
		Name:        tlm.Name,
		Plans:       pls,
		TermRecords: trs,
	}
	return tl, nil
}

func GetTimelinesByUserId(ctx context.Context, userId users.UserId) (timelines.Timelines, error) {
	sess := db.PrimarySession
	var tlms []TimelineModel
	err := sess.Timelines().FindOne(ctx, bson.M{"user_id": userId}).Decode(&tlms)
	if err != nil {
		return nil, err
	}
	var tls timelines.Timelines
	for _, tlm := range tlms {
		tl, err := tlm.ToTimeline(ctx)
		if err != nil {
			return nil, err
		}
		tls = append(tls, tl)
	}
	return tls, err
}

func GetTimelineById(ctx context.Context, timelineId timelines.TimelineId) (*timelines.Timeline, error) {
	sess := db.PrimarySession
	var res TimelineModel
	err := sess.Timelines().FindOne(ctx, bson.M{"_id": timelineId}).Decode(&res)
	if err != nil {
		return nil, err
	}
	tl, err := res.ToTimeline(ctx)
	return tl, err
}

func CreateTimeline(ctx context.Context, name timelines.TimelineName, userId users.UserId) (*timelines.Timeline, error) {
	tl := timelines.NewTimeline(name, nil)
	tlm := TimelineModel{
		Id:            tl.Id,
		Name:          name,
		TermRecordIds: []terms.TermRecordId{},
		PlanIds:       []plans.PlanId{},
		UserId:        userId,
	}
	sess := db.PrimarySession
	_, err := sess.Timelines().InsertOne(ctx, tlm)
	if err != nil {
		return nil, errors.Wrap(err, "Error creating new timeline: ")
	}
	return tl, nil
}
