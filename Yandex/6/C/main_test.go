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
		want []int
	}{
		{
			name: "1",
			args: args{
				s: 3,
				m: 4,
				edges: [][]int{
					{3, 2},
					{4, 3},
					{1, 4},
					{1, 2},
				},
			},
			want: []int{3, 2, 1, 4},
		},
		{
			name: "2",
			args: args{
				s: 1,
				m: 2,
				edges: [][]int{
					{2, 1},
					{1, 2},
				},
			},
			want: []int{1, 2},
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
