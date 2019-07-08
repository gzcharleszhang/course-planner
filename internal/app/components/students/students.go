package students

import (
	"github.com/gzcharleszhang/course-planner/internal/app/components/courses"
	"github.com/rs/xid"
)

type StudentFirstName string
type StudentLastName string
type StudentId string

type Student struct {
	Id            StudentId
	FirstName     StudentFirstName
	LastName      StudentLastName
	CourseRecords courses.CourseRecords
	CoursePlans   []*courses.CoursePlan
}

func newStudentId() StudentId {
	return StudentId(xid.New().String())
}

func NewStudent(firstName StudentFirstName, lastName StudentLastName) *Student {
	return &Student{
		Id:            newStudentId(),
		FirstName:     firstName,
		LastName:      lastName,
		CourseRecords: courses.CourseRecords{},
		CoursePlans:   []*courses.CoursePlan{},
	}
}
