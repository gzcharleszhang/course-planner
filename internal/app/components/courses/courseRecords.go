package courses

import (
	"context"
	"github.com/rs/xid"
	"time"
)

type CourseGrade float64
type CourseRecordId string

type CourseRecord struct {
	Course
	Id             CourseRecordId `json:"id"`
	Grade          CourseGrade    `json:"grade"`
	CompletionDate *time.Time     `json:"completion_date"`
	Override       bool           `json:"override"` // user specified this course was overridden so no need to check pre-requisites
}

type CourseRecords []*CourseRecord

func newCourseRecordId() CourseRecordId {
	return CourseRecordId(xid.New().String())
}

func (records CourseRecords) Copy() CourseRecords {
	var newCourseRecords CourseRecords
	for _, cr := range records {
		newRecord := cr.Copy()
		newCourseRecords = append(newCourseRecords, &newRecord)
	}
	return newCourseRecords
}

func (cr CourseRecord) Copy() CourseRecord {
	newCompletionDate := cr.CompletionDate
	if cr.CompletionDate != nil {
		copyTime := *cr.CompletionDate
		newCompletionDate = &copyTime
	}
	id := newCourseRecordId()
	newRecord := CourseRecord{cr.Course, id, cr.Grade, newCompletionDate, cr.Override}
	return newRecord
}


func GetCourseRecordById(ctx context.Context, recordId CourseRecordId) (*CourseRecord, error) {
	// TODO: implement
	return nil, nil
}

func GetCourseRecordsByIds(ctx context.Context, recordIds []CourseRecordId) (CourseRecords, error) {
	// TODO: implement
	return nil, nil
}

// convert course record to a map with course id as key, using the higher grade as tie breaker
func (cr CourseRecords) ToCourseIdMap() map[CourseId]*CourseRecord {
	idMap := map[CourseId]*CourseRecord{}
	for _, record := range cr {
		oldRecord, exists := idMap[record.Course.Id]
		// always take a future course instead of a completed one
		if !exists || record.CompletionDate == nil ||
			(oldRecord.CompletionDate != nil && oldRecord.Grade < record.Grade) {
			idMap[record.Course.Id] = record
		}
	}
	return idMap
}

// merge two course records into one slice
func (cr CourseRecords) Merge(records CourseRecords) CourseRecords {
	result := CourseRecords{}
	for _, record := range cr {
		result = append(result, record)
	}
	for _, record := range records {
		result = append(result, record)
	}
	return result
}

// exclude any course record in records from cr by course id
func (cr CourseRecords) Exclude(records CourseRecords) CourseRecords {
	result := CourseRecords{}
	excludeMap := records.ToCourseIdMap()
	for _, record := range cr {
		// only append record to result if not in the exclude map
		_, inExcludeMap := excludeMap[record.Course.Id]
		if !inExcludeMap {
			result = append(result, record)
		}
	}
	return result
}

// calculates the cumulative average of completed courses
func (cr CourseRecords) CurrentCAV() CourseGrade {
	total, count := CourseGrade(0), 0
	for _, record := range cr {
		// only accumulate if course is completed
		if record.IsCompleted() {
			total += record.Grade
			count += 1
		}
	}
	if count == 0 {
		return CourseGrade(0)
	}
	return CourseGrade(int(total) / count)
}

func (cr CourseRecord) IsPrereqSatisfied(pastRecords *CourseRecords) bool {
	// if it's overridden or course has no pre-reqs, then it's satisfied
	prereqs := cr.Prereqs
	if cr.Override || prereqs == nil {
		return true
	}
	for _, prereq := range prereqs {
		if !prereq.IsSatisfied(pastRecords) {
			return false
		}
	}
	return true
}

// whether or not the user has completed this course
func (cr CourseRecord) IsCompleted() bool {
	return cr.CompletionDate != nil
}
