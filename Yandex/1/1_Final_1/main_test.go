package main

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	type args struct {
		ints []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				ints: []int{0, 1, 4, 9, 0},
			},
			want: []int{0, 1, 2, 1, 0},
		},
		{
			name: "2",
			args: args{
				ints: []int{0, 7, 9, 4, 8, 20},
			},
			want: []int{0, 1, 2, 3, 4, 5},
		},
		{
			name: "3",
			args: args{
				ints: []int{0, 7, 0, 4, 0, 20},
			},
			want: []int{0, 1, 0, 1, 0, 1},
		},
		{
			name: "4",
			args: args{
				ints: []int{0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1},
			},
			want: []int{0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1},
		},
		{
			name: "5",
			args: args{
				ints: []int{5, 4, 3, 2, 1, 0, 1, 2, 3, 4, 5},
			},
			want: []int{5, 4, 3, 2, 1, 0, 1, 2, 3, 4, 5},
		},
		{
			name: "6",
			args: args{
				ints: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
			},
			want: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		},
		{
			name: "7",
			args: args{
				ints: []int{0, 9898},
			},
			want: []int{0, 1},
		},
		{
			name: "8",
			args: args{
				ints: []int{0, 9898},
			},
			want: []int{0, 1},
		},
		{
			name: "9",
			args: args{
				ints: []int{0, 4, 3, 2, 1, 0, 1, 2, 3, 4, 0},
			},
			want: []int{0, 1, 2, 2, 1, 0, 1, 2, 2, 1, 0},
		},
		{
			name: "10",
			args: args{
				ints: []int{0, 4, 3, 2, 1, 0, 1, 2, 3, 4, 0, 1, 2, 3, 4, 0},
			},
			want: []int{0, 1, 2, 2, 1, 0, 1, 2, 2, 1, 0, 1, 2, 2, 1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.ints); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
