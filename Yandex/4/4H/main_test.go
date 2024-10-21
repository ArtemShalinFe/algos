package main

import "testing"

func Test_solution(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				str: "abcabcbb",
			},
			want: "3",
		},
		{
			name: "2",
			args: args{
				str: "bbbbb",
			},
			want: "1",
		},
		{
			name: "3",
			args: args{
				str: "oqwieruskjdfhqwoieur",
			},
			want: "13",
		},
		{
			name: "4",
			args: args{
				str: "awe",
			},
			want: "3",
		},
		{
			name: "5",
			args: args{
				str: "ojodx",
			},
			want: "4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.str); got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
