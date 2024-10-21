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
		want [][]int
	}{
		{
			name: "1",
			args: args{
				m: 5,
				edges: [][]int{
					{1, 3},
					{2, 3},
					{5, 2},
				},
			},
			want: [][]int{
				{},
				{3},
				{3},
				{},
				{},
				{2},
			},
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
