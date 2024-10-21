package main

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				n: 2,
			},
			want: 2,
		},
		{
			name: "1",
			args: args{
				n: 3,
			},
			want: 5,
		},
		{
			name: "1",
			args: args{
				n: 4,
			},
			want: 14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.n); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
