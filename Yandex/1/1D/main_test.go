package main

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		t   int
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				t:   7,
				arr: []int{-1, -10, -8, 0, 2, 0, 5},
			},
			want: 3,
		},
		{
			name: "2",
			args: args{
				t:   40,
				arr: []int{-254, -234, -223, -214, -211, -211, -208, -195, -175, -175, -164, -146, -142, -126, -116, -109, -104, -93, -87, -82, 5, 8, 19, 30, 105, 108, 111, 112, 123, 130, 169, 171, 191, 200, 234, 237, 245, 251, 253, 261},
			},
			want: 1,
		},
		{
			name: "3",
			args: args{
				t:   10,
				arr: []int{5, 5, 5, 5, 5, 5, 5, 5, 5, 5},
			},
			want: 0,
		},
		{
			name: "4",
			args: args{
				t:   1,
				arr: []int{-133},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.t, tt.args.arr); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
