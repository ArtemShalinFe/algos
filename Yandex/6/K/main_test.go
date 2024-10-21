package main

import (
	"math"
	"reflect"
	"testing"
)

var maxInt = math.MaxInt32

func Test_solution(t *testing.T) {
	type args struct {
		n  int
		m  int
		mx [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "1",
			args: args{
				n: 4,
				m: 4,
				mx: [][]int{
					//1  2   3   4
					{0, 1, maxInt, 2}, // 1
					{1, 0, 3, maxInt}, // 2
					{maxInt, 3, 0, 5}, // 3
					{2, maxInt, 5, 0}, // 4
				},
			},
			want: [][]int{
				{0, 1, 4, 2},
				{1, 0, 3, 3},
				{4, 3, 0, 5},
				{2, 3, 5, 0},
			},
		},
		{
			name: "2",
			args: args{
				n: 3,
				m: 2,
				mx: [][]int{
					//1  2   3
					{0, 1, maxInt},      // 1
					{1, 0, maxInt},      // 2
					{maxInt, maxInt, 0}, // 3
				},
			},
			want: [][]int{
				//1  2   3
				{0, 1, -1},  // 1
				{1, 0, -1},  // 2
				{-1, -1, 0}, // 3
			},
		},
		{
			name: "3",
			args: args{
				n: 2,
				m: 0,
				mx: [][]int{
					//	1		2
					{0, maxInt}, // 1
					{maxInt, 0}, // 2
				},
			},
			want: [][]int{
				//1   2
				{0, -1}, // 1
				{-1, 0}, // 2
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.n, tt.args.m, tt.args.mx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
