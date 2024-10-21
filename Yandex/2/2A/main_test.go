package main

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	type args struct {
		matrix [][]int
		rows   int
		col    int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "1",
			args: args{
				matrix: [][]int{
					{1, 2, 3},
					{0, 2, 6},
					{7, 4, 1},
					{2, 7, 0},
				},
				rows: 4,
				col:  3,
			},
			want: [][]int{
				{1, 0, 7, 2},
				{2, 2, 4, 7},
				{3, 6, 1, 0},
			},
		},
		{
			name: "2",
			args: args{
				matrix: [][]int{
					{1, 2, 3},
				},
				rows: 4,
				col:  3,
			},
			want: [][]int{
				{1},
				{2},
				{3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.matrix, tt.args.rows, tt.args.col); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
