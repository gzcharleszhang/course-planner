package termRecordModel

import (
	"context"
	"github.com/gzcharleszhang/course-planner/internal/app/components/courses"
	"github.com/gzcharleszhang/course-planner/internal/app/components/terms"
)

type TermRecordModel struct {
	Id              terms.TermRecordId       `bson:"_id"`
	Term            terms.Term               `bson:"term"`
	CourseRecordIds []courses.CourseRecordId `bson:"course_record_ids"`
}

func GetTermRecordById(ctx context.Context, id terms.TermRecordId) (*terms.TermRecord, error) {
	// TODO: implement
	return nil, nil
}

func GetTermRecordsByIds(ctx context.Context, ids []terms.TermRecordId) (terms.TermRecords, error) {
	var records terms.TermRecords
	for _, id := range ids {
		record, err := GetTermRecordById(ctx, id)
		if err != nil {
			return nil, err
		}
		records = append(records, record)
	}
	return records, nil
}
