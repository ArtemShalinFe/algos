package main

import (
	"reflect"
	"testing"
)

func Test_solution(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// {
		// 	name: "1",
		// 	args: args{
		// 		arr: []int{0, 2, 1, 2, 0, 0, 1},
		// 	},
		// 	want: []int{0, 0, 0, 1, 1, 2, 2},
		// },
		// {
		// 	name: "2",
		// 	args: args{
		// 		arr: []int{2, 1},
		// 	},
		// 	want: []int{1, 2},
		// },
		{
			name: "3",
			args: args{
				arr: []int{},
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
