package courses

type CourseId int
type CourseName string
type CourseSubject string
type CourseCatalog int

type Course struct {
	Id                CourseId               `json:"id"`
	Name              CourseName             `json:"name"`
	Subject           CourseSubject          `json:"subject"`
	Catalog           CourseCatalog          `json:"catalog"`
	Description       string                 `json:"description"`
	PrereqDescription string                 `json:"prereq_description"`
	Prereqs           CourseRequirementRules `json:"prereqs"`
}
