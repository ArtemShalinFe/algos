package main

import (
	"reflect"
	"testing"
)

func Test_solution(t *testing.T) {
	type args struct {
		s   int
		arr []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "0",
			args: args{
				s:   0,
				arr: []int{1, 0, -1, 0, 2, -2},
			},
			want: [][]int{
				{-2, -1, 1, 2},
				{-2, 0, 0, 2},
				{-1, 0, 0, 1},
			},
		},
		{
			name: "1",
			args: args{
				s:   10,
				arr: []int{2, 3, 2, 4, 1, 10, 3, 0},
			},
			want: [][]int{
				{0, 3, 3, 4},
				{1, 2, 3, 4},
				{2, 2, 3, 3},
			},
		},
		{
			name: "2",
			args: args{
				s:   381350422,
				arr: []int{-142721627, 206575381, 564131650, -648707194, 995752548, 393312490, 435642014, -9069088, 326565756, 140484837, -533059899, 488966447, 0, -780815037, 699133600, -268205879, -70733143, 496260289, -777196728, 85424651},
			},
			want: [][]int{
				{-648707194, 140484837, 393312490, 496260289},
				{-268205879, 0, 85424651, 564131650},
				{-142721627, -9069088, 206575381, 326565756},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.s, tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
