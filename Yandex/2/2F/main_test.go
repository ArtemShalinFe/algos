package main

import (
	"reflect"
	"testing"
)

func Test_solution(t *testing.T) {
	type args struct {
		cmds []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "1",
			args: args{[]string{
				"pop",
				"pop",
				"top",
				"push 4",
				"push -5",
				"top",
				"push 7",
				"pop",
				"pop",
				"get_max",
				"top",
				"pop",
				"get_max",
			}},
			want: []string{
				"error",
				"error",
				"error",
				"-5",
				"4",
				"4",
				"None",
			},
		},
		{
			name: "2",
			args: args{[]string{
				"push 6",
				"push 6",
				"get_max",
				"pop",
				"pop",
				"pop",
			}},
			want: []string{
				"6",
				"error",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.cmds); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
