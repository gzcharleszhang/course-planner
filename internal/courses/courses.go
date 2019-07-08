package courses

type CourseId int
type CourseName string
type CourseSubject string
type CourseCatalog int

type Course struct {
	Id                CourseId
	Name              CourseName
	Subject           CourseSubject
	Catalog           CourseCatalog
	Description       string
	PrereqDescription string
	Prereqs           CourseRequirementRule
}
