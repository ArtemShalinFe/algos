package main

import "testing"

func Test_solution(t *testing.T) {
	type args struct {
		borders []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				borders: []int{6, 3, 3, 2},
			},
			want: 8,
		},
		{
			name: "2",
			args: args{
				borders: []int{5, 3, 7, 2, 8, 3},
			},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.borders); got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
