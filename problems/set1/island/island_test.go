package main

import "testing"

func Test_numIslands(t *testing.T) {
	type args struct {
		grid [][]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name:"test-case#1",
			args:args{
				grid:nil,
			},
			want:0,

		},

		{
			name:"test-case#2",
			args:args{
				grid:[][]byte{
					{48},
				},
			},
			want:0,

		},

		{
			name:"test-case#2",
			args:args{
				grid:[][]byte{
					{49},
				},
			},
			want:1,

		},

		{
			name:"test-case#3",
			args:args{
				grid:[][]byte{
					{49},
				},
			},
			want:1,

		},

		{
			name:"test-case#4",
			args:args{
				grid:[][]byte{
					{49,49,48,48,48},
					{49,49,48,48,48},
					{48,48,49,48,48},
					{48,48,48,49,49},
				},
			},
			want:3,
		},

		{
			name:"test-case#4",
			args:args{
				grid:[][]byte{
					{49,49,49,49,48},
					{49,49,48,49,48},
					{49,49,48,48,48},
					{48,48,48,48,48},
				},
			},
			want:1,
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numIslands(tt.args.grid); got != tt.want {
				t.Errorf("numIslands() = %v, want %v", got, tt.want)
			}
		})
	}
}
