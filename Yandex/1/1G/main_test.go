package main

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		t int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				t: 5,
			},
			want: "101",
		},
		{
			name: "2",
			args: args{
				t: 14,
			},
			want: "1110",
		},
		{
			name: "3",
			args: args{
				t: 9876,
			},
			want: "10011010010100",
		},
		{
			name: "4",
			args: args{
				t: 0,
			},
			want: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.t); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
