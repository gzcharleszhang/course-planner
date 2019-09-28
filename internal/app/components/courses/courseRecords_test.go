// +build all unit

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
						Id: "0",
					},
					Grade:          50,
					CompletionDate: &currTime,
				},
				&CourseRecord{
					Course: Course{
						Id: "1",
					},
				},
				&CourseRecord{
					Course: Course{
						Id: "2",
					},
					Grade:          70,
					CompletionDate: &currTime,
				},
			},
			want: map[CourseId]*CourseRecord{
				CourseId("0"): {
					Course: Course{
						Id: "0",
					},
					Grade:          50,
					CompletionDate: &currTime,
				},
				CourseId("1"): {
					Course: Course{
						Id: "1",
					},
				},
				CourseId("2"): {
					Course: Course{
						Id: "2",
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
						Id: "0",
					},
				},
				&CourseRecord{
					Course: Course{
						Id: "0",
					},
					Grade:          50,
					CompletionDate: &currTime,
				},
				&CourseRecord{
					Course: Course{
						Id: "2",
					},
					Grade:          60,
					CompletionDate: &currTime,
				},
				&CourseRecord{
					Course: Course{
						Id: "1",
					},
				},
				&CourseRecord{
					Course: Course{
						Id: "2",
					},
					Grade:          90,
					CompletionDate: &currTime,
				},
				&CourseRecord{
					Course: Course{
						Id: "2",
					},
					Grade:          70,
					CompletionDate: &currTime,
				},
			},
			want: map[CourseId]*CourseRecord{
				CourseId("0"): {
					Course: Course{
						Id: "0",
					},
				},
				CourseId("1"): {
					Course: Course{
						Id: "1",
					},
				},
				CourseId("2"): {
					Course: Course{
						Id: "2",
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
						Id: "0",
					},
					Grade: 50,
				},
				&CourseRecord{
					Course: Course{
						Id: "1",
					},
					Grade: 70,
				},
			},
			args: args{
				records: CourseRecords{
					&CourseRecord{
						Course: Course{
							Id: "2",
						},
						Grade: 90,
					},
					&CourseRecord{
						Course: Course{
							Id: "1",
						},
						Grade: 60,
					},
				},
			},
			want: CourseRecords{
				&CourseRecord{
					Course: Course{
						Id: "0",
					},
					Grade: 50,
				},
				&CourseRecord{
					Course: Course{
						Id: "1",
					},
					Grade: 70,
				},
				&CourseRecord{
					Course: Course{
						Id: "2",
					},
					Grade: 90,
				},
				&CourseRecord{
					Course: Course{
						Id: "1",
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
							Id: "2",
						},
						Grade: 90,
					},
					&CourseRecord{
						Course: Course{
							Id: "1",
						},
						Grade: 60,
					},
				},
			},
			want: CourseRecords{
				&CourseRecord{
					Course: Course{
						Id: "2",
					},
					Grade: 90,
				},
				&CourseRecord{
					Course: Course{
						Id: "1",
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
						Id: "0",
					},
					Grade: 50,
				},
				&CourseRecord{
					Course: Course{
						Id: "1",
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
						Id: "0",
					},
					Grade: 50,
				},
				&CourseRecord{
					Course: Course{
						Id: "1",
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
						Id: "1",
					},
					Grade: 100,
				},
				&CourseRecord{
					Course: Course{
						Id: "0",
					},
					Grade: 50,
				},
				&CourseRecord{
					Course: Course{
						Id: "1",
					},
					Grade: 70,
				},
			},
			args: args{
				records: CourseRecords{
					&CourseRecord{
						Course: Course{
							Id: "2",
						},
						Grade: 90,
					},
					&CourseRecord{
						Course: Course{
							Id: "1",
						},
						Grade: 60,
					},
				},
			},
			want: CourseRecords{
				&CourseRecord{
					Course: Course{
						Id: "0",
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
							Id: "2",
						},
						Grade: 90,
					},
					&CourseRecord{
						Course: Course{
							Id: "1",
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
						Id: "0",
					},
					Grade: 50,
				},
				&CourseRecord{
					Course: Course{
						Id: "1",
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
						Id: "0",
					},
					Grade: 50,
				},
				&CourseRecord{
					Course: Course{
						Id: "1",
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
					Id: "0",
					Prereqs: CourseRequirementRules{
						CourseRequirement{
							CourseId: "1",
							MinGrade: 60,
						},
					},
				},
			},
			args: args{
				pastRecords: &CourseRecords{
					&CourseRecord{
						Course: Course{
							Id: "1",
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
					Id: "0",
					Prereqs: CourseRequirementRules{
						CourseRequirement{
							CourseId: "1",
							MinGrade: 60,
						},
					},
				},
			},
			args: args{
				pastRecords: &CourseRecords{
					&CourseRecord{
						Course: Course{
							Id: "1",
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
					Id: "0",
					Prereqs: CourseRequirementRules{
						CourseRequirement{
							CourseId: "1",
							MinGrade: 60,
						},
					},
				},
				Override: true,
			},
			args: args{
				pastRecords: &CourseRecords{
					&CourseRecord{
						Course: Course{
							Id: "1",
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

func TestCourseRecord_IsCompleted(t *testing.T) {
	currTime := time.Now()
	type fields struct {
		Course         Course
		Id             CourseRecordId
		Grade          CourseGrade
		CompletionDate *time.Time
		Override       bool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "completed without grade",
			fields: fields{
				Course: Course{
					Id: CourseId(0),
				},
				CompletionDate: &currTime,
			},
			want: true,
		},
		{
			name: "completed with grade",
			fields: fields{
				Course: Course{
					Id: CourseId(0),
				},
				Grade:          CourseGrade(10),
				CompletionDate: &currTime,
			},
			want: true,
		},
		{
			name: "incomplete with grade",
			fields: fields{
				Course: Course{
					Id: CourseId(0),
				},
				Grade: CourseGrade(100),
			},
			want: false,
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
			if got := cr.IsCompleted(); got != tt.want {
				t.Errorf("CourseRecord.IsCompleted() = %v, want %v", utils.ToJson(got), utils.ToJson(tt.want))
			}
		})
	}
}

func TestCourseRecords_CurrentCAV(t *testing.T) {
	currTime := time.Now()
	tests := []struct {
		name string
		cr   CourseRecords
		want CourseGrade
	}{
		{
			name: "all completed",
			cr: CourseRecords{
				{
					Grade:          CourseGrade(10),
					CompletionDate: &currTime,
				},
				{
					Grade:          CourseGrade(20),
					CompletionDate: &currTime,
				},
				{
					Grade:          CourseGrade(30),
					CompletionDate: &currTime,
				},
				{
					Grade:          CourseGrade(40),
					CompletionDate: &currTime,
				},
			},
			want: CourseGrade(25),
		},
		{
			name: "some incomplete",
			cr: CourseRecords{
				{
					Grade: CourseGrade(10),
				},
				{
					Grade:          CourseGrade(20),
					CompletionDate: &currTime,
				},
				{
					Grade: CourseGrade(30),
				},
				{
					Grade:          CourseGrade(40),
					CompletionDate: &currTime,
				},
			},
			want: CourseGrade(30),
		},
		{
			name: "all incomplete",
			cr: CourseRecords{
				{
					Grade: CourseGrade(10),
				},
				{
					Grade: CourseGrade(20),
				},
				{
					Grade: CourseGrade(30),
				},
				{
					Grade: CourseGrade(40),
				},
			},
			want: CourseGrade(0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cr.CurrentCAV(); got != tt.want {
				t.Errorf("CourseRecords.CurrentCAV() = %v, want %v", utils.ToJson(got), utils.ToJson(tt.want))
			}
		})
	}
}

func TestCourseRecords_Copy(t *testing.T) {
	currTime := time.Now()
	courseRecord1 := CourseRecord{
		Course: Course{
			Id: "3",
		},
		Id:             CourseRecordId("random"),
		Grade:          50,
		CompletionDate: &currTime,
	}
	courseRecord1Expect := CourseRecord{
		Course: Course{
			Id: "3",
		},
		Id:             CourseRecordId("asdfghjkl"),
		Grade:          50,
		CompletionDate: &currTime,
	}
	courseRecord2 := CourseRecord{
		Course: Course{
			Id: "2",
		},
		Id:             CourseRecordId("asdfghjkl"),
		Grade:          50,
		CompletionDate: &currTime,
	}
	courseRecord3 := CourseRecord{
		Course: Course{
			Id: "1",
		},
		Id:             CourseRecordId("asdfghjkl"),
		Grade:          70,
		CompletionDate: &currTime,
	}
	tests := []struct {
		name    string
		records CourseRecords
		want    CourseRecords
	}{
		{
			name:    "empty",
			records: CourseRecords{},
			want:    nil,
		},
		{
			name: "single",
			records: CourseRecords{
				&courseRecord1,
			},
			want: CourseRecords{
				&courseRecord1Expect,
			},
		},
		{
			name: "mutli",
			records: CourseRecords{
				&courseRecord1,
				&courseRecord2,
				&courseRecord3,
			},
			want: CourseRecords{
				&courseRecord1Expect,
				&courseRecord2,
				&courseRecord3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.records.Copy()
			for i := range got {
				if got[i].Id == tt.want[i].Id {
					t.Errorf("Expected different id for recieved %v and expected %v", got[i].Id, tt.want[i].Id)
				}
				got[i].Id = CourseRecordId("asdfghjkl")
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CourseRecords.Copy() = %v, want %v", utils.ToJson(got), utils.ToJson(tt.want))
			}
		})
	}
}

func TestCourseRecord_Copy(t *testing.T) {
	currTime := time.Now()
	courseRecord1 := CourseRecord{
		Course: Course{
			Id: "3",
		},
		Id:             CourseRecordId("random"),
		Grade:          50,
		CompletionDate: &currTime,
	}
	courseRecord1Expect := CourseRecord{
		Course: Course{
			Id: "3",
		},
		Id:             CourseRecordId("asdfghjkl"),
		Grade:          50,
		CompletionDate: &currTime,
	}
	courseRecord2 := CourseRecord{
		Course: Course{
			Id: "2",
		},
		Id:             CourseRecordId("asdfghjkl"),
		Grade:          50,
		CompletionDate: &currTime,
	}
	type fields struct {
		Course         Course
		Id             CourseRecordId
		Grade          CourseGrade
		CompletionDate *time.Time
		Override       bool
	}
	tests := []struct {
		name   string
		fields CourseRecord
		want   CourseRecord
	}{
		{
			name:   "empty",
			fields: CourseRecord{},
			want: CourseRecord{
				Id: CourseRecordId("asdfghjkl"),
			},
		},
		{
			name:   "single1",
			fields: courseRecord1,
			want:   courseRecord1Expect,
		},
		{
			name:   "single2",
			fields: courseRecord2,
			want:   courseRecord2,
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
			got := cr.Copy()
			if got.Id == cr.Id {
				t.Errorf("Expected different id for recieved %v and expected %v", got.Id, cr.Id)
			}
			got.Id = CourseRecordId("asdfghjkl")
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CourseRecord.Copy() = %v, want %v", got, tt.want)
			}
		})
	}
}
