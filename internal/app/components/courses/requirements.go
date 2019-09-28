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
	CatalogMin CatalogNumber `json:"catalog_min"`
	CatalogMax CatalogNumber `json:"catalog_max"`
	MinGrade   CourseGrade   `json:"min_grade"`
}

// CourseRequirementCAV is a concrete impl for checking if records have the minimum required CAV
type CAVRequirement struct {
	MinCAV CourseGrade `json:"min_cav"`
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
	return completed && checkGradeRequirement(course, req.MinGrade)
}

func (rang CourseRequirementRange) IsSatisfied(courseRecords *CourseRecords) bool {
	for _, cr := range *courseRecords {
		if cr.Subject == rang.Subject && cr.Catalog.Number >= rang.CatalogMin && cr.Catalog.Number <= rang.CatalogMax &&
			checkGradeRequirement(cr, rang.MinGrade) {
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

func (req CAVRequirement) IsSatisfied(courseRecords *CourseRecords) bool {
	return courseRecords.CurrentCAV() >= req.MinCAV
}

// grade requirement is met if cr has a higher grade or if cr has a nil completion date (signifying a future course)
func checkGradeRequirement(cr *CourseRecord, minGrade CourseGrade) bool {
	return cr.Grade >= minGrade || cr.CompletionDate == nil
}
