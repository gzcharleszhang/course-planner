package Courses

import "time"

type CourseId string
type CourseSubject string
type CourseCatalog int
type CourseDescription string
type CourseGrade float64

type Course struct {
	Id                CourseId
	Subject           CourseSubject
	Catalog           CourseCatalog
	Description       CourseDescription
	PrerequisiteRules CourseRequirementRules
}

type CompletedCourse struct {
	Course
	Grade          CourseGrade
	CompletionDate *time.Time
}

type CompletedCourses map[CourseId]CompletedCourse
