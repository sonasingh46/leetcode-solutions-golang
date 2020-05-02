package main

import "testing"

func Test_searchMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test-case#1",
			args: args{
				matrix: [][]int{},
				target: 0,
			},
			want: false,
		},
		{
			name: "test-case#2",
			args: args{
				matrix: nil,
				target: 0,
			},
			want: false,
		},
		{
			name: "test-case#3",
			args: args{
				matrix: [][]int{{}},
				target: 0,
			},
			want: false,
		},
		{
			name: "test-case#4",
			args: args{
				matrix: [][]int{{}, {}},
				target: 0,
			},
			want: false,
		},
		{
			name: "test-case#5",
			args: args{
				matrix: [][]int{{1}},
				target: 0,
			},
			want: false,
		},
		{
			name: "test-case#5",
			args: args{
				matrix: [][]int{{1}},
				target: 1,
			},
			want: true,
		},
		{
			name: "test-case#5",
			args: args{
				matrix: [][]int{{1}},
				target: 2,
			},
			want: false,
		},
		{
			name: "test-case#5",
			args: args{
				matrix: [][]int{
					{1, 3, 5, 7},
					{10, 12, 18, 26},
					{29, 32, 67, 78},
					{80, 89, 100, 133},
				},
				target: 3,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchMatrix(tt.args.matrix, tt.args.target); got != tt.want {
				t.Errorf("searchMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
