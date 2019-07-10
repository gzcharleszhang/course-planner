package plans

import (
	"github.com/gzcharleszhang/course-planner/internal/app/components/courses"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDegree_IsCompleted(t *testing.T) {
	degree := Degree{
		Name: "Easy BCS",
		Requirements: DegreeRequirements(courses.CourseRequirementSet{
			MinCoursesToSatisfy: 2,
			Requirements: courses.CourseRequirementRules{
				courses.CourseRequirement{
					MinGrade: 50,
					Course: &courses.Course{
						Id: 0,
					},
				},
				courses.CourseRequirement{
					MinGrade: 60,
					Course: &courses.Course{
						Id: 1,
					},
				},
			},
		}),
	}
	currTime := time.Now()

	// completed
	courseRecords := courses.CourseRecords{
		courses.CourseId(0): &courses.CourseRecord{
			Course: courses.Course{
				Id: 0,
			},
			Grade:          50,
			CompletionDate: &currTime,
		},
		courses.CourseId(1): &courses.CourseRecord{
			Course: courses.Course{
				Id: 1,
			},
			Grade:          60,
			CompletionDate: &currTime,
		},
	}
	assert.Equal(t, true, degree.IsCompleted(&courseRecords))

	// incomplete
	courseRecords = courses.CourseRecords{
		courses.CourseId(0): &courses.CourseRecord{
			Course: courses.Course{
				Id: 0,
			},
			Grade:          50,
			CompletionDate: &currTime,
		},
		courses.CourseId(1): &courses.CourseRecord{
			Course: courses.Course{
				Id: 1,
			},
			Grade:          59,
			CompletionDate: &currTime,
		},
	}
	assert.Equal(t, false, degree.IsCompleted(&courseRecords))
}

func TestDegree_GetName(t *testing.T) {
	name := "Easy BCS"
	degree := Degree{
		Name: DegreeName(name),
		Requirements: DegreeRequirements(courses.CourseRequirementSet{
			MinCoursesToSatisfy: 2,
			Requirements: courses.CourseRequirementRules{
				courses.CourseRequirement{
					MinGrade: 50,
					Course: &courses.Course{
						Id: 0,
					},
				},
				courses.CourseRequirement{
					MinGrade: 60,
					Course: &courses.Course{
						Id: 1,
					},
				},
			},
		}),
	}
	assert.Equal(t, name, degree.GetName())
}
