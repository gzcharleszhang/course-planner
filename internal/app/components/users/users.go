package users

import (
	"context"
	"github.com/gzcharleszhang/course-planner/internal/app/components/courses"
	"github.com/gzcharleszhang/course-planner/internal/app/components/timelines"
	"github.com/rs/xid"
	"golang.org/x/crypto/bcrypt"
)

type FirstName string
type LastName string
type UserId string
type PasswordHash string

type User struct {
	Id            UserId
	FirstName     FirstName
	LastName      LastName
	CourseHistory courses.CourseRecords
	Timelines     []*timelines.Timeline
	Password      PasswordHash
}

func newUserId() UserId {
	return UserId(xid.New().String())
}

func CreateUser(ctx context.Context, firstName FirstName, lastName LastName, password PasswordHash) error {
	user := User{
		Id:            newUserId(),
		FirstName:     firstName,
		LastName:      lastName,
		Password:      password,
		CourseHistory: courses.CourseRecords{},
		Timelines:     []*timelines.Timeline{},
	}
}

func HashPassword(password string) (PasswordHash, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return PasswordHash(bytes), err
}
