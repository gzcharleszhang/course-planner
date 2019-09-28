package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gzcharleszhang/course-planner/internal/app/components/courses"
	"github.com/gzcharleszhang/course-planner/internal/app/db"
	"github.com/gzcharleszhang/course-planner/internal/app/models/courseModel"
	"github.com/gzcharleszhang/course-planner/internal/app/scripts"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	ctx, err := scripts.InitScript()
	if err != nil {
		log.Fatalf("Error starting script: %v", err)
	}
	subjectsToLoad := []string{"CS", "MATH", "AFM", "STAT", "CO", "ECON", "SPCOM", "ENGL"}
	// loading cs for now
	err = loadCoursesBySubjects(ctx, subjectsToLoad)
	if err != nil {
		log.Fatalf("Error loading courses: %v", err)
	}
}

func strToCourseCatalog(s string) (*courses.CourseCatalog, error) {
	num, err := strconv.Atoi(s)
	if err == nil {
		return &courses.CourseCatalog{
			Number: courses.CatalogNumber(num),
			Suffix: courses.CatalogSuffix(""),
		}, nil
	}
	// otherwise, find the non int character
	firstNonInt := -1
	for i, c := range s {
		_, err := strconv.Atoi(string(c))
		if err != nil {
			firstNonInt = i
			break
		}
	}

	if firstNonInt >= 0 {
		num, err = strconv.Atoi(s[:firstNonInt])
		if err != nil {
			return nil, err
		}
		return &courses.CourseCatalog{
			Number: courses.CatalogNumber(num),
			Suffix: courses.CatalogSuffix(s[firstNonInt:]),
		}, nil
	}
	return nil, nil
}

func loadCoursesBySubjects(ctx context.Context, subjects []string) error {
	for _, subject := range subjects {
		err := loadCoursesBySubject(ctx, subject)
		if err != nil {
			return err
		}
	}
	return nil
}

func loadCoursesBySubject(ctx context.Context, subject string) error {
	upperSubject := strings.ToUpper(subject)
	log.Printf("Loading %v courses...", upperSubject)
	rawData, _ := ioutil.ReadFile(fmt.Sprintf("data/courses/%v.json", upperSubject))
	type courseRaw struct {
		Id          string `json:"course_id"`
		Title       string `json:"title"`
		Subject     string `json:"subject"`
		CatalogStr  string `json:"catalog_number"`
		Description string `json:"description"`
	}
	type coursesRaw []*courseRaw
	var data struct {
		Courses coursesRaw `json:"data"`
	}
	err := json.Unmarshal(rawData, &data)
	count := 0
	dupCount := 0
	for _, cr := range data.Courses {
		catalog, err := strToCourseCatalog(cr.CatalogStr)
		if err != nil {
			return err
		}
		err = courseModel.CreateCourse(
			ctx,
			courses.CourseId(cr.Id),
			courses.CourseName(cr.Title),
			courses.CourseSubject(cr.Subject),
			*catalog,
			cr.Description,
		)
		if err != nil {
			switch err.(type) {
			case db.DocumentExistsError:
				dupCount += 1
			default:
				return err
			}
		} else {
			count += 1
		}
		if count%25 == 0 && count != 0 {
			log.Printf("Created %v courses, skipped %v duplicates", count, dupCount)
		}
	}
	if err != nil {
		return err
	}
	log.Printf("Created %v %v courses in total", count, upperSubject)
	return nil
}
