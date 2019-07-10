package courses

// CourseRequirement(Set) are concrete
type CourseRequirement struct {
	Course   *Course     `json:"course"`
	MinGrade CourseGrade `json:"min_grade"`
}

type CourseRequirementSet struct {
	MinCoursesToSatisfy int                    `json:"min_courses_to_satisfy"`
	Requirements        CourseRequirementRules `json:"requirements"`
}

// CourseRequirementRule(s) are abstract
type CourseRequirementRule interface {
	IsSatisfied(*CourseRecords) bool
}

type CourseRequirementRules []CourseRequirementRule

func (req CourseRequirement) IsSatisfied(courseRecords *CourseRecords) bool {
	idMap := courseRecords.ToCourseIdMap()
	course, completed := idMap[req.Course.Id]
	return completed && (course.Grade >= req.MinGrade || course.CompletionDate == nil)
}

func (set CourseRequirementSet) IsSatisfied(courseRecords *CourseRecords) bool {
	count := 0
	for _, req := range set.Requirements {
		if req.IsSatisfied(courseRecords) {
			count += 1
		}
	}
	return count >= set.MinCoursesToSatisfy
}
