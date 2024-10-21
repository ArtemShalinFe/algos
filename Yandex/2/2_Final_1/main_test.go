package main

import (
	"reflect"
	"testing"
)

func Test_solution(t *testing.T) {
	type args struct {
		dequeSize int
		cmds      []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "case yandex 1",
			args: args{
				dequeSize: 4,
				cmds: []string{
					"push_front 861",
					"push_front -819",
					"pop_back",
					"pop_back",
				},
			},
			want: []string{
				"861",
				"-819",
			},
		},
		{
			name: "case yandex 2",
			args: args{
				dequeSize: 10,
				cmds: []string{
					"push_front -855",
					"push_front 0",
					"pop_back",
					"pop_back",
					"push_back 844",
					"pop_back",
					"push_back 823",
				},
			},
			want: []string{
				"-855",
				"0",
				"844",
			},
		},
		{
			name: "case yandex 3",
			args: args{
				dequeSize: 6,
				cmds: []string{
					"push_front -201",
					"push_back 959",
					"push_back 102",
					"push_front 20",
					"pop_front",
					"pop_back",
				},
			},
			want: []string{
				"20",
				"102",
			},
		},
		{
			name: "case yandex 4",
			args: args{
				dequeSize: 8,
				cmds: []string{
					"push_back -977",
					"pop_back",
					"pop_back",
					"push_front -86",
					"pop_back",
					"push_back 81",
					"pop_front",
					"pop_back",
					"pop_back",
				},
			},
			want: []string{
				"-977",
				"error",
				"-86",
				"81",
				"error",
				"error",
			},
		},
		{
			name: "case 4",
			args: args{
				dequeSize: 10,
				cmds: []string{
					"push_front 1",
					"push_back 2",
					"push_front 3",
					"push_back 4",
					"push_front 5",
					"push_back 6",
					"push_front 7",
					"push_back 8",
					"pop_back",
					"pop_back",
					"pop_back",
					"pop_back",
					"pop_front",
					"pop_front",
					"pop_front",
					"pop_front",
					"pop_back",
					"pop_front",
				},
			},
			want: []string{
				"8",
				"6",
				"4",
				"2",
				"7",
				"5",
				"3",
				"1",
				"error",
				"error",
			},
		},
		{
			name: "case when deque is full",
			args: args{
				dequeSize: 6,
				cmds: []string{
					"push_front 10",
					"push_back 20",
					"push_front 30",
					"push_back 40",
					"push_front 50",
					"push_back 60",
					"push_front 70",
					"push_back 80",
					"pop_back",
					"pop_back",
					"pop_back",
					"pop_back",
					"pop_back",
					"pop_back",
					"pop_front",
					"pop_front",
				},
			},
			want: []string{
				"error",
				"error",
				"60",
				"40",
				"20",
				"10",
				"30",
				"50",
				"error",
				"error",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := solution(tt.args.dequeSize, tt.args.cmds); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
