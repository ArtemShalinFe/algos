package main

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				n: 8,
			},
			want: []int{2, 2, 2},
		},
		{
			name: "2",
			args: args{
				n: 13,
			},
			want: []int{13},
		},
		{
			name: "3",
			args: args{
				n: 100,
			},
			want: []int{2, 2, 5, 5},
		},
		{
			name: "4",
			args: args{
				n: 917521579,
			},
			want: []int{13, 70578583},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
