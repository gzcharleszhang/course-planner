package plans

import (
	"github.com/gzcharleszhang/course-planner/internal/app/components/courses"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDegree_IsCompleted(t *testing.T) {
	currTime := time.Now()
	degree := initDegree()

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

	// incomplete, one course missing
	courseRecords = courses.CourseRecords{
		courses.CourseId(0): &courses.CourseRecord{
			Course: courses.Course{
				Id: 0,
			},
			Grade:          50,
			CompletionDate: &currTime,
		},
	}
	assert.Equal(t, false, degree.IsCompleted(&courseRecords))

	// incomplete, one course did not meet grade requirement
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
	degree := initDegree()
	assert.Equal(t, "Easy BCS", degree.GetName())
}

func initDegree() *Degree {
	return &Degree{
		Name: "Easy BCS",
		Requirements: DegreeRequirements{
			courses.CourseRequirementSet{
				NumCoursesToSatisfy: 2,
				Requirements: courses.CourseRequirementRules{
					courses.CourseRequirement{
						MinGrade: 50,
						CourseId: 0,
					},
					courses.CourseRequirement{
						MinGrade: 60,
						CourseId: 1,
					},
				},
			},
		},
	}
}
