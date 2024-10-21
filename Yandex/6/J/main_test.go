package main

import (
	"reflect"
	"testing"
)

func Test_solution(t *testing.T) {
	type args struct {
		s     int
		m     int
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "1",
			args: args{
				s: 1,
				m: 3,
				edges: [][]int{
					{1, 2},
					{2, 3},
				},
			},
			want: [][]int{
				{0, 5},
				{1, 4},
				{2, 3},
			},
		},
		{
			name: "2",
			args: args{
				s: 1,
				m: 6,
				edges: [][]int{
					{2, 6},
					{1, 6},
					{3, 1},
					{2, 5},
					{4, 3},
					{3, 2},
					{1, 2},
					{1, 4},
				},
			},
			want: [][]int{
				{0, 11},
				{1, 6},
				{8, 9},
				{7, 10},
				{2, 3},
				{4, 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.m, tt.args.s, tt.args.edges); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
