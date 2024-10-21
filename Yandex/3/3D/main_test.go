package main

import "testing"

func Test_solution(t *testing.T) {
	type args struct {
		childCap     []int
		coockieSizes []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				childCap:     []int{1, 2},
				coockieSizes: []int{2, 1, 3},
			},
			want: 2,
		},
		{
			name: "2",
			args: args{
				childCap:     []int{2, 1, 3},
				coockieSizes: []int{1, 1},
			},
			want: 1,
		},
		{
			name: "3",
			args: args{
				childCap:     []int{8, 5, 5, 8, 6, 9, 8, 2, 4, 7},
				coockieSizes: []int{9, 8, 1, 1, 1, 5, 10, 8},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.childCap, tt.args.coockieSizes); got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
