package main

import "testing"

func Test_solution(t *testing.T) {
	type args struct {
		substr string
		str    string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				substr: "abc",
				str:    "ahbgdcu",
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				substr: "abcp",
				str:    "ahpc",
			},
			want: false,
		},
		{
			name: "3",
			args: args{
				substr: "abcd",
				str:    "ahu",
			},
			want: false,
		},
		{
			name: "4",
			args: args{
				substr: "abcd",
				str:    "abcd",
			},
			want: true,
		},
		{
			name: "5",
			args: args{
				substr: "islx",
				str:    "yoytgtshldmogkdburkbcfvoapepjpcuwemusfkfztrzxstytrnarlizjhuoscuzlraezlaweipuuqdgvhwkhhoufexojaps",
			},
			want: true,
		},
		{
			name: "6",
			args: args{
				substr: "ijha",
				str:    "hmrqvftefyixinahlzgbkidroxiptbbkjmtwpsujevkulgrjiwiwzyhngulrodiwyg",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.substr, tt.args.str); got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
