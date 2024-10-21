package main

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	type args struct {
		ints []int
		i    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				ints: []int{1, 2, 0, 0},
				i:    34,
			},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "2",
			args: args{
				ints: []int{9, 5},
				i:    17,
			},
			want: []int{1, 1, 2},
		},
		{
			name: "3",
			args: args{
				ints: []int{0, 7},
				i:    3,
			},
			want: []int{1, 0},
		},
		{
			name: "4",
			args: args{
				ints: []int{0},
				i:    3,
			},
			want: []int{3},
		},
		{
			name: "5",
			args: args{
				ints: []int{1},
				i:    1000,
			},
			want: []int{1, 0, 0, 1},
		},
		{
			name: "6",
			args: args{
				ints: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 0},
				i:    12345678910,
			},
			want: []int{2, 4, 6, 9, 1, 3, 5, 7, 8, 2, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.ints, tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
