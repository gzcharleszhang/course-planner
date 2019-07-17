package terms

import (
	"github.com/gzcharleszhang/course-planner/internal/app/components/courses"
	"github.com/rs/xid"
)

type TermRecordId string

type TermRecord struct {
	Term          Term
	Id            TermRecordId
	CourseRecords courses.CourseRecords
}

func newTermRecordId() TermRecordId {
	return TermRecordId(xid.New().String())
}

func NewTermRecord(termName TermName, uwTermId int) *TermRecord {
	return &TermRecord{
		Term:          NewTerm(termName, uwTermId),
		Id:						 newTermRecordId(),
		CourseRecords: courses.CourseRecords{},
	}
}

// Returns the courses whose pre-requisites are not satisfied
func (tr TermRecord) InvalidCourses(pastRecords courses.CourseRecords) courses.CourseRecords {
	invalidRecords := courses.CourseRecords{}
	for _, record := range tr.CourseRecords {
		if !isPrereqSatisfied(record, &pastRecords) {
			invalidRecords = append(invalidRecords, record)
		}
	}
	return invalidRecords
}

func isPrereqSatisfied(record *courses.CourseRecord, pastRecords *courses.CourseRecords) bool {
	// if no pre-reqs, then it's satisfied
	prereqs := record.Prereqs
	if prereqs == nil {
		return true
	}
	return record.Prereqs.IsSatisfied(pastRecords)
}

// TODO: implement
func GetTermRecordById(id TermRecordId) TermRecord {
	return nil
}
