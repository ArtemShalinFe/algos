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
		k      int
		m      int
	}
	tests := []struct {
		name string
		args args
		want []int
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
				k:    3,
				m:    0,
			},
			want: []int{7, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.matrix, tt.args.rows, tt.args.col, tt.args.k, tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
