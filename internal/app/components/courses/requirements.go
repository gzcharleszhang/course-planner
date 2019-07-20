package courses

// CourseRequirementRule is the abstract interface for requirements
type CourseRequirementRule interface {
	IsSatisfied(*CourseRecords) bool
}

type CourseRequirementRules []CourseRequirementRule

// CourseRequirement is a concrete impl for one single course (e.g. CS 246)
type CourseRequirement struct {
	CourseId CourseId    `json:"course_id"`
	MinGrade CourseGrade `json:"min_grade"`
}

// CourseRequirementRange is a concrete impl for one course within a range (e.g. CS 340-CS 398)
type CourseRequirementRange struct {
	Subject    CourseSubject `json:"subject"`
	CatalogMin CourseCatalog `json:"catalog_min"`
	CatalogMax CourseCatalog `json:"catalog_max"`
	MinGrade   CourseGrade   `json:"min_grade"`
}

// CourseRequirementSet is a concrete impl for several courses that can satisfy any of the requirements (e.g. one of
// CS 115, CS 135, CS 145)
type CourseRequirementSet struct {
	NumCoursesToSatisfy int                    `json:"num_courses_to_satisfy"`
	Requirements        CourseRequirementRules `json:"requirements"`
}

func (req CourseRequirement) IsSatisfied(courseRecords *CourseRecords) bool {
	idMap := courseRecords.ToCourseIdMap()
	course, completed := idMap[req.CourseId]
	return completed && (course.Grade >= req.MinGrade || course.CompletionDate == nil)
}

func (rang CourseRequirementRange) IsSatisfied(courseRecords *CourseRecords) bool {
	for _, cr := range *courseRecords {
		if cr.Subject == rang.Subject && cr.Catalog >= rang.CatalogMin && cr.Catalog <= rang.CatalogMax &&
			(cr.Grade >= rang.MinGrade || cr.CompletionDate == nil) {
			return true
		}
	}
	return false
}

func (set CourseRequirementSet) IsSatisfied(courseRecords *CourseRecords) bool {
	count := 0
	for _, req := range set.Requirements {
		if req.IsSatisfied(courseRecords) {
			count += 1
		}
	}
	return count >= set.NumCoursesToSatisfy
}
