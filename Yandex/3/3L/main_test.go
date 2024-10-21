package main

import "testing"

func Test_solution(t *testing.T) {
	type args struct {
		arr []int
		n   int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{
			name: "1",
			args: args{
				arr: []int{
					1, 2, 4, 4, 6, 8,
				},
				n: 3,
			},
			want:  3,
			want1: 5,
		},
		{
			name: "2",
			args: args{
				arr: []int{
					1, 2, 4, 4, 4, 4,
				},
				n: 3,
			},
			want:  3,
			want1: -1,
		},
		{
			name: "3",
			args: args{
				arr: []int{
					1, 2, 4, 4, 4, 4,
				},
				n: 10,
			},
			want:  -1,
			want1: -1,
		},
		{
			name: "4",
			args: args{
				arr: []int{
					1, 2, 4, 4, 4, 4,
				},
				n: 1,
			},
			want:  1,
			want1: 2,
		},
		{
			name: "5",
			args: args{
				arr: []int{
					999900, 999900, 999900, 999900, 999900, 999900,
				},
				n: 999900,
			},
			want:  1,
			want1: -1,
		},
		{
			name: "6",
			args: args{
				arr: []int{
					1, 1, 1, 2, 2, 2, 4, 6, 15, 22, 200, 2123, 21236, 212315, 212600, 212800, 212900,
				},
				n: 212400,
			},
			want:  15,
			want1: -1,
		},
		{
			name: "7",
			args: args{
				arr: []int{
					1, 1, 4, 4, 4, 4, 8,
				},
				n: 4,
			},
			want:  3,
			want1: 7,
		},
		{
			name: "8",
			args: args{
				arr: genInts(1000000),
				n:   999900,
			},
			want:  1,
			want1: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := solution(len(tt.args.arr), tt.args.arr, tt.args.n)
			if got != tt.want {
				t.Errorf("solution() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("solution() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func genInts(count int) []int {
	var ints []int
	for i := 0; i < count; i++ {
		ints = append(ints, 999909)
	}

	return ints
}

// func BenchmarkSolution(b *testing.B) {
// 	type args struct {
// 		arr []int
// 		n   int
// 	}
// 	tests := []struct {
// 		name  string
// 		args  args
// 		want  int
// 		want1 int
// 	}{
// 		{
// 			name: "1",
// 			args: args{
// 				arr: []int{
// 					1, 2, 4, 4, 6, 8,
// 				},
// 				n: 3,
// 			},
// 			want:  3,
// 			want1: 5,
// 		},
// 		{
// 			name: "2",
// 			args: args{
// 				arr: []int{
// 					1, 2, 4, 4, 4, 4,
// 				},
// 				n: 3,
// 			},
// 			want:  3,
// 			want1: -1,
// 		},
// 		{
// 			name: "2",
// 			args: args{
// 				arr: []int{
// 					1, 2, 4, 4, 4, 4,
// 				},
// 				n: 10,
// 			},
// 			want:  -1,
// 			want1: -1,
// 		},
// 	}
// 	for _, tt := range tests {
// 		b.Run(tt.name, func(t *testing.B) {
// 			got, got1 := solution(tt.args.arr, tt.args.n)
// 			if got != tt.want {
// 				t.Errorf("solution() got = %v, want %v", got, tt.want)
// 			}
// 			if got1 != tt.want1 {
// 				t.Errorf("solution() got1 = %v, want %v", got1, tt.want1)
// 			}
// 		})
// 	}
// }

// func BenchmarkSolutionRec(b *testing.B) {
// 	type args struct {
// 		arr []int
// 		n   int
// 	}
// 	tests := []struct {
// 		name  string
// 		args  args
// 		want  int
// 		want1 int
// 	}{
// 		{
// 			name: "1",
// 			args: args{
// 				arr: []int{
// 					1, 2, 4, 4, 6, 8,
// 				},
// 				n: 3,
// 			},
// 			want:  3,
// 			want1: 5,
// 		},
// 		{
// 			name: "2",
// 			args: args{
// 				arr: []int{
// 					1, 2, 4, 4, 4, 4,
// 				},
// 				n: 3,
// 			},
// 			want:  3,
// 			want1: -1,
// 		},
// 		{
// 			name: "2",
// 			args: args{
// 				arr: []int{
// 					1, 2, 4, 4, 4, 4,
// 				},
// 				n: 10,
// 			},
// 			want:  -1,
// 			want1: -1,
// 		},
// 	}
// 	for _, tt := range tests {
// 		b.Run(tt.name, func(t *testing.B) {
// 			got, got1 := solution(tt.args.arr, tt.args.n)
// 			if got != tt.want {
// 				t.Errorf("solution() got = %v, want %v", got, tt.want)
// 			}
// 			if got1 != tt.want1 {
// 				t.Errorf("solution() got1 = %v, want %v", got1, tt.want1)
// 			}
// 		})
// 	}
// }
