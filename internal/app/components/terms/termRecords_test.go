// +build all unit

package terms

import (
	"reflect"
	"testing"
	"time"

	"github.com/gzcharleszhang/course-planner/internal/app/components/courses"
	"github.com/gzcharleszhang/course-planner/internal/app/components/utils"
	_ "github.com/stretchr/testify/assert"
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

func TestTermRecords_Copy(t *testing.T) {
	currTime := time.Now()
	termRecord1 := TermRecord{
		Term: Term{
			Name:   TermName("1A"),
			Id:     TermId(1189),
			Season: TermSeason(9),
			Year:   TermYear(2018),
		},
		Id: TermRecordId("cfvgbhnj"),
		CourseRecords: courses.CourseRecords{
			&courses.CourseRecord{
				Course: courses.Course{
					Id: 0,
				},
				Id:             courses.CourseRecordId("xdcfvgbhnj"),
				Grade:          courses.CourseGrade(85),
				CompletionDate: &currTime,
				Override:       false,
			},
		},
	}
	termRecord1Expect := TermRecord{
		Term: Term{
			Name:   TermName("1A"),
			Id:     TermId(1189),
			Season: TermSeason(9),
			Year:   TermYear(2018),
		},
		Id: TermRecordId("asdfghjkl"),
		CourseRecords: courses.CourseRecords{
			&courses.CourseRecord{
				Course: courses.Course{
					Id: 0,
				},
				Id:             courses.CourseRecordId("asdfghjkl"),
				Grade:          courses.CourseGrade(85),
				CompletionDate: &currTime,
				Override:       false,
			},
		},
	}
	termRecord2 := TermRecord{
		Term: Term{
			Name:   TermName("3A"),
			Id:     TermId(1205),
			Season: TermSeason(5),
			Year:   TermYear(2020),
		},
		Id: TermRecordId("asdfghjkl"),
		CourseRecords: courses.CourseRecords{
			&courses.CourseRecord{
				Course: courses.Course{
					Id: 0,
				},
				Id:             courses.CourseRecordId("asdfghjkl"),
				Grade:          courses.CourseGrade(99),
				CompletionDate: &currTime,
				Override:       false,
			},
			&courses.CourseRecord{
				Course: courses.Course{
					Id: 1,
				},
				Id:             courses.CourseRecordId("asdfghjkl"),
				Grade:          courses.CourseGrade(31),
				CompletionDate: &currTime,
				Override:       true,
			},
		},
	}
	termRecord3 := TermRecord{
		Term: Term{
			Name:   TermName("2B"),
			Id:     TermId(1195),
			Season: TermSeason(5),
			Year:   TermYear(2019),
		},
		Id: TermRecordId("asdfghjkl"),
		CourseRecords: courses.CourseRecords{
			&courses.CourseRecord{
				Course: courses.Course{
					Id: 30,
				},
				Id:             courses.CourseRecordId("asdfghjkl"),
				CompletionDate: &currTime,
				Override:       true,
			},
			&courses.CourseRecord{
				Course: courses.Course{
					Id: 20,
				},
				Id:             courses.CourseRecordId("asdfghjkl"),
				Grade:          courses.CourseGrade(15),
				CompletionDate: &currTime,
				Override:       false,
			},
			&courses.CourseRecord{
				Course: courses.Course{
					Id: 10,
				},
				Id:             courses.CourseRecordId("asdfghjkl"),
				Grade:          courses.CourseGrade(22),
				CompletionDate: &currTime,
				Override:       false,
			},
		},
	}

	type args struct {
		records TermRecords
	}
	tests := []struct {
		name string
		args args
		want TermRecords
	}{
		{
			name: "empty",
			args: args{
				TermRecords{},
			},
			want: nil,
		},
		{
			name: "single + different course id",
			args: args{
				TermRecords{
					&termRecord1,
				},
			},
			want: TermRecords{
				&termRecord1Expect,
			},
		},
		{
			name: "multiple",
			args: args{
				TermRecords{
					&termRecord1,
					&termRecord2,
					&termRecord3,
				},
			},
			want: TermRecords{
				&termRecord1Expect,
				&termRecord2,
				&termRecord3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.records.Copy()
			if len(got) != len(tt.want) {
				t.Errorf("Expected length of %v, received %v", len(tt.want), len(got))
			}
			for i, tr := range got {
				if got[i].Id == tt.want[i].Id {
					t.Errorf("Expected record id %v, received %v", tt.want[i].Id, got[i].Id)
				}
				got[i].Id = TermRecordId("asdfghjkl")
				for j := range tr.CourseRecords {
					if got[i].CourseRecords[j].Id == tt.want[i].CourseRecords[j].Id {
						t.Errorf("Expected course record id %v, received %v", tt.want[i].CourseRecords[j].Id, got[i].CourseRecords[j].Id)
					}
					got[i].CourseRecords[j].Id = courses.CourseRecordId("asdfghjkl")
				}
				if !reflect.DeepEqual(got[i], tt.want[i]) {
					t.Errorf("CopyRecords() = %v, want %v", utils.ToJson(got[i]), utils.ToJson(tt.want[i]))
				}
			}
		})
	}
}

func TestTermRecord_Copy(t *testing.T) {
	currTime := time.Now()
	termRecord1 := TermRecord{
		Term: Term{
			Name:   "1B",
			Id:     1191,
			Season: TermSeason(1),
			Year:   2019,
		},
		Id: "generic",
		CourseRecords: courses.CourseRecords{
			&courses.CourseRecord{
				Course: courses.Course{
					Id: 1,
				},
				Id:             courses.CourseRecordId("different"),
				Grade:          70,
				CompletionDate: &currTime,
			},
		},
	}
	termRecord1Expect := TermRecord{
		Term: Term{
			Name:   "1B",
			Id:     1191,
			Season: TermSeason(1),
			Year:   2019,
		},
		Id: "generic",
		CourseRecords: courses.CourseRecords{
			&courses.CourseRecord{
				Course: courses.Course{
					Id: 1,
				},
				Id:             courses.CourseRecordId("asdfghjkl"),
				Grade:          70,
				CompletionDate: &currTime,
			},
		},
	}
	termRecord2 := TermRecord{
		Term: Term{
			Name:   "1A",
			Id:     1189,
			Season: TermSeason(9),
			Year:   2018,
		},
		Id: "generic",
			CourseRecords: courses.CourseRecords{
			&courses.CourseRecord{
				Course: courses.Course{
					Id: 0,
				},
				Id:             courses.CourseRecordId("asdfghjkl"),
				Grade:          85,
				CompletionDate: &currTime,
			},
			&courses.CourseRecord{
				Course: courses.Course{
					Id: 3,
				},
				Id:             courses.CourseRecordId("asdfghjkl"),
				Grade:          50,
				CompletionDate: &currTime,
			},
			&courses.CourseRecord{
				Course: courses.Course{
					Id: 2,
				},
				Id:             courses.CourseRecordId("asdfghjkl"),
				Grade:          50,
				CompletionDate: &currTime,
			},
		},
	}
	type fields struct {
		Term          Term
		Id            TermRecordId
		CourseRecords courses.CourseRecords
	}
	tests := []struct {
		name   string
		fields fields
		want   TermRecord
	}{
		{
			name: "empty",
			fields: fields{},
			want: TermRecord{
				Term: Term{
					Year: TermYear(1900), // by default
				},
				Id: "generic",
			},
		},
		{
			name: "justId",
			fields: fields{
				Id: "generic",
			},
			want: TermRecord{
				Term: Term{
					Year: TermYear(1900), // by default
				},
				Id: "generic",
			},
		},
		{
			name: 	"singleCourse",
			fields: fields(termRecord1),
			want: 	termRecord1Expect,
		},
		{
			name: 	"multiCourse",
			fields: fields(termRecord2),
			want:	termRecord2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := TermRecord{
				Term:          tt.fields.Term,
				Id:            tt.fields.Id,
				CourseRecords: tt.fields.CourseRecords,
			}
			got := tr.Copy()
			if got.Id == tr.Id {
				t.Errorf("Expected different id for recieved %v and expected %v", got.Id, tr.Id)
			}
			got.Id = "generic"
			for j := range got.CourseRecords {
				if got.CourseRecords[j].Id == tt.want.CourseRecords[j].Id {
					t.Errorf("Expected different course record id %v, received %v",
										tt.want.CourseRecords[j].Id, got.CourseRecords[j].Id)
				}
				got.CourseRecords[j].Id = courses.CourseRecordId("asdfghjkl")
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TermRecord.Copy() = %v, want %v", utils.ToJson(got), utils.ToJson(tt.want))
			}
		})
	}
}
