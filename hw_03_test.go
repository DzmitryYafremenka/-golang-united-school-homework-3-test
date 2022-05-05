package homework

import (
	"reflect"
	"testing"
)

func Test_average(t *testing.T) {
	type args struct {
		input [15]float32
	}
	tests := []struct {
		name       string
		args       args
		wantResult float32
	}{
		{
			name: "positive_case_01",
			args: args{
				input: [15]float32{3, 8, 16, 4, 5, 6, 5, 134, 2, 5, 6, 7, 12, 34, 2},
			},
			wantResult: 16.6,
		},
		{
			name: "positive_case_02",
			args: args{
				input: [15]float32{4, 12, 16, 4, 5, 6, 5, 444, 4, 5, 6, 12, 12, 34, 2},
			},
			wantResult: 38.066666,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := average(tt.args.input); gotResult != tt.wantResult {
				t.Errorf("average() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func Test_reverse(t *testing.T) {
	type args struct {
		input []int64
	}
	tests := []struct {
		name       string
		args       args
		wantResult []int64
	}{
		{
			name: "positive_case_01",
			args: args{
				input: []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			},
			wantResult: []int64{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		},
		{
			name: "positive_case_02",
			args: args{
				input: []int64{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5},
			},
			wantResult: []int64{5, 3, 5, 6, 2, 9, 5, 1, 4, 1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := reverse(tt.args.input); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("reverse() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func Test_sortMapValues(t *testing.T) {
	type args struct {
		input map[int]string
	}
	tests := []struct {
		name       string
		args       args
		wantResult []string
	}{
		{
			name: "positive_case_01",
			args: args{
				input: map[int]string{
					0:  "M",
					1:  "A",
					2:  "Y",
					3:  " ",
					4:  "T",
					5:  "H",
					6:  "E",
					7:  " ",
					8:  "F",
					9:  "O",
					10: "R",
					11: "C",
					12: "E",
					13: " ",
					14: "B",
					15: "E",
					16: " ",
					17: "W",
					18: "I",
					19: "T",
					20: "H",
					21: " ",
					22: "Y",
					23: "O",
					24: "U",
				},
			},
			wantResult: []string{"M", "A", "Y", " ", "T", "H", "E", " ", "F", "O", "R", "C", "E", " ", "B", "E", " ", "W", "I", "T", "H", " ", "Y", "O", "U"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := sortMapValues(tt.args.input); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("sortMapValues() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
