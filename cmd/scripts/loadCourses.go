package main

import (
	"context"
	"encoding/json"
	"github.com/gzcharleszhang/course-planner/internal/app/components/courses"
	"github.com/gzcharleszhang/course-planner/internal/app/models/courseModel"
	"github.com/gzcharleszhang/course-planner/internal/app/scripts"
	"io/ioutil"
	"log"
	"strconv"
)

func main() {
	ctx, err := scripts.InitScript()
	if err != nil {
		log.Fatalf("Error starting script: %v", err)
	}
	// loading cs for now
	err = loadCSCourses(ctx)
	if err != nil {
		log.Fatalf("Error loading CS courses: %v", err)
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

func loadCSCourses(ctx context.Context) error {
	log.Print("Loading CS courses...")
	rawData, _ := ioutil.ReadFile("data/courses/CS.json")
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
		count += 1
		if count%100 == 0 {
			log.Printf("Created %v courses", count)
		}
	}
	if err != nil {
		return err
	}
	log.Printf("Created %v courses in total", count)
	return nil
}
