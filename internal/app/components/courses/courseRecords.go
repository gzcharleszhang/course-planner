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
