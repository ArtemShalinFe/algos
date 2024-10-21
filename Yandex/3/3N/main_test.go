package main

import (
	"reflect"
	"testing"
)

func Test_solution(t *testing.T) {
	type args struct {
		arr [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "1",
			args: args{
				arr: [][]int{
					{7, 8},
					{7, 8},
					{2, 3},
					{6, 10},
				},
			},
			want: [][]int{
				{2, 3},
				{6, 10},
			},
		},
		{
			name: "2",
			args: args{
				arr: [][]int{
					{1, 3},
					{3, 5},
					{4, 6},
					{5, 6},
					{2, 4},
					{7, 10},
				},
			},
			want: [][]int{
				{1, 6},
				{7, 10},
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
