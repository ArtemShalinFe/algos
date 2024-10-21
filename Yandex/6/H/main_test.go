package main

import (
	"reflect"
	"testing"
)

func Test_solution(t *testing.T) {
	type args struct {
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
				m: 5,
				edges: [][]int{
					{3, 2},
					{3, 4},
					{2, 5},
				},
			},
			want: []int{1, 3, 2, 4, 5},
		},
		{
			name: "2",
			args: args{
				m: 6,
				edges: [][]int{
					{6, 4},
					{4, 1},
					{5, 1},
				},
			},
			want: []int{2, 3, 5, 6, 4, 1},
		},
		{
			name: "3",
			args: args{
				m:     4,
				edges: [][]int{},
			},
			want: []int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.m, tt.args.edges); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
