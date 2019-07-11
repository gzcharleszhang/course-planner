package users

import (
	"github.com/gzcharleszhang/course-planner/internal/app/components/courses"
	"github.com/gzcharleszhang/course-planner/internal/app/components/timelines"
	"github.com/rs/xid"
)

type FirstName string
type LastName string
type UserId string

type User struct {
	Id            UserId
	FirstName     FirstName
	LastName      LastName
	CourseHistory courses.CourseRecords
	Timelines     []*timelines.Timeline
}

func newUserId() UserId {
	return UserId(xid.New().String())
}

func NewUser(firstName FirstName, lastName LastName) *User {
	return &User{
		Id:            newUserId(),
		FirstName:     firstName,
		LastName:      lastName,
		CourseHistory: courses.CourseRecords{},
		Timelines:     []*timelines.Timeline{},
	}
}
