package main

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	type args struct {
		field [][]int
		k     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				field: [][]int{
					{1, 2, 3, 1},
					{2, 0, 0, 2},
					{2, 0, 0, 2},
					{2, 0, 0, 2},
				},
				k: 3,
			},
			want: 2,
		},
		{
			name: "2",
			args: args{
				field: [][]int{
					{1, 1, 1, 1},
					{9, 9, 9, 9},
					{1, 1, 1, 1},
					{9, 9, 1, 1},
				},
				k: 4,
			},
			want: 1,
		},
		{
			name: "3",
			args: args{
				field: [][]int{
					{1, 1, 1, 1},
					{1, 1, 1, 1},
					{1, 1, 1, 1},
					{1, 1, 1, 1},
				},
				k: 4,
			},
			want: 0,
		},
		{
			name: "4",
			args: args{
				field: [][]int{
					{1, 1, 1, 1},
					{1, 1, 1, 1},
					{0, 0, 0, 0},
					{1, 1, 0, 0},
				},
				k: 5,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.k, tt.args.field); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
