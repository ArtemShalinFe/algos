package main

import (
	"testing"
)

func Test_brokenSearch(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "-2",
			args: args{
				arr: []int{1},
				k:   1,
			},
			want: 0,
		},
		{
			name: "-1",
			args: args{
				arr: []int{19, 21, 100, 101, 1, 4, 5, 7, 12},
				k:   21,
			},
			want: 1,
		},
		{
			name: "0",
			args: args{
				arr: []int{3, 6, 7},
				k:   8,
			},
			want: -1,
		},

		{
			name: "1 yandex",
			args: args{
				arr: []int{19, 21, 100, 101, 1, 4, 5, 7, 12},
				k:   5,
			},
			want: 6,
		},
		{
			name: "2 yandex",
			args: args{
				arr: []int{5, 1},
				k:   1,
			},
			want: 1,
		},
		{
			name: "3 my",
			args: args{
				arr: []int{1, 2, 3, 4, 5, 6, 7, 8},
				k:   5,
			},
			want: 4,
		},
		{
			name: "4 my",
			args: args{
				arr: []int{5, 6, 7, 8, 9, 19, 29, 45, 1, 2, 3, 4},
				k:   4,
			},
			want: 11,
		},
		{
			name: "5 my",
			args: args{
				arr: []int{5, 6, 7, 8, 9, 19, 29, 45, 1, 2, 3, 4},
				k:   5,
			},
			want: 0,
		},
		{
			name: "6 my",
			args: args{
				arr: []int{5, 6, 7, 8, 9, 19, 29, 45, 1, 2, 3, 4},
				k:   6,
			},
			want: 1,
		},
		{
			name: "7 my",
			args: args{
				arr: []int{5, 6, 7, 8, 9, 19, 29, 45, 1, 2, 3, 4},
				k:   1,
			},
			want: 8,
		},
		{
			name: "8 my",
			args: args{
				arr: []int{3, 5, 6, 7, 9, 1, 2},
				k:   4,
			},
			want: -1,
		},
		{
			name: "9 my",
			args: args{
				arr: []int{6, 7, 10, 0, 2, 4, 5},
				k:   3,
			},
			want: -1,
		},
		{
			name: "10 my",
			args: args{
				arr: []int{1, 5},
				k:   2,
			},
			want: -1,
		},
		{
			name: "11 my",
			args: args{
				arr: []int{1},
				k:   2,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := brokenSearch(tt.args.arr, tt.args.k); got != tt.want {
				t.Errorf("brokenSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
