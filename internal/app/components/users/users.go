package users

import (
	"context"
	"github.com/gzcharleszhang/course-planner/internal/app/components/contextKeys"
	"github.com/gzcharleszhang/course-planner/internal/app/components/roles"
	"github.com/gzcharleszhang/course-planner/internal/app/components/terms"
	"github.com/gzcharleszhang/course-planner/internal/app/components/timelines"
	"github.com/rs/xid"
	"golang.org/x/crypto/bcrypt"
)

type FirstName string
type LastName string
type Email string
type UserId string
type PasswordHash string

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

func NewUserId() UserId {
	return UserId(xid.New().String())
}

func HashPassword(password string) (PasswordHash, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return PasswordHash(bytes), err
}

// Creates a new timemline with the courses added to the course CourseHistory
func (usr User) NewTimeline(name timelines.TimelineName) {
	usr.Timelines = append(usr.Timelines, timelines.NewTimeline(name, usr.CourseHistory))
}

func GetUserIdFromContext(ctx context.Context) UserId {
	return ctx.Value(contextKeys.UserIdKey).(UserId)
}
