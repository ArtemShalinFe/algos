package main

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		t int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				t: 16,
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				t: 14,
			},
			want: false,
		},
		{
			name: "3",
			args: args{
				t: 1,
			},
			want: true,
		},
		{
			name: "4",
			args: args{
				t: 5,
			},
			want: false,
		},
		{
			name: "5",
			args: args{
				t: 2408,
			},
			want: false,
		},
		{
			name: "6",
			args: args{
				t: 356,
			},
			want: false,
		},
		{
			name: "7",
			args: args{
				t: 8304,
			},
			want: false,
		},
		{
			name: "8",
			args: args{
				t: 1024,
			},
			want: true,
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
