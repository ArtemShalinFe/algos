package main

import "testing"

func TestSolution(t *testing.T) {
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
			name: "1",
			args: args{
				a: "abcd",
				b: "abcde",
			},
			want: "e",
		},
		{
			name: "2",
			args: args{
				a: "go",
				b: "ogg",
			},
			want: "g",
		},
		{
			name: "3",
			args: args{
				a: "xtkpx",
				b: "xkctpx",
			},
			want: "c",
		},
		{
			name: "4",
			args: args{
				a: "11111",
				b: "111111",
			},
			want: "1",
		},
		{
			name: "5",
			args: args{
				a: "11111asdlkjsfdkjhasjkfhasdjkfhasjdkfhsajkfhsafsahdflkashdfkaslhdlksadfhalksdfh",
				b: "11111asdlkjsfdkjhasjkfhasdjkfhasj0dkfhsajkfhsafsahdflkashdfkaslhdlksadfhalksdfh",
			},
			want: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
