package courseSelections

import (
	"reflect"
	"testing"
	"time"

	"github.com/gzcharleszhang/course-planner/internal/app/components/courses"
	"github.com/gzcharleszhang/course-planner/internal/app/components/plans"
	"github.com/stretchr/testify/assert"
)

func TestCourseSelection_Aggregate(t *testing.T) {
	currTime := time.Now()
	courseSelection := CourseSelection{
		TermSelections: []*TermSelection{
			{
				CourseRecords: courses.CourseRecords{
					&courses.CourseRecord{
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
					&courses.CourseRecord{
						Course: courses.Course{
							Id: 1,
						},
						Grade:          60,
						CompletionDate: &currTime,
					},
					&courses.CourseRecord{
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
					&courses.CourseRecord{
						Course: courses.Course{
							Id: 3,
						},
					},
				},
			},
		},
	}
	expectedRecords := courses.CourseRecords{
		&courses.CourseRecord{
			Course: courses.Course{
				Id: 0,
			},
			Grade:          50,
			CompletionDate: &currTime,
		},
		&courses.CourseRecord{
			Course: courses.Course{
				Id: 1,
			},
			Grade:          60,
			CompletionDate: &currTime,
		},
		&courses.CourseRecord{
			Course: courses.Course{
				Id: 2,
			},
			Grade:          60,
			CompletionDate: &currTime,
		},
		&courses.CourseRecord{
			Course: courses.Course{
				Id: 3,
			},
		},
	}
	assert.Equal(t, &expectedRecords, courseSelection.Aggregate())
}

func TestCourseSelection_IncompletePlans(t *testing.T) {
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

	csPlan, econMinorPlan, mathPlan := plans.Plan(csDegree), plans.Plan(econMinor), plans.Plan(mathDegree)
	var csEconPlan plans.Plans
	csEconPlan = append(csEconPlan, &csPlan)
	csEconPlan = append(csEconPlan, &econMinorPlan)
	var csMathPlan plans.Plans
	csMathPlan = append(csMathPlan, &csPlan)
	csMathPlan = append(csMathPlan, &mathPlan)

	// satisfied
	courseSelection := CourseSelection{
		TermSelections: []*TermSelection{
			{
				CourseRecords: courses.CourseRecords{
					&courses.CourseRecord{
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
					&courses.CourseRecord{
						Course: courses.Course{
							Id: 1,
						},
						Grade:          60,
						CompletionDate: &currTime,
					},
					&courses.CourseRecord{
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
	assert.Equal(t, 0, len(courseSelection.IncompletePlans()))

	// satisfied overlapping plans
	courseSelection = CourseSelection{
		TermSelections: []*TermSelection{
			{
				CourseRecords: courses.CourseRecords{
					&courses.CourseRecord{
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
					&courses.CourseRecord{
						Course: courses.Course{
							Id: 1,
						},
						Grade:          60,
						CompletionDate: &currTime,
					},
					&courses.CourseRecord{
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
	assert.Equal(t, 0, len(courseSelection.IncompletePlans()))

	// not satisfied
	courseSelection = CourseSelection{
		TermSelections: []*TermSelection{
			{
				CourseRecords: courses.CourseRecords{
					&courses.CourseRecord{
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
					&courses.CourseRecord{
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
	inCompletePlans := courseSelection.IncompletePlans()
	assert.Equal(t, 1, len(inCompletePlans))
	assert.Equal(t, string(csDegree.Name), (*inCompletePlans[0]).GetName())

	// grade requirement not met
	courseSelection = CourseSelection{
		TermSelections: []*TermSelection{
			{
				CourseRecords: courses.CourseRecords{
					&courses.CourseRecord{
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
					&courses.CourseRecord{
						Course: courses.Course{
							Id: 3,
						},
						Grade:          60,
						CompletionDate: &currTime,
					},
					&courses.CourseRecord{
						Course: courses.Course{
							Id: 1,
						},
						Grade:          59,
						CompletionDate: &currTime,
					},
				},
			},
		},
		Plans: csEconPlan,
	}
	inCompletePlans = courseSelection.IncompletePlans()
	assert.Equal(t, 1, len(inCompletePlans))
	assert.Equal(t, string(csDegree.Name), (*inCompletePlans[0]).GetName())

	// grade requirement met after repeating course
	courseSelection = CourseSelection{
		TermSelections: []*TermSelection{
			{
				CourseRecords: courses.CourseRecords{
					&courses.CourseRecord{
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
					&courses.CourseRecord{
						Course: courses.Course{
							Id: 3,
						},
						Grade:          60,
						CompletionDate: &currTime,
					},
					&courses.CourseRecord{
						Course: courses.Course{
							Id: 1,
						},
						Grade:          59,
						CompletionDate: &currTime,
					},
				},
			},
			{
				CourseRecords: courses.CourseRecords{
					&courses.CourseRecord{
						Course: courses.Course{
							Id: 1,
						},
						Grade:          60,
						CompletionDate: &currTime,
					},
				},
			},
		},
		Plans: csEconPlan,
	}
	inCompletePlans = courseSelection.IncompletePlans()
	assert.Equal(t, 0, len(inCompletePlans))
}

func TestCourseSelection_InvalidCourses(t *testing.T) {
	course2 := courses.CourseRecord{
		Course: courses.Course{
			Id: 1,
			Prereqs: courses.CourseRequirement{
				Course: &courses.Course{
					Id: 0,
				},
				MinGrade: 60,
			},
		},
	}
	course3 := courses.CourseRecord{
		Course: courses.Course{
			Id: 2,
			Prereqs: courses.CourseRequirement{
				Course: &courses.Course{
					Id: 1,
				},
				MinGrade: 60,
			},
		},
	}
	currTime := time.Now()
	type fields struct {
		Id             CourseSelectionId
		Name           CourseSelectionName
		TermSelections []*TermSelection
		Plans          plans.Plans
	}
	tests := []struct {
		name   string
		fields fields
		want   courses.CourseRecords
	}{
		{
			name: "all valid",
			fields: fields{
				TermSelections: []*TermSelection{
					{
						CourseRecords: courses.CourseRecords{
							&courses.CourseRecord{
								Course: courses.Course{
									Id: 0,
								},
							},
						},
					},
					{
						CourseRecords: courses.CourseRecords{
							&course2,
						},
					},
					{
						CourseRecords: courses.CourseRecords{
							&course3,
						},
					},
				},
			},
			want: courses.CourseRecords{},
		},
		{
			name: "second and third term invalid",
			fields: fields{
				TermSelections: []*TermSelection{
					{
						CourseRecords: courses.CourseRecords{
							&courses.CourseRecord{
								Course: courses.Course{
									Id: 0,
								},
								Grade:          59,
								CompletionDate: &currTime,
							},
						},
					},
					{
						CourseRecords: courses.CourseRecords{
							&course2,
						},
					},
					{
						CourseRecords: courses.CourseRecords{
							&course3,
						},
					},
				},
			},
			want: courses.CourseRecords{
				&course2,
				&course3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := CourseSelection{
				Id:             tt.fields.Id,
				Name:           tt.fields.Name,
				TermSelections: tt.fields.TermSelections,
				Plans:          tt.fields.Plans,
			}
			if got := cs.InvalidCourses(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CourseSelection.InvalidCourses() = %v, want %v", got, tt.want)
			}
		})
	}
}
