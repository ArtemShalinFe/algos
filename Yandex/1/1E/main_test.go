package main

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		l    int
		text string
	}
	tests := []struct {
		name string
		args args
		want *Word
	}{
		{
			name: "1",
			args: args{
				l:    19,
				text: "i love segment tree",
			},
			want: &Word{
				Len:   7,
				Value: "segment",
			},
		},
		{
			name: "2",
			args: args{
				l:    21,
				text: "frog jumps from river",
			},
			want: &Word{
				Len:   5,
				Value: "jumps",
			},
		},
		{
			name: "3",
			args: args{
				l:    6,
				text: "mymbg",
			},
			want: &Word{
				Len:   5,
				Value: "mymbg",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.l, tt.args.text); got.Len != tt.want.Len || got.Value != tt.want.Value {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
