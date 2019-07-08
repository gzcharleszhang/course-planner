package courses

import "time"

type CourseGrade float64

type CourseRecord struct {
	Grade          CourseGrade
	CompletionDate *time.Time
}

type CourseRecords map[*Course]CourseRecord
