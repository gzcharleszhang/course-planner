package courses

import (
	"time"
)

type CourseGrade float64

type CourseRecord struct {
	Course
	Grade          CourseGrade
	CompletionDate *time.Time
}

type CourseRecords map[CourseId]*CourseRecord

func (cr CourseRecords) Merge(records CourseRecords) CourseRecords {
	result := CourseRecords{}
	for id, record := range cr {
		result[id] = record
	}
	for id, record := range records {
		result[id] = record
	}
	return result
}

func (cr CourseRecords) Exclude(records CourseRecords) CourseRecords {
	result := CourseRecords{}
	for id, record := range cr {
		result[id] = record
	}
	for id := range records {
		delete(result, id)
	}
	return result
}
