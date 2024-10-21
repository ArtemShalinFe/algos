package main

import "testing"

func Test_solution(t *testing.T) {
	type args struct {
		arr1 []int
		arr2 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				arr1: []int{1, 2, 3, 2, 1},
				arr2: []int{3, 2, 1, 5, 6},
			},
			want: 3,
		},
		{
			name: "2",
			args: args{
				arr1: []int{1, 2, 3, 4, 5},
				arr2: []int{4, 5, 9},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.arr1, tt.args.arr2); got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
