package users

import (
	"context"
	"github.com/gzcharleszhang/course-planner/internal/app/components/roles"
	"github.com/gzcharleszhang/course-planner/internal/app/components/terms"
	"github.com/gzcharleszhang/course-planner/internal/app/components/timelines"
	"github.com/gzcharleszhang/course-planner/internal/app/db"
	"github.com/pkg/errors"
	"github.com/rs/xid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type FirstName string
type LastName string
type Email string
type UserId string
type PasswordHash string

type UserData struct {
	Id            UserId                 `bson:"_id"`
	FirstName     FirstName              `bson:"first_name"`
	LastName      LastName               `bson:"last_name"`
	Email         Email                  `bson:"email"`
	Password      PasswordHash           `bson:"password"`
	CourseHistory []terms.TermRecordId   `bson:"course_history"`
	Timelines     []timelines.TimelineId `bson:"timelines"`
	RoleId        roles.RoleId           `bson:"role_id"`
}

type User struct {
	Id            UserId                `json:"_id"`
	FirstName     FirstName             `json:"first_name"`
	LastName      LastName              `json:"last_name"`
	Email         Email                 `json:"email"`
	Password      PasswordHash          `json:"password"`
	CourseHistory terms.TermRecords     `json:"course_history"`
	Timelines     []*timelines.Timeline `json:"timelines"`
	Role          roles.Role            `json:"role"`
}

func newUserId() UserId {
	return UserId(xid.New().String())
}

func CreateUser(ctx context.Context, firstName FirstName, lastName LastName,
	email Email, password PasswordHash) (UserId, error) {
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
		return "", errors.New("Email already exists")
	}
	newUserId := newUserId()
	user := UserData{
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

func GetTimelinesByUserId(ctx context.Context, userId UserId) ([]*timelines.Timeline, error) {
	// TODO: implement
	return nil, nil
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
	user, err := result.ToUser(ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func HashPassword(password string) (PasswordHash, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return PasswordHash(bytes), err
}

func VerifyPassword(ctx context.Context, email Email, password string) error {
	user, err := GetUserByEmail(ctx, email)
	if err != nil {
		return err
	}
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

func GetUserRole(ctx context.Context, id UserId) (*roles.Role, error) {
	user, err := GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}
	return &user.Role, nil
}

func GetUserByEmail(ctx context.Context, email Email) (*User, error) {
	sess, err := db.NewSession(ctx)
	if err != nil {
		return nil, err
	}
	defer sess.Close(ctx)
	var result UserData
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

// return true if there already exists an user with the given email
func checkDuplicateEmail(ctx context.Context, email Email) (bool, error) {
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

func (u UserData) ToUser(ctx context.Context) (*User, error) {
	history, err := terms.GetTermRecordsByIds(ctx, u.CourseHistory)
	if err != nil {
		return nil, err
	}
	tls, err := GetTimelinesByUserId(ctx, u.Id)
	if err != nil {
		return nil, err
	}
	user := User{
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

// Creates a new timemline with the courses added to the course CourseHistory
func (usr User) NewTimeline(name timelines.TimelineName) {
	usr.Timelines = append(usr.Timelines, timelines.NewTimeline(name, usr.CourseHistory))
}
