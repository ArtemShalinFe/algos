package main

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				text: "A man, a plan, a canal: Panama",
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				text: "zo",
			},
			want: false,
		},
		{
			name: "3",
			args: args{
				text: "ThiS_String-is-@-PALIND0m3```~3m0DNILAP-()-si*!gnirts>>>siht",
			},
			want: true,
		},
		{
			name: "4",
			args: args{
				text: "10 ten animals I slam in a net 012",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.text); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
