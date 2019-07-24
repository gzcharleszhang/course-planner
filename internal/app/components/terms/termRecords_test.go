package terms

import (
	"reflect"
	"testing"
	"time"

	"github.com/gzcharleszhang/course-planner/internal/app/components/courses"
	"github.com/gzcharleszhang/course-planner/internal/app/components/utils"
	"github.com/stretchr/testify/assert"
)

func TestTermRecord_InvalidCourses(t *testing.T) {
	course1 := courses.CourseRecord{
		Course: courses.Course{
			Id: 0,
			Prereqs: courses.CourseRequirement{
				CourseId: 1,
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
							CourseId: 1,
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
							CourseId: 1,
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

func X(t *testing.T) {
	currTime := time.Now()
	termRecords := []*TermRecord{
		{
			Term: Term{
				Name:   "1A",
				Id:     1189,
				Season: TermSeason(9),
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
			Term: Term{
				Name:   "1B",
				Id:     1191,
				Season: TermSeason(1),
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
			Term: Term{
				Name:   "2A",
				Id:     1195,
				Season: TermSeason(5),
				Year:   2019,
			},
			CourseRecords: courses.CourseRecords{
				&courses.CourseRecord{
					Course: courses.Course{
						Id: 3,
					},
					Grade:          50,
					CompletionDate: &currTime,
				},
			},
		},
	}

	newTermRecord := CopyRecords(termRecords)

	termRecords[1].Term.Season = TermSeason(5)
	termRecords[1].CourseRecords[1].Course.Id = 10
	termRecords[1].CourseRecords[1].Grade = 100
	termRecords[1].Term.Name = "3A"
	termRecords[1].Term.Year = 3019

	assert.NotEqual(t, termRecords[0].Id, newTermRecord[0].Id)
	assert.Equal(t, termRecords[0].Term, newTermRecord[0].Term)
	assert.NotEqual(t, termRecords[1].Term, newTermRecord[1].Term)

}

func TestTermRecord_CopyTermRecords(t *testing.T) {
	currTime := time.Now()
	type args struct {
		records []*TermRecord
	}
	tests := []struct {
		name string
		args args
		want []*TermRecord
	}{
		// TODO: Add test cases.
		{
			name: "empty",
			args: args{
				[]*TermRecord{},
			},
			want: []*TermRecord{},
		},
		{
			name: "single",
			args: args{
				 []*TermRecord{
					{
						Term: Term{
							Name:   "1A",
							Id:     1189,
							Season: TermSeason(9),
							Year:   2018,
						},
						Id: "131ijj2",
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
				},
			},
			want: []*TermRecord{
				{
					Term: Term{
						Name:   "1A",
						Id:     1189,
						Season: TermSeason(9),
						Year:   2018,
					},
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
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CopyRecords(tt.args.records); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CopyRecords() = %v, want %v", got, tt.want)
			}
		})
	}
}
