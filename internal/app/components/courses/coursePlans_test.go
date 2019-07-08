package courses

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCoursePlan_Aggregate(t *testing.T) {
	currTime := time.Now()
	coursePlan := CoursePlan{
		TermPlans: []*TermPlan{
			{
				CourseRecords: CourseRecords{
					CourseId(0): &CourseRecord{
						Course: Course{
							Id: 0,
						},
						Grade:          50,
						CompletionDate: &currTime,
					},
				},
			},
			{
				CourseRecords: CourseRecords{
					CourseId(1): &CourseRecord{
						Course: Course{
							Id: 1,
						},
						Grade:          60,
						CompletionDate: &currTime,
					},
					CourseId(2): &CourseRecord{
						Course: Course{
							Id: 2,
						},
						Grade:          60,
						CompletionDate: &currTime,
					},
				},
			},
			{
				CourseRecords: CourseRecords{
					CourseId(3): &CourseRecord{
						Course: Course{
							Id: 3,
						},
					},
				},
			},
		},
	}
	expectedRecords := CourseRecords{
		CourseId(0): &CourseRecord{
			Course: Course{
				Id: 0,
			},
			Grade:          50,
			CompletionDate: &currTime,
		},
		CourseId(1): &CourseRecord{
			Course: Course{
				Id: 1,
			},
			Grade:          60,
			CompletionDate: &currTime,
		},
		CourseId(2): &CourseRecord{
			Course: Course{
				Id: 2,
			},
			Grade:          60,
			CompletionDate: &currTime,
		},
		CourseId(3): &CourseRecord{
			Course: Course{
				Id: 3,
			},
		},
	}
	assert.Equal(t, &expectedRecords, coursePlan.Aggregate())
}
