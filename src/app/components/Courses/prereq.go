package Courses

type CourseRequirementRules []*CourseRequirementRule

type CourseRequirement struct {
	Course   *Course
	MinGrade CourseGrade
}

type CourseRequirementSet struct {
	MinCoursesToSatisfy int
	Requirements        CourseRequirementRules
}

type CourseRequirementRule interface {
	IsSatisfied(*CompletedCourses) bool
}

func isCourseReqSatisfied(req *CourseRequirement, completedCourses *CompletedCourses) bool {
	completedCourse, inCompletedCourses := (*completedCourses)[req.Course.Id]
	return inCompletedCourses && completedCourse.Grade >= req.MinGrade
}

func (req CourseRequirement) IsSatisfied(completedCourses *CompletedCourses) bool {
	return isCourseReqSatisfied(&req, completedCourses)
}

func (set CourseRequirementSet) IsSatisfied(completedCourses *CompletedCourses) bool {
	count := 0
	for _, req := range set.Requirements {
		if (*req).IsSatisfied(completedCourses) {
			count += 1
		}
	}
	return count >= set.MinCoursesToSatisfy
}
