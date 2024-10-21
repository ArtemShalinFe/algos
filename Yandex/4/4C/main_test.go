package main

import "testing"

func Test_solution(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "-1",
			args: args{
				a: "aba",
				b: "xxx",
			},
			want: "NO",
		},
		{
			name: "0",
			args: args{
				a: "mxyskaoghi",
				b: "qodfrgmslc",
			},
			want: "YES",
		},
		{
			name: "1",
			args: args{
				a: "aadd",
				b: "ccff",
			},
			want: "YES",
		},
		{
			name: "2",
			args: args{
				a: "agg",
				b: "xda",
			},
			want: "NO",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
