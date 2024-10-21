package main

import (
	"reflect"
	"testing"
)

func Test_solution(t *testing.T) {
	type args struct {
		size int
		cmds []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "1",
			args: args{
				size: 2,
				cmds: []string{
					"peek",
					"push 5",
					"push 2",
					"peek",
					"size",
					"size",
					"push 1",
					"size",
				},
			},
			want: []string{
				"None",
				"5",
				"2",
				"2",
				"error",
				"2",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.size, tt.args.cmds); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
