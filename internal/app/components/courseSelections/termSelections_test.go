package courseSelections

import (
	"github.com/gzcharleszhang/course-planner/internal/app/components/utils"
	"reflect"
	"testing"
	"time"

	"github.com/gzcharleszhang/course-planner/internal/app/components/courses"
	"github.com/gzcharleszhang/course-planner/internal/app/components/terms"
)

func TestTermSelection_InvalidCourses(t *testing.T) {
	course1 := courses.CourseRecord{
		Course: courses.Course{
			Id: 0,
			Prereqs: courses.CourseRequirement{
				Course: &courses.Course{
					Id: 1,
				},
				MinGrade: 60,
			},
		},
	}
	course2 := courses.CourseRecord{
		Course: courses.Course{
			Id: 2,
			Prereqs: courses.CourseRequirementSet{
				MinCoursesToSatisfy: 1,
				Requirements: courses.CourseRequirementRules{
					courses.CourseRequirement{
						Course: &courses.Course{
							Id: 3,
						},
						MinGrade: 60,
					},
					courses.CourseRequirement{
						Course: &courses.Course{
							Id: 4,
						},
						MinGrade: 60,
					},
				},
			},
		},
	}
	currTime := time.Now()
	type fields struct {
		Term          terms.Term
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
			ts := TermSelection{
				Term:          tt.fields.Term,
				CourseRecords: tt.fields.CourseRecords,
			}
			if got := ts.InvalidCourses(tt.args.pastRecords); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TermSelection.InvalidCourses() = %v, want %v", utils.ToJson(got), utils.ToJson(tt.want))
			}
		})
	}
}

func Test_isPrereqSatisfied(t *testing.T) {
	currTime := time.Now()
	type args struct {
		record      *courses.CourseRecord
		pastRecords *courses.CourseRecords
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "satisfied",
			args: args{
				record: &courses.CourseRecord{
					Course: courses.Course{
						Id: 0,
						Prereqs: courses.CourseRequirement{
							Course: &courses.Course{
								Id: 1,
							},
							MinGrade: 60,
						},
					},
				},
				pastRecords: &courses.CourseRecords{
					&courses.CourseRecord{
						Course: courses.Course{
							Id: 1,
						},
						Grade: 10, // should still be satisfied because it's a future record
					},
				},
			},
			want: true,
		},
		{
			name: "not satisfied",
			args: args{
				record: &courses.CourseRecord{
					Course: courses.Course{
						Id: 0,
						Prereqs: courses.CourseRequirement{
							Course: &courses.Course{
								Id: 1,
							},
							MinGrade: 60,
						},
					},
				},
				pastRecords: &courses.CourseRecords{
					&courses.CourseRecord{
						Course: courses.Course{
							Id: 1,
						},
						Grade:          10,
						CompletionDate: &currTime,
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPrereqSatisfied(tt.args.record, tt.args.pastRecords); got != tt.want {
				t.Errorf("isPrereqSatisfied() = %v, want %v", got, tt.want)
			}
		})
	}
}
