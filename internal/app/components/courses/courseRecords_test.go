package courses

import (
	"reflect"
	"testing"
	"time"

	"github.com/gzcharleszhang/course-planner/internal/app/components/utils"
)

func TestCourseRecords_ToCourseIdMap(t *testing.T) {
	currTime := time.Now()
	tests := []struct {
		name string
		cr   CourseRecords
		want map[CourseId]*CourseRecord
	}{
		{
			name: "all unique",
			cr: CourseRecords{
				&CourseRecord{
					Course: Course{
						Id: 0,
					},
					Grade:          50,
					CompletionDate: &currTime,
				},
				&CourseRecord{
					Course: Course{
						Id: 1,
					},
				},
				&CourseRecord{
					Course: Course{
						Id: 2,
					},
					Grade:          70,
					CompletionDate: &currTime,
				},
			},
			want: map[CourseId]*CourseRecord{
				CourseId(0): {
					Course: Course{
						Id: 0,
					},
					Grade:          50,
					CompletionDate: &currTime,
				},
				CourseId(1): {
					Course: Course{
						Id: 1,
					},
				},
				CourseId(2): {
					Course: Course{
						Id: 2,
					},
					Grade:          70,
					CompletionDate: &currTime,
				},
			},
		},
		{
			name: "duplicate courses",
			cr: CourseRecords{
				&CourseRecord{
					Course: Course{
						Id: 0,
					},
				},
				&CourseRecord{
					Course: Course{
						Id: 0,
					},
					Grade:          50,
					CompletionDate: &currTime,
				},
				&CourseRecord{
					Course: Course{
						Id: 2,
					},
					Grade:          60,
					CompletionDate: &currTime,
				},
				&CourseRecord{
					Course: Course{
						Id: 1,
					},
				},
				&CourseRecord{
					Course: Course{
						Id: 2,
					},
					Grade:          90,
					CompletionDate: &currTime,
				},
				&CourseRecord{
					Course: Course{
						Id: 2,
					},
					Grade:          70,
					CompletionDate: &currTime,
				},
			},
			want: map[CourseId]*CourseRecord{
				CourseId(0): {
					Course: Course{
						Id: 0,
					},
				},
				CourseId(1): {
					Course: Course{
						Id: 1,
					},
				},
				CourseId(2): {
					Course: Course{
						Id: 2,
					},
					Grade:          90,
					CompletionDate: &currTime,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cr.ToCourseIdMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CourseRecords.ToCourseIdMap() = %v, want %v", utils.ToJson(got), utils.ToJson(tt.want))
			}
		})
	}
}

func TestCourseRecords_Merge(t *testing.T) {
	type args struct {
		records CourseRecords
	}
	tests := []struct {
		name string
		cr   CourseRecords
		args args
		want CourseRecords
	}{
		{
			name: "both have values",
			cr: CourseRecords{
				&CourseRecord{
					Course: Course{
						Id: 0,
					},
					Grade: 50,
				},
				&CourseRecord{
					Course: Course{
						Id: 1,
					},
					Grade: 70,
				},
			},
			args: args{
				records: CourseRecords{
					&CourseRecord{
						Course: Course{
							Id: 2,
						},
						Grade: 90,
					},
					&CourseRecord{
						Course: Course{
							Id: 1,
						},
						Grade: 60,
					},
				},
			},
			want: CourseRecords{
				&CourseRecord{
					Course: Course{
						Id: 0,
					},
					Grade: 50,
				},
				&CourseRecord{
					Course: Course{
						Id: 1,
					},
					Grade: 70,
				},
				&CourseRecord{
					Course: Course{
						Id: 2,
					},
					Grade: 90,
				},
				&CourseRecord{
					Course: Course{
						Id: 1,
					},
					Grade: 60,
				},
			},
		},
		{
			name: "empty first",
			cr:   CourseRecords{},
			args: args{
				records: CourseRecords{
					&CourseRecord{
						Course: Course{
							Id: 2,
						},
						Grade: 90,
					},
					&CourseRecord{
						Course: Course{
							Id: 1,
						},
						Grade: 60,
					},
				},
			},
			want: CourseRecords{
				&CourseRecord{
					Course: Course{
						Id: 2,
					},
					Grade: 90,
				},
				&CourseRecord{
					Course: Course{
						Id: 1,
					},
					Grade: 60,
				},
			},
		},
		{
			name: "empty second",
			cr: CourseRecords{
				&CourseRecord{
					Course: Course{
						Id: 0,
					},
					Grade: 50,
				},
				&CourseRecord{
					Course: Course{
						Id: 1,
					},
					Grade: 70,
				},
			},
			args: args{
				records: CourseRecords{},
			},
			want: CourseRecords{
				&CourseRecord{
					Course: Course{
						Id: 0,
					},
					Grade: 50,
				},
				&CourseRecord{
					Course: Course{
						Id: 1,
					},
					Grade: 70,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cr.Merge(tt.args.records); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CourseRecords.Merge() = %v, want %v", utils.ToJson(got), utils.ToJson(tt.want))
			}
		})
	}
}

func TestCourseRecords_Exclude(t *testing.T) {
	type args struct {
		records CourseRecords
	}
	tests := []struct {
		name string
		cr   CourseRecords
		args args
		want CourseRecords
	}{
		{
			name: "both have values",
			cr: CourseRecords{
				&CourseRecord{
					Course: Course{
						Id: 1,
					},
					Grade: 100,
				},
				&CourseRecord{
					Course: Course{
						Id: 0,
					},
					Grade: 50,
				},
				&CourseRecord{
					Course: Course{
						Id: 1,
					},
					Grade: 70,
				},
			},
			args: args{
				records: CourseRecords{
					&CourseRecord{
						Course: Course{
							Id: 2,
						},
						Grade: 90,
					},
					&CourseRecord{
						Course: Course{
							Id: 1,
						},
						Grade: 60,
					},
				},
			},
			want: CourseRecords{
				&CourseRecord{
					Course: Course{
						Id: 0,
					},
					Grade: 50,
				},
			},
		},
		{
			name: "empty first",
			cr:   CourseRecords{},
			args: args{
				records: CourseRecords{
					&CourseRecord{
						Course: Course{
							Id: 2,
						},
						Grade: 90,
					},
					&CourseRecord{
						Course: Course{
							Id: 1,
						},
						Grade: 60,
					},
				},
			},
			want: CourseRecords{},
		},
		{
			name: "empty second",
			cr: CourseRecords{
				&CourseRecord{
					Course: Course{
						Id: 0,
					},
					Grade: 50,
				},
				&CourseRecord{
					Course: Course{
						Id: 1,
					},
					Grade: 70,
				},
			},
			args: args{
				records: CourseRecords{},
			},
			want: CourseRecords{
				&CourseRecord{
					Course: Course{
						Id: 0,
					},
					Grade: 50,
				},
				&CourseRecord{
					Course: Course{
						Id: 1,
					},
					Grade: 70,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cr.Exclude(tt.args.records); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CourseRecords.Exclude() = %v, want %v", utils.ToJson(got), utils.ToJson(tt.want))
			}
		})
	}
}

func TestCourseRecord_IsPrereqSatisfied(t *testing.T) {
	currTime := time.Now()
	type fields struct {
		Course         Course
		Id             CourseRecordId
		Grade          CourseGrade
		CompletionDate *time.Time
		Override       bool
	}
	type args struct {
		pastRecords *CourseRecords
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "satisfied",
			fields: fields{
				Course: Course{
					Id: 0,
					Prereqs: CourseRequirement{
						CourseId: 1,
						MinGrade: 60,
					},
				},
			},
			args: args{
				pastRecords: &CourseRecords{
					&CourseRecord{
						Course: Course{
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
			fields: fields{
				Course: Course{
					Id: 0,
					Prereqs: CourseRequirement{
						CourseId: 1,
						MinGrade: 60,
					},
				},
			},
			args: args{
				pastRecords: &CourseRecords{
					&CourseRecord{
						Course: Course{
							Id: 1,
						},
						Grade:          10,
						CompletionDate: &currTime,
					},
				},
			},
			want: false,
		},
		{
			name: "overridden",
			fields: fields{
				Course: Course{
					Id: 0,
					Prereqs: CourseRequirement{
						CourseId: 1,
						MinGrade: 60,
					},
				},
				Override: true,
			},
			args: args{
				pastRecords: &CourseRecords{
					&CourseRecord{
						Course: Course{
							Id: 1,
						},
						Grade:          10,
						CompletionDate: &currTime,
					},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cr := CourseRecord{
				Course:         tt.fields.Course,
				Id:             tt.fields.Id,
				Grade:          tt.fields.Grade,
				CompletionDate: tt.fields.CompletionDate,
				Override:       tt.fields.Override,
			}
			if got := cr.IsPrereqSatisfied(tt.args.pastRecords); got != tt.want {
				t.Errorf("CourseRecord.IsPrereqSatisfied() = %v, want %v", utils.ToJson(got), utils.ToJson(tt.want))
			}
		})
	}
}
