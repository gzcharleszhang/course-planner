package terms

import (
	"github.com/gzcharleszhang/course-planner/internal/app/components/utils"
	"reflect"
	"testing"
	"time"

	"github.com/gzcharleszhang/course-planner/internal/app/components/courses"
)

func TestTermRecord_InvalidCourses(t *testing.T) {
	course1 := courses.CourseRecord{
		Course: courses.Course{
			Id: 0,
			Prereqs: courses.CourseRequirementRules{
				courses.CourseRequirement{
					CourseId: 1,
					MinGrade: 60,
				},
			},
		},
	}
	course2 := courses.CourseRecord{
		Course: courses.Course{
			Id: 2,
			Prereqs: courses.CourseRequirementRules{
				courses.CourseRequirementSet{
					NumCoursesToSatisfy: 1,
					Requirements: courses.CourseRequirementRules{
						courses.CourseRequirement{
							CourseId: 3,
							MinGrade: 60,
						},
						courses.CourseRequirement{
							CourseId: 4,
							MinGrade: 60,
						},
					},
				},
			},
		},
	}
	currTime := time.Now()
	type fields struct {
		Term          Term
		CourseRecords courses.CourseRecords
	}
	type args struct {
		pastRecords courses.CourseRecords
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   courses.CourseRecords
	}{
		{
			name: "no invalid courses",
			fields: fields{
				CourseRecords: courses.CourseRecords{
					&course1,
				},
			},
			args: args{
				pastRecords: courses.CourseRecords{
					&courses.CourseRecord{
						Course: courses.Course{
							Id: 1,
						},
						Grade:          60,
						CompletionDate: &currTime,
					},
				},
			},
			want: courses.CourseRecords{},
		},
		{
			name: "one invalid course",
			fields: fields{
				CourseRecords: courses.CourseRecords{
					&course1,
					&course2,
				},
			},
			args: args{
				pastRecords: courses.CourseRecords{
					&courses.CourseRecord{
						Course: courses.Course{
							Id: 1,
						},
						Grade:          59,
						CompletionDate: &currTime,
					},
					&courses.CourseRecord{
						Course: courses.Course{
							Id: 4,
						},
						Grade:          60,
						CompletionDate: &currTime,
					},
				},
			},
			want: courses.CourseRecords{
				&course1,
			},
		},
		{
			name: "multiple invalid courses",
			fields: fields{
				CourseRecords: courses.CourseRecords{
					&course1,
					&course2,
				},
			},
			args: args{
				pastRecords: courses.CourseRecords{
					&courses.CourseRecord{
						Course: courses.Course{
							Id: 1,
						},
						Grade:          59,
						CompletionDate: &currTime,
					},
					&courses.CourseRecord{
						Course: courses.Course{
							Id: 4,
						},
						Grade:          59,
						CompletionDate: &currTime,
					},
				},
			},
			want: courses.CourseRecords{
				&course1,
				&course2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := TermRecord{
				Term:          tt.fields.Term,
				CourseRecords: tt.fields.CourseRecords,
			}
			if got := tr.InvalidCourses(tt.args.pastRecords); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TermRecord.InvalidCourses() = %v, want %v", utils.ToJson(got), utils.ToJson(tt.want))
			}
		})
	}
}
