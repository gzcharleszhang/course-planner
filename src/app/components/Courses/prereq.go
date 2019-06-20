package Courses

type CourseRequirement struct {
	Course   *Course
	MinGrade CourseGrade
}

type CourseRequirementSet struct {
	MinCoursesToSatisfy int
	Courses             []*CourseRequirement
}

type CourseRequirementRuleset interface {
	IsSatisfied(CompletedCourses) bool
}

type CourseRequirementRules []*CourseRequirementRuleset

func (req CourseRequirement) IsSatisfied(completedCourses CompletedCourses) bool {
	completedCourse, inCompletedCourses := completedCourses[req.Course.Id]
	return inCompletedCourses && completedCourse.Grade >= req.MinGrade
}
