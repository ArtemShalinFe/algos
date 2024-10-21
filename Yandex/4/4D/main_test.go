package main

import "testing"

func Test_solution(t *testing.T) {
	type args struct {
		a int
		m int
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				a: 123,
				m: 100003,
				s: "a",
			},
			want: 97,
		},
		{
			name: "2",
			args: args{
				a: 123,
				m: 100003,
				s: "hash",
			},
			want: 6080,
		},
		{
			name: "3",
			args: args{
				a: 123,
				m: 100003,
				s: "HaSH",
			},
			want: 56156,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.a, tt.args.m, tt.args.s); got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
