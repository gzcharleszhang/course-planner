package Courses

import "time"

type CourseId string
type CourseSubject string
type CourseCatalog int
type CourseGrade float64

type Course struct {
	Id                CourseId
	Subject           CourseSubject
	Catalog           CourseCatalog
	Description       string
	PrereqDescription string
	Prerequisites     CourseRequirementRule
}

type Courses map[CourseId]Course

type CompletedCourse struct {
	Course
	Grade          CourseGrade
	CompletionDate *time.Time
}

type CompletedCourses map[CourseId]CompletedCourse
