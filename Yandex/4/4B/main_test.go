package main

import "testing"

// func solution(results []int) int {
// 	var max int
// 	var allSum int

// 	resultIdx := make(map[int][]int)
// 	resultIdx[0] = append(resultIdx[0], 0)

// 	for i := 0; i < len(results); i++ {
// 		if results[i] == 0 {
// 			allSum++
// 			resultIdx[allSum] = append(resultIdx[allSum], i+1)
// 		} else {
// 			allSum--
// 			resultIdx[allSum] = append(resultIdx[allSum], i+1)
// 		}
// 	}
// 	for _, v := range resultIdx {
// 		if len(v) == 0 {
// 			continue
// 		}
// 		curr := v[len(v)-1] - v[0]
// 		if curr > max {
// 			max = curr
// 		}
// 	}

//		return max
//	}
func Test_solution(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "0",
			args: args{
				arr: []int{0, 0, 1, 0, 0, 0, 1, 0, 0, 1},
			},
			want: 4,
		},
		{
			name: "1",
			args: args{
				arr: []int{0, 0, 1, 0, 1, 1, 1, 0, 0, 0},
			},
			want: 8,
		},
		{
			name: "2",
			args: args{
				arr: []int{0, 1},
			},
			want: 2,
		},
		{
			name: "3",
			args: args{
				arr: []int{0, 1, 0},
			},
			want: 2,
		},
		{
			name: "3",
			args: args{
				arr: []int{0, 0, 0, 0, 0, 0, 1, 1, 1, 1},
			},
			want: 8,
		},
		{
			name: "3",
			args: args{
				arr: []int{0},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.arr); got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
