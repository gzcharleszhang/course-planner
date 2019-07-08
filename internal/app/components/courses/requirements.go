package courses

type CourseRequirementRules []CourseRequirementRule

type CourseRequirement struct {
	Course   *Course
	MinGrade CourseGrade
}

type CourseRequirementSet struct {
	MinCoursesToSatisfy int
	Requirements        CourseRequirementRules
}

type CourseRequirementRule interface {
	IsSatisfied(*CourseRecords) bool
}

func (req CourseRequirement) IsSatisfied(courseRecords *CourseRecords) bool {
	course, completed := (*courseRecords)[req.Course.Id]
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
