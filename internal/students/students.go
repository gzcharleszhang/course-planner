package students

import (
	"github.com/gzcharleszhang/course-planner/internal/courses"
	"github.com/gzcharleszhang/course-planner/internal/terms"
)

type StudentName string
type StudentId int

type Student struct {
	Name             StudentName
	Id               StudentId
	CurrentTerm      terms.Term
	CoursePlans      []*courses.CoursePlan
	CompletedCourses courses.CompletedCourses
}

func NewStudent(name StudentName, startYear terms.TermYear) *Student {
	startTerm := (int(startYear)-1900)*10 + 9 // assuming student starts in fall
	return &Student{
		Name:             name,
		Id:               0, // TODO: generate unique student id
		CurrentTerm:      terms.NewTerm("1A", startTerm),
		CoursePlans:      []*courses.CoursePlan{},
		CompletedCourses: map[courses.CourseId]courses.CompletedCourse{},
	}
}
