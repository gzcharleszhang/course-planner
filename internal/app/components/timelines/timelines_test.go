package timelines

import (
	"github.com/gzcharleszhang/course-planner/internal/app/components/terms"
	"reflect"
	"testing"
	"time"

	"github.com/gzcharleszhang/course-planner/internal/app/components/courses"
	"github.com/gzcharleszhang/course-planner/internal/app/components/plans"
	"github.com/stretchr/testify/assert"
)

func TestTimeline_Aggregate(t *testing.T) {
	currTime := time.Now()
	timeline := Timeline{
		TermRecords: []*terms.TermRecord{
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
	assert.Equal(t, &expectedRecords, timeline.Aggregate())
}

func TestTimeline_IncompletePlans(t *testing.T) {
	currTime := time.Now()
	csDegree := plans.Degree{
		Name: "Easy BCS",
		Requirements: plans.DegreeRequirements(courses.CourseRequirementSet{
			MinCoursesToSatisfy: 2,
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
		}),
	}
	mathDegree := plans.Degree{
		Name: "Easy BMath",
		Requirements: plans.DegreeRequirements(courses.CourseRequirementSet{
			MinCoursesToSatisfy: 2,
			Requirements: courses.CourseRequirementRules{
				courses.CourseRequirement{
					MinGrade: 50,
					CourseId: 0,
				},
				courses.CourseRequirement{
					MinGrade: 60,
					CourseId: 2,
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
					CourseId: 3,
				},
				courses.CourseRequirement{
					MinGrade: 50,
					CourseId: 4,
				},
			},
		}),
	}

	csPlan, econMinorPlan, mathPlan := plans.Plan(csDegree), plans.Plan(econMinor), plans.Plan(mathDegree)
	var csEconPlan plans.Plans
	csEconPlan = append(csEconPlan, csPlan)
	csEconPlan = append(csEconPlan, econMinorPlan)
	var csMathPlan plans.Plans
	csMathPlan = append(csMathPlan, csPlan)
	csMathPlan = append(csMathPlan, mathPlan)

	// satisfied
	timeline := Timeline{
		TermRecords: []*terms.TermRecord{
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
	assert.Equal(t, 0, len(timeline.IncompletePlans()))

	// satisfied overlapping plans
	timeline = Timeline{
		TermRecords: []*terms.TermRecord{
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
	assert.Equal(t, 0, len(timeline.IncompletePlans()))

	// not satisfied
	timeline = Timeline{
		TermRecords: []*terms.TermRecord{
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
	inCompletePlans := timeline.IncompletePlans()
	assert.Equal(t, 1, len(inCompletePlans))
	assert.Equal(t, string(csDegree.Name), inCompletePlans[0].GetName())

	// grade requirement not met
	timeline = Timeline{
		TermRecords: []*terms.TermRecord{
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
	inCompletePlans = timeline.IncompletePlans()
	assert.Equal(t, 1, len(inCompletePlans))
	assert.Equal(t, string(csDegree.Name), inCompletePlans[0].GetName())

	// grade requirement met after repeating course
	timeline = Timeline{
		TermRecords: []*terms.TermRecord{
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
	inCompletePlans = timeline.IncompletePlans()
	assert.Equal(t, 0, len(inCompletePlans))
}

func TestTimeline_InvalidCourses(t *testing.T) {
	course2 := courses.CourseRecord{
		Course: courses.Course{
			Id: 1,
			Prereqs: courses.CourseRequirement{
				CourseId: 0,
				MinGrade: 60,
			},
		},
	}
	course3 := courses.CourseRecord{
		Course: courses.Course{
			Id: 2,
			Prereqs: courses.CourseRequirement{
				CourseId: 1,
				MinGrade: 60,
			},
		},
	}
	currTime := time.Now()
	type fields struct {
		Id          TimelineId
		Name        TimelineName
		TermRecords []*terms.TermRecord
		Plans       plans.Plans
	}
	tests := []struct {
		name   string
		fields fields
		want   courses.CourseRecords
	}{
		{
			name: "all valid",
			fields: fields{
				TermRecords: []*terms.TermRecord{
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
				TermRecords: []*terms.TermRecord{
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
			timeline := Timeline{
				Id:          tt.fields.Id,
				Name:        tt.fields.Name,
				TermRecords: tt.fields.TermRecords,
				Plans:       tt.fields.Plans,
			}
			if got := timeline.InvalidCourses(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Timeline.InvalidCourses() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeline_NewTimeline(t *testing.T) {
	currTime := time.Now()
	timelineName := TimelineName("timelineName")
	courseHistory := []*terms.TermRecord{
		{
			Term: terms.Term{
				Name:   "1A",
				Season: terms.TermSeason(9),
				Year:   2018,
			},
			Id: "#131ijj2",
			CourseRecords: courses.CourseRecords{
				&courses.CourseRecord{
					Course: courses.Course{
						Id: 0,
					},
					Grade:          85,
					CompletionDate: &currTime,
				},
			},
		},
		{
			Term: terms.Term{
				Name:   "1B",
				Season: terms.TermSeason(1),
				Year:   2019,
			},
			CourseRecords: courses.CourseRecords{
				&courses.CourseRecord{
					Course: courses.Course{
						Id: 1,
					},
					Grade:          70,
					CompletionDate: &currTime,
				},
				&courses.CourseRecord{
					Course: courses.Course{
						Id: 2,
					},
					Grade:          50,
					CompletionDate: &currTime,
				},
			},
		},
		{
			Term: terms.Term{
				Name:   "2A",
				Season: terms.TermSeason(5),
				Year:   2019,
			},
			CourseRecords: courses.CourseRecords{
				&courses.CourseRecord{
					Course: courses.Course{
						Id: 3,
					},
				},
			},
		},
	}
	newTimeline := NewTimeline(timelineName, courseHistory)
	courseHistory[1].Term.Season = terms.TermSeason(5)
	courseHistory[1].CourseRecords[1].Course.Id = 10
	courseHistory[1].CourseRecords[1].Grade = 100
	courseHistory[1].Term.Name = "3A"
	courseHistory[1].Term.Year = 3019

	// have to check that each individual one is not equals
	assert.Equal(t, courseHistory[0].Term, newTimeline)
	//assert.NotEqual(t, courseHistory[0].Id, newTimeline.TermRecords[0].Id)
	//assert.Equal(t, courseHistory[0].CourseRecords, newTimeline.TermRecords[0].CourseRecords)
	//assert.NotEqual(t, courseHistory[1].Term, newTimeline.TermRecords[1].Term)
	//assert.NotEqual(t, courseHistory[1].CourseRecords, newTimeline.TermRecords[1].CourseRecords)
	//assert.Equal(t, courseHistory[2].Term, newTimeline.TermRecords[2].Term)
	//assert.Equal(t, courseHistory[2].CourseRecords, newTimeline.TermRecords[2].CourseRecords)
}
