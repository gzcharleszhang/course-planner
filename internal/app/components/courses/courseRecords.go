package courses

import (
	"time"
)

type CourseGrade float64
type CourseRecordId string

type CourseRecord struct {
	Course
	Id             CourseRecordId `json:"id"`
	Grade          CourseGrade    `json:"grade"`
	CompletionDate *time.Time     `json:"completion_date"`
}

type CourseRecords []*CourseRecord

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
