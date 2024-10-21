package main

import (
	"reflect"
	"testing"
)

func Test_solution(t *testing.T) {
	type args struct {
		a     int
		m     int
		n     int
		s     string
		pairs []pair
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				a: 1000,
				m: 1000009,
				n: 7,
				s: "abcdefgh",
				pairs: []pair{
					{left: 1, right: 1},
					{left: 1, right: 5},
					{left: 2, right: 3},
					{left: 3, right: 4},
					{left: 4, right: 4},
					{left: 1, right: 8},
					{left: 5, right: 8},
				},
			},
			want: []int{
				97,
				225076,
				98099,
				99100,
				100,
				436420,
				193195,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.a, tt.args.m, tt.args.n, tt.args.s, tt.args.pairs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
