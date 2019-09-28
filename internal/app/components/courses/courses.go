package courses

type CourseId string
type CourseName string
type CourseSubject string
type CourseCatalog struct {
	Number CatalogNumber
	Suffix CatalogSuffix
}
type CatalogNumber int
type CatalogSuffix string

type Course struct {
	Id                CourseId               `json:"id"`
	Name              CourseName             `json:"name"`
	Subject           CourseSubject          `json:"subject"`
	Catalog           CourseCatalog          `json:"catalog"`
	Description       string                 `json:"description"`
	PrereqDescription string                 `json:"prereq_description"`
	Prereqs           CourseRequirementRules `json:"prereqs"`
}
