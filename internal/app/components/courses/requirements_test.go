package courses

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCourseRequirement_IsSatisfied(t *testing.T) {
	records := initRecords()
	// testing using future course record (record3 )
	courseId := CourseId(789)
	req := CourseRequirement{
		MinGrade: 100, // should ignore this requirement
		CourseId: courseId,
	}
	assert.Equal(t, true, req.IsSatisfied(records), "Test future record")

	// missing requirement
	courseId = CourseId(0)
	req = CourseRequirement{
		MinGrade: 0, // should ignore this requirement
		CourseId: courseId,
	}
	assert.Equal(t, false, req.IsSatisfied(records), "Test missing record")

	// doesn't meet grades requirement
	courseId = CourseId(456)

	req = CourseRequirement{
		MinGrade: 51, // should ignore this requirement
		CourseId: courseId,
	}
	assert.Equal(t, false, req.IsSatisfied(records), "Test grade requirement not met")

	// meets grades requirement
	courseId = CourseId(456)
	req = CourseRequirement{
		MinGrade: 50, // should ignore this requirement
		CourseId: courseId,
	}
	assert.Equal(t, true, req.IsSatisfied(records), "Test grade requirement met")

	// test repeated courses
	currTime := time.Now()
	repeatedCourse := CourseRecord{
		Course: Course{
			Id: CourseId(456),
		},
		Grade:          60,
		CompletionDate: &currTime,
	}
	*records = append(*records, &repeatedCourse)
	courseId = CourseId(456)
	req = CourseRequirement{
		MinGrade: 51, // should ignore this requirement
		CourseId: courseId,
	}
	assert.Equal(t, true, req.IsSatisfied(records), "Test grade requirement met after course repeated")
}

func TestCourseRequirementSet_IsSatisfied(t *testing.T) {
	records := initRecords()
	// meet all requirements
	req := CourseRequirementSet{
		MinCoursesToSatisfy: 3,
		Requirements: CourseRequirementRules{
			CourseRequirement{
				MinGrade: 80,
				CourseId: CourseId(123),
			},
			CourseRequirement{
				MinGrade: 50,
				CourseId: CourseId(456),
			},
			CourseRequirement{
				MinGrade: 100,
				CourseId: CourseId(789),
			},
		},
	}
	assert.Equal(t, true, req.IsSatisfied(records), "Test all requirements met")

	// doesn't meet one course grade requirement
	req = CourseRequirementSet{
		MinCoursesToSatisfy: 3,
		Requirements: CourseRequirementRules{
			CourseRequirement{
				MinGrade: 80,
				CourseId: CourseId(123),
			},
			CourseRequirement{
				MinGrade: 51,
				CourseId: CourseId(456),
			},
			CourseRequirement{
				MinGrade: 100,
				CourseId: CourseId(789),
			},
		},
	}
	assert.Equal(t, false, req.IsSatisfied(records), "Test one requirement not met")

	// meet 2 out of the 3 courses
	req = CourseRequirementSet{
		MinCoursesToSatisfy: 2,
		Requirements: CourseRequirementRules{
			CourseRequirement{
				MinGrade: 80,
				CourseId: CourseId(123),
			},
			CourseRequirement{
				MinGrade: 51,
				CourseId: CourseId(456),
			},
			CourseRequirement{
				MinGrade: 100,
				CourseId: CourseId(789),
			},
		},
	}
	assert.Equal(t, true, req.IsSatisfied(records), "Test one requirement not met")

	// doesn't satisfy min number of courses
	req = CourseRequirementSet{
		MinCoursesToSatisfy: 2,
		Requirements: CourseRequirementRules{
			CourseRequirement{
				MinGrade: 80,
				CourseId: CourseId(123),
			},
			CourseRequirement{
				MinGrade: 51,
				CourseId: CourseId(456),
			},
			CourseRequirement{
				MinGrade: 0,
				CourseId: CourseId(0),
			},
		},
	}
	assert.Equal(t, false, req.IsSatisfied(records), "Test one requirement not met")
}

// setting up course records
func initRecords() *CourseRecords {
	currTime := time.Now()
	courseId1 := CourseId(123)
	course1 := Course{
		Id:      courseId1,
		Name:    "Test 1",
		Subject: "Test",
		Catalog: 1,
	}
	record1 := CourseRecord{
		Course:         course1,
		Grade:          80,
		CompletionDate: &currTime,
	}
	courseId2 := CourseId(456)
	course2 := Course{
		Id:      courseId2,
		Name:    "Test 2",
		Subject: "Test",
		Catalog: 2,
	}
	record2 := CourseRecord{
		Course:         course2,
		Grade:          50,
		CompletionDate: &currTime,
	}
	courseId3 := CourseId(789)
	course3 := Course{
		Id:      courseId3,
		Name:    "Test 3",
		Subject: "Test",
		Catalog: 3,
	}
	record3 := CourseRecord{
		Course: course3,
	}
	records := CourseRecords{
		&record1,
		&record2,
		&record3,
	}

	return &records
}
