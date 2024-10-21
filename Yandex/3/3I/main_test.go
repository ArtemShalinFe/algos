package main

import "testing"

func Test_solution(t *testing.T) {
	type args struct {
		n   int
		arr []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "0",
			args: args{
				n:   3,
				arr: []int{1, 1, 1, 2, 2, 3, 7, 5, 5},
			},
			want: "1 2 5",
		},
		{
			name: "1",
			args: args{
				n:   3,
				arr: []int{1, 2, 3, 1, 2, 3, 4},
			},
			want: "1 2 3",
		},
		{
			name: "2",
			args: args{
				n:   1,
				arr: []int{1, 1, 1, 2, 2, 3},
			},
			want: "1",
		},
		{
			name: "3",
			args: args{
				n:   1,
				arr: []int{1, 2, 2, 2, 2, 3},
			},
			want: "2",
		},
		{
			name: "4",
			args: args{
				n:   3,
				arr: []int{5, 5, 3, 3, 4, 4},
			},
			want: "5 3 4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.n, tt.args.arr); got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
