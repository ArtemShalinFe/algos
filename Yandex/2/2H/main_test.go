package main

import (
	"testing"
)

func TestSolution(t *testing.T) {
	type args struct {
		brackets string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				brackets: "{[()]}",
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				brackets: "{}",
			},
			want: true,
		},
		{
			name: "3",
			args: args{
				brackets: "",
			},
			want: true,
		},
		{
			name: "4",
			args: args{
				brackets: "[])",
			},
			want: false,
		},
		{
			name: "5",
			args: args{
				brackets: "[)",
			},
			want: false,
		},
		{
			name: "6",
			args: args{
				brackets: "((((((((((((((((((((((((((((((((((((((())))))))))))))))))))))))))))))))))))))))",
			},
			want: false,
		},
		{
			name: "7",
			args: args{
				brackets: "((())",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.brackets); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
