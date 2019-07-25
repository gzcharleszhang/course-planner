package terms

import (
	"context"
	"github.com/gzcharleszhang/course-planner/internal/app/components/courses"
	"github.com/rs/xid"
)

type TermRecordId string

type TermRecord struct {
	Term          Term
	Id            TermRecordId
	CourseRecords courses.CourseRecords
}

type TermRecords []*TermRecord

func newTermRecordId() TermRecordId {
	return TermRecordId(xid.New().String())
}

func NewTermRecord(termName TermName, uwTermId int) *TermRecord {
	return &TermRecord{
		Term:          NewTerm(termName, uwTermId),
		Id:            newTermRecordId(),
		CourseRecords: courses.CourseRecords{},
	}
}

func (records TermRecords) Copy() TermRecords {
	var newTermRecords TermRecords
	for _, tr := range records {
		newRecord := tr.Copy()
		newTermRecords = append(newTermRecords, &newRecord)
	}
	return newTermRecords
}

func (tr TermRecord) Copy() TermRecord {
	termName, termId := tr.Term.Name, int(tr.Term.Id)
	newRecord := NewTermRecord(termName, termId)
	newCourseRecords := tr.CourseRecords.Copy()
	newRecord.CourseRecords = newCourseRecords
	return *newRecord
}

// Returns the courses whose pre-requisites are not satisfied
func (tr TermRecord) InvalidCourses(pastRecords courses.CourseRecords) courses.CourseRecords {
	invalidRecords := courses.CourseRecords{}
	for _, record := range tr.CourseRecords {
		if !record.IsPrereqSatisfied(&pastRecords) {
			invalidRecords = append(invalidRecords, record)
		}
	}
	return invalidRecords
}

// TODO: implement
func GetTermRecordById(ctx context.Context, id TermRecordId) (*TermRecord, error) {
	return nil, nil
}

// TODO: implement
func GetTermRecordsByIds(ctx context.Context, recordIds []TermRecordId) ([]*TermRecord, error) {
	return nil, nil
}
