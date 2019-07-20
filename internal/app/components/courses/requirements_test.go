package courses

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCourseRequirement_IsSatisfied(t *testing.T) {
	records := initRecords()
	// testing using future course record
	req := CourseRequirement{
		MinGrade: 100, // should ignore this requirement
		CourseId: CourseId(789),
	}
	assert.Equal(t, true, req.IsSatisfied(records), "Test future record")

	// missing requirement
	req = CourseRequirement{
		MinGrade: 0,
		CourseId: CourseId(0),
	}
	assert.Equal(t, false, req.IsSatisfied(records), "Test missing record")

	// doesn't meet grades requirement
	req = CourseRequirement{
		MinGrade: 51,
		CourseId: CourseId(456),
	}
	assert.Equal(t, false, req.IsSatisfied(records), "Test grade requirement not met")

	// meets grades requirement
	req = CourseRequirement{
		MinGrade: 50,
		CourseId: CourseId(456),
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
	req = CourseRequirement{
		MinGrade: 51, // should ignore this requirement
		CourseId: CourseId(456),
	}
	assert.Equal(t, true, req.IsSatisfied(records), "Test grade requirement met after course repeated")
}

func TestCourseRequirementRange_IsSatisfied(t *testing.T) {
	records := initRecords()
	// wrong subject
	req := CourseRequirementRange{
		Subject:    "NotTest",
		CatalogMin: 1,
		CatalogMax: 4,
		MinGrade:   50,
	}
	assert.Equal(t, false, req.IsSatisfied(records), "Test wrong subject")

	// right subject, not in range
	req = CourseRequirementRange{
		Subject:    "Test",
		CatalogMin: 4,
		CatalogMax: 7,
		MinGrade:   50,
	}
	assert.Equal(t, false, req.IsSatisfied(records), "Test catalog not in range")

	// right subject, in range, does not pass grade requirement
	req = CourseRequirementRange{
		Subject:    "Test",
		CatalogMin: 1,
		CatalogMax: 2,
		MinGrade:   81,
	}
	assert.Equal(t, false, req.IsSatisfied(records), "Test grade requirement not met")

	// right subject, on min boundary
	req = CourseRequirementRange{
		Subject:    "Test",
		CatalogMin: 2,
		CatalogMax: 3,
		MinGrade:   50,
	}
	assert.Equal(t, true, req.IsSatisfied(records), "Test on min boundary of range")

	// right subject, on max boundary
	req = CourseRequirementRange{
		Subject:    "Test",
		CatalogMin: 0,
		CatalogMax: 1,
		MinGrade:   50,
	}
	assert.Equal(t, true, req.IsSatisfied(records), "Test on max boundary of range")

	// right subject, in range, future course
	req = CourseRequirementRange{
		Subject:    "Test",
		CatalogMin: 3,
		CatalogMax: 6,
		MinGrade:   50,
	}
	assert.Equal(t, true, req.IsSatisfied(records), "Test future course")
}

func TestCourseRequirementSet_IsSatisfied(t *testing.T) {
	records := initRecords()
	// meet all requirements
	req := CourseRequirementSet{
		NumCoursesToSatisfy: 3,
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
		NumCoursesToSatisfy: 3,
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
		NumCoursesToSatisfy: 2,
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
		NumCoursesToSatisfy: 2,
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
	course1 := Course{
		Id:      CourseId(123),
		Name:    "Test 1",
		Subject: "Test",
		Catalog: 1,
	}
	record1 := CourseRecord{
		Course:         course1,
		Grade:          80,
		CompletionDate: &currTime,
	}
	course2 := Course{
		Id:      CourseId(456),
		Name:    "Test 2",
		Subject: "Test",
		Catalog: 2,
	}
	record2 := CourseRecord{
		Course:         course2,
		Grade:          50,
		CompletionDate: &currTime,
	}
	course3 := Course{
		Id:      CourseId(789),
		Name:    "Test 3",
		Subject: "Test",
		Catalog: 3,
	}
	record3 := CourseRecord{
		Course: course3,
	}

	return &CourseRecords{&record1, &record2, &record3}
}
