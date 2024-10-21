package main

import "testing"

func Test_solution(t *testing.T) {
	type args struct {
		money []int
		homes []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				money: []int{3, 300},
				homes: []int{999, 999, 999},
			},
			want: 0,
		},
		{
			name: "2",
			args: args{
				money: []int{3, 1000},
				homes: []int{350, 999, 200},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.money, tt.args.homes); got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
