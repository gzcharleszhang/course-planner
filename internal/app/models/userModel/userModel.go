package userModel

import (
	"context"
	"errors"
	"github.com/gzcharleszhang/course-planner/internal/app/components/roles"
	"github.com/gzcharleszhang/course-planner/internal/app/components/terms"
	"github.com/gzcharleszhang/course-planner/internal/app/components/timelines"
	"github.com/gzcharleszhang/course-planner/internal/app/components/users"
	"github.com/gzcharleszhang/course-planner/internal/app/db"
	"github.com/gzcharleszhang/course-planner/internal/app/models/timelineModel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserModel struct {
	Id            users.UserId           `bson:"_id"`
	FirstName     users.FirstName        `bson:"first_name"`
	LastName      users.LastName         `bson:"last_name"`
	Email         users.Email            `bson:"email"`
	Password      users.PasswordHash     `bson:"password"`
	CourseHistory []terms.TermRecordId   `bson:"course_history"`
	Timelines     []timelines.TimelineId `bson:"timelines"`
	RoleId        roles.RoleId           `bson:"role_id"`
}

func CreateUser(ctx context.Context, firstName users.FirstName, lastName users.LastName,
	email users.Email, password users.PasswordHash) (users.UserId, error) {
	sess, err := db.NewSession(ctx)
	if err != nil {
		return "", err
	}
	defer sess.Close(ctx)
	// check for duplicate emails
	userExists, err := checkDuplicateEmail(ctx, email)
	if err != nil {
		return "", err
	}
	if userExists {
		return "", errors.New("email already exists")
	}
	newUserId := users.NewUserId()
	user := UserModel{
		Id:        newUserId,
		FirstName: firstName,
		LastName:  lastName,
		Password:  password,
		Email:     email,
		RoleId:    roles.ConradId, // default to conrad
	}
	if _, err := sess.Users().InsertOne(ctx, user); err != nil {
		return "", err
	}
	return newUserId, nil
}

// return true if there already exists an user with the given email
func checkDuplicateEmail(ctx context.Context, email users.Email) (bool, error) {
	_, err := GetUserByEmail(ctx, email)
	if err != nil {
		// no result found means no duplicates
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func GetUserById(ctx context.Context, userId users.UserId) (*users.User, error) {
	sess, err := db.NewSession(ctx)
	if err != nil {
		return nil, err
	}
	var result UserModel
	err = sess.Users().FindOne(ctx, bson.M{"_id": userId}).Decode(&result)
	if err != nil {
		return nil, err
	}
	user, err := result.ToUser(ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u UserModel) ToUser(ctx context.Context) (*users.User, error) {
	history, err := terms.GetTermRecordsByIds(ctx, u.CourseHistory)
	if err != nil {
		return nil, err
	}
	tls, err := timelineModel.GetTimelinesByUserId(ctx, u.Id)
	if err != nil {
		return nil, err
	}
	user := users.User{
		Id:            u.Id,
		FirstName:     u.FirstName,
		LastName:      u.LastName,
		Password:      u.Password,
		Email:         u.Email,
		CourseHistory: history,
		Timelines:     tls,
		Role:          roles.GetRoleFromId(u.RoleId),
	}
	return &user, nil
}

func GetUserByEmail(ctx context.Context, email users.Email) (*users.User, error) {
	sess, err := db.NewSession(ctx)
	if err != nil {
		return nil, err
	}
	defer sess.Close(ctx)
	var result UserModel
	err = sess.Users().FindOne(ctx, bson.M{"email": email}).Decode(&result)
	if err != nil {
		return nil, err
	}
	user, err := result.ToUser(ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserRole(ctx context.Context, id users.UserId) (*roles.Role, error) {
	user, err := GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}
	return &user.Role, nil
}
