package planModel

import (
	"context"
	"github.com/gzcharleszhang/course-planner/internal/app/components/plans"
)

type PlanModel struct {
	Id   plans.PlanId   `bson:"_id"`
	Name plans.PlanName `bson:"name"`
	Type plans.PlanType `bson:"type"`
}

func GetPlanById(ctx context.Context, id plans.PlanId) (*plans.Plan, error) {
	// TODO: implement
	return nil, nil
}

func GetPlansByIds(ctx context.Context, ids []plans.PlanId) (plans.Plans, error) {
	var res plans.Plans
	for _, id := range ids {
		plan, err := GetPlanById(ctx, id)
		if err != nil {
			return nil, err
		}
		res = append(res, plan)
	}
	return res, nil
}
