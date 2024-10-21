package main

import (
	"reflect"
	"testing"
)

func Test_solution(t *testing.T) {
	type args struct {
		keys    []string
		querrys []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "1",
			args: args{
				keys: []string{
					"i love coffee",
					"coffee with milk and sugar",
					"free tea for everyone",
				},
				querrys: []string{
					"i like black coffee without milk",
					"everyone loves new year",
					"mary likes black coffee without milk",
				},
			},
			want: []string{
				"1 2",
				"3",
				"2 1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.keys, tt.args.querrys); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
