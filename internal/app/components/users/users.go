package users

import (
	"context"
	"github.com/gzcharleszhang/course-planner/internal/app/components/courses"
	"github.com/gzcharleszhang/course-planner/internal/app/components/permissions"
	"github.com/gzcharleszhang/course-planner/internal/app/components/timelines"
	"github.com/gzcharleszhang/course-planner/internal/app/db"
	"github.com/rs/xid"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

type FirstName string
type LastName string
type UserId string
type PasswordHash string

type UserData struct {
	Id               UserId                   `bson:"_id"`
	FirstName        FirstName                `bson:"first_name"`
	LastName         LastName                 `bson:"last_name"`
	Password         PasswordHash             `bson:"password"`
	CourseHistory    []courses.CourseRecordId `bson:"course_history"`
	PermissionAccess permissions.Permission   `bson:"permission_access"`
}

type User struct {
	Id               UserId                 `bson:"_id"`
	FirstName        FirstName              `bson:"first_name"`
	LastName         LastName               `bson:"last_name"`
	Password         PasswordHash           `bson:"password"`
	CourseHistory    courses.CourseRecords  `bson:"course_history"`
	Timelines        []*timelines.Timeline  `bson:"timelines"`
	PermissionAccess permissions.Permission `bson:"permission_access"`
}

func newUserId() UserId {
	return UserId(xid.New().String())
}

func CreateUser(ctx context.Context, firstName FirstName, lastName LastName,
	password PasswordHash, perm permissions.Permission) (UserId, error) {
	sess, err := db.NewSession(ctx)
	if err != nil {
		return "", err
	}
	defer sess.Close(ctx)
	newUserId := newUserId()
	user := UserData{
		Id:               newUserId,
		FirstName:        firstName,
		LastName:         lastName,
		Password:         password,
		PermissionAccess: perm,
	}
	if _, err := sess.Users().InsertOne(ctx, user); err != nil {
		return "", err
	}
	return newUserId, nil
}

func GetUserById(ctx context.Context, userId UserId) (*User, error) {
	sess, err := db.NewSession(ctx)
	if err != nil {
		return nil, err
	}
	var result UserData
	err = sess.Users().FindOne(ctx, bson.M{"_id": userId}).Decode(&result)
	if err != nil {
		return nil, err
	}
	history, err := courses.GetCourseRecordsByIds(ctx, result.CourseHistory)
	if err != nil {
		return nil, err
	}
	tls, err := timelines.GetTimelinesByUserId(ctx, result.Id)
	if err != nil {
		return nil, err
	}
	user := User{
		Id:               result.Id,
		FirstName:        result.FirstName,
		LastName:         result.LastName,
		Password:         result.Password,
		CourseHistory:    history,
		Timelines:        tls,
		PermissionAccess: result.PermissionAccess,
	}
	return &user, nil
}

func HashPassword(password string) (PasswordHash, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return PasswordHash(bytes), err
}

func GetUserPermissionAccess(ctx context.Context, id UserId) (*permissions.Permission, error) {
	user, err := GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}
	return &user.PermissionAccess, nil
}
