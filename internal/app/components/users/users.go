package users

import (
	"context"
	"github.com/gzcharleszhang/course-planner/internal/app/components/courses"
	"github.com/gzcharleszhang/course-planner/internal/app/components/timelines"
	"github.com/gzcharleszhang/course-planner/internal/app/db"
	"github.com/rs/xid"
	"golang.org/x/crypto/bcrypt"
)

type FirstName string
type LastName string
type UserId string
type PasswordHash string

type UserBase struct {
	Id        UserId       `bson:"_id"`
	FirstName FirstName    `bson:"first_name"`
	LastName  LastName     `bson:"last_name"`
	Password  PasswordHash `bson:"password"`
}

type UserData struct {
	UserBase      `bson:"user_base"`
	CourseHistory []courses.CourseRecordId `bson:"course_history"`
	Timelines     []timelines.TimelineId   `bson:"timelines"`
}

type User struct {
	UserBase      `bson:"user_base"`
	CourseHistory courses.CourseRecords `bson:"course_history"`
	Timelines     []*timelines.Timeline `bson:"timelines"`
}

func newUserId() UserId {
	return UserId(xid.New().String())
}

func CreateUser(ctx context.Context, firstName FirstName, lastName LastName, password PasswordHash) (UserId, error) {
	sess, err := db.NewSession(ctx)
	newUserId := newUserId()
	if err != nil {
		return "", err
	}
	user := UserData{
		UserBase: UserBase{
			Id:        newUserId,
			FirstName: firstName,
			LastName:  lastName,
			Password:  password,
		},
	}
	if _, err := sess.Users().InsertOne(ctx, user); err != nil {
		return "", err
	}
	return newUserId, nil
}

func HashPassword(password string) (PasswordHash, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return PasswordHash(bytes), err
}
