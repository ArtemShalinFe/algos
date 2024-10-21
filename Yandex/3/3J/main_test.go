package main

import (
	"reflect"
	"testing"
)

func Test_solution(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "1",
			args: args{
				arr: []int{4, 3, 9, 2, 1},
			},
			want: [][]int{
				{3, 4, 2, 1, 9},
				{3, 2, 1, 4, 9},
				{2, 1, 3, 4, 9},
				{1, 2, 3, 4, 9},
			},
		},
		{
			name: "2",
			args: args{
				arr: []int{4, 5},
			},
			want: [][]int{
				{4, 5},
			},
		},
		{
			name: "3",
			args: args{
				arr: []int{1, 1, 1, 1, 1},
			},
			want: [][]int{
				{1, 1, 1, 1, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
