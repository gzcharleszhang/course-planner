package Courses

import "time"

type CourseName string
type CourseSubject string
type CourseId int
type CourseGrade float64

type Course struct {
	Name              CourseName
	Id                CourseId
	Subject           CourseSubject
	Description       string
	PrereqDescription string
	Prerequisites     CourseRequirementRule
}

type Courses map[CourseId]*Course

type CompletedCourse struct {
	Course
	Grade          CourseGrade
	CompletionDate *time.Time
}

type CompletedCourses map[CourseId]CompletedCourse
