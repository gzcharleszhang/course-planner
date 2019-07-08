package students

import (
	"github.com/gzcharleszhang/course-planner/internal/courses"
	"github.com/gzcharleszhang/course-planner/internal/terms"
)

type StudentName string
type StudentId int

type Student struct {
	Id            StudentId
	Name          StudentName
	CurrentTerm   terms.Term
	CourseRecords courses.CourseRecords
	CoursePlans   []*courses.CoursePlan
}

func NewStudent(name StudentName, startYear terms.TermYear) *Student {
	startTerm := (int(startYear)-1900)*10 + 9 // assuming student starts in fall
	return &Student{
		Id:            0, // TODO: generate unique student id
		Name:          name,
		CurrentTerm:   terms.NewTerm("1A", startTerm),
		CourseRecords: map[*courses.Course]courses.CourseRecord{},
		CoursePlans:   []*courses.CoursePlan{},
	}
}
