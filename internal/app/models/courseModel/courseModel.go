package courseModel

import (
	"context"
	"fmt"
	"github.com/gzcharleszhang/course-planner/internal/app/components/courses"
	"github.com/gzcharleszhang/course-planner/internal/app/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CourseModel struct {
	Id          courses.CourseId      `bson:"_id"`
	Name        courses.CourseName    `bson:"name"`
	Subject     courses.CourseSubject `bson:"subject"`
	Catalog     courses.CourseCatalog `bson:"catalog"`
	Description string                `bson:"description"`
}

func (cm CourseModel) ToCourse() *courses.Course {
	course := courses.Course{
		Id:          cm.Id,
		Name:        cm.Name,
		Subject:     cm.Subject,
		Catalog:     cm.Catalog,
		Description: cm.Description,
	}
	return &course
}

func CreateCourse(
	ctx context.Context,
	id courses.CourseId,
	name courses.CourseName,
	subject courses.CourseSubject,
	catalog courses.CourseCatalog,
	desc string,
) error {
	exists, err := checkDuplicateCourse(ctx, id)
	if err != nil {
		return err
	}
	if exists {
		return db.DocumentExistsError{Message: fmt.Sprintf("course_id: %v", id)}
	}
	cm := CourseModel{
		Id:          id,
		Name:        name,
		Subject:     subject,
		Catalog:     catalog,
		Description: desc,
	}
	sess := db.PrimarySession
	_, err = sess.Courses().InsertOne(ctx, cm)
	if err != nil {
		return err
	}
	return nil
}

func GetCourseById(ctx context.Context, id courses.CourseId) (*courses.Course, error) {
	sess := db.PrimarySession
	var result CourseModel
	err := sess.Courses().FindOne(ctx, bson.M{"_id": id}).Decode(&result)
	if err != nil {
		return nil, err
	}
	course := result.ToCourse()
	return course, nil
}

// return true if course already exists
func checkDuplicateCourse(ctx context.Context, id courses.CourseId) (bool, error) {
	_, err := GetCourseById(ctx, id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
