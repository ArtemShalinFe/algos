package main

import (
	"strings"
	"testing"
)

func Test_solution(t *testing.T) {
	type args struct {
		arr string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "yandex 1",
			args: args{
				arr: "2 1 + 3 *",
			},
			want: 9,
		},
		{
			name: "yandex 2",
			args: args{
				arr: "7 2 + 4 * 2 +",
			},
			want: 38,
		},
		{
			name: "yandex 3",
			args: args{
				arr: "2 5 - 4 /",
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(strings.Split(tt.args.arr, " ")); got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
