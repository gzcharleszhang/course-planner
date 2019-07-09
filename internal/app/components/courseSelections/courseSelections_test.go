package courseSelections

import (
	"github.com/gzcharleszhang/course-planner/internal/app/components/courses"
	"github.com/gzcharleszhang/course-planner/internal/app/components/plans"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCourseSelection_Aggregate(t *testing.T) {
	currTime := time.Now()
	courseSelection := CourseSelection{
		TermSelections: []*courses.TermSelection{
			{
				CourseRecords: courses.CourseRecords{
					courses.CourseId(0): &courses.CourseRecord{
						Course: courses.Course{
							Id: 0,
						},
						Grade:          50,
						CompletionDate: &currTime,
					},
				},
			},
			{
				CourseRecords: courses.CourseRecords{
					courses.CourseId(1): &courses.CourseRecord{
						Course: courses.Course{
							Id: 1,
						},
						Grade:          60,
						CompletionDate: &currTime,
					},
					courses.CourseId(2): &courses.CourseRecord{
						Course: courses.Course{
							Id: 2,
						},
						Grade:          60,
						CompletionDate: &currTime,
					},
				},
			},
			{
				CourseRecords: courses.CourseRecords{
					courses.CourseId(3): &courses.CourseRecord{
						Course: courses.Course{
							Id: 3,
						},
					},
				},
			},
		},
	}
	expectedRecords := courses.CourseRecords{
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
		courses.CourseId(2): &courses.CourseRecord{
			Course: courses.Course{
				Id: 2,
			},
			Grade:          60,
			CompletionDate: &currTime,
		},
		courses.CourseId(3): &courses.CourseRecord{
			Course: courses.Course{
				Id: 3,
			},
		},
	}
	assert.Equal(t, &expectedRecords, courseSelection.Aggregate())
}

func TestCourseSelection_IsSatisfied(t *testing.T) {
	currTime := time.Now()
	csDegree := plans.Degree{
		Name: "Easy BCS",
		Requirements: plans.DegreeRequirements(courses.CourseRequirementSet{
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
	mathDegree := plans.Degree{
		Name: "Easy BMath",
		Requirements: plans.DegreeRequirements(courses.CourseRequirementSet{
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
						Id: 2,
					},
				},
			},
		}),
	}
	econMinor := plans.Degree{
		Name: "Easy Econ Minor",
		Requirements: plans.DegreeRequirements(courses.CourseRequirementSet{
			MinCoursesToSatisfy: 1,
			Requirements: courses.CourseRequirementRules{
				courses.CourseRequirement{
					MinGrade: 50,
					Course: &courses.Course{
						Id: 3,
					},
				},
				courses.CourseRequirement{
					MinGrade: 50,
					Course: &courses.Course{
						Id: 4,
					},
				},
			},
		}),
	}

	csEconPlan := make([]plans.Plan, 2)
	csEconPlan[0] = plans.Plan(csDegree)
	csEconPlan[1] = plans.Plan(econMinor)
	csMathPlan := make([]plans.Plan, 2)
	csMathPlan[0] = plans.Plan(csDegree)
	csMathPlan[1] = plans.Plan(mathDegree)

	// satisfied
	courseSelection := CourseSelection{
		TermSelections: []*courses.TermSelection{
			{
				CourseRecords: courses.CourseRecords{
					courses.CourseId(0): &courses.CourseRecord{
						Course: courses.Course{
							Id: 0,
						},
						Grade:          50,
						CompletionDate: &currTime,
					},
				},
			},
			{
				CourseRecords: courses.CourseRecords{
					courses.CourseId(1): &courses.CourseRecord{
						Course: courses.Course{
							Id: 1,
						},
						Grade:          60,
						CompletionDate: &currTime,
					},
					courses.CourseId(3): &courses.CourseRecord{
						Course: courses.Course{
							Id: 3,
						},
						Grade:          60,
						CompletionDate: &currTime,
					},
				},
			},
		},
		Plans: csEconPlan,
	}
	assert.Equal(t, true, courseSelection.IsSatisfied())

	// satisfied overlapping plans
	courseSelection = CourseSelection{
		TermSelections: []*courses.TermSelection{
			{
				CourseRecords: courses.CourseRecords{
					courses.CourseId(0): &courses.CourseRecord{
						Course: courses.Course{
							Id: 0,
						},
						Grade:          50,
						CompletionDate: &currTime,
					},
				},
			},
			{
				CourseRecords: courses.CourseRecords{
					courses.CourseId(1): &courses.CourseRecord{
						Course: courses.Course{
							Id: 1,
						},
						Grade:          60,
						CompletionDate: &currTime,
					},
					courses.CourseId(2): &courses.CourseRecord{
						Course: courses.Course{
							Id: 2,
						},
						Grade:          60,
						CompletionDate: &currTime,
					},
				},
			},
		},
		Plans: csMathPlan,
	}
	assert.Equal(t, true, courseSelection.IsSatisfied())

	// not satisfied
	courseSelection = CourseSelection{
		TermSelections: []*courses.TermSelection{
			{
				CourseRecords: courses.CourseRecords{
					courses.CourseId(0): &courses.CourseRecord{
						Course: courses.Course{
							Id: 0,
						},
						Grade:          50,
						CompletionDate: &currTime,
					},
				},
			},
			{
				CourseRecords: courses.CourseRecords{
					courses.CourseId(3): &courses.CourseRecord{
						Course: courses.Course{
							Id: 3,
						},
						Grade:          60,
						CompletionDate: &currTime,
					},
				},
			},
		},
		Plans: csEconPlan,
	}
	assert.Equal(t, false, courseSelection.IsSatisfied())
}
