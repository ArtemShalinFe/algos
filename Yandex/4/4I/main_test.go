package main

import (
	"reflect"
	"testing"
)

func Test_solution(t *testing.T) {
	type args struct {
		str []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "1",
			args: args{
				str: []string{
					"tan", "eat", "tea", "ate", "nat", "bat",
				},
			},
			want: []string{
				"0 4",
				"1 2 3",
				"5",
			},
		},
		{
			name: "1",
			args: args{
				str: []string{
					"kfumpzsaao", "ilvnolmacs", "fpwmgnuonb", "gfwfwtwmvo", "bruucwbbjw", "uzpmkfaoas", "afkpuamszo",
				},
			},
			want: []string{
				"0 5 6",
				"1",
				"2",
				"3",
				"4",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
