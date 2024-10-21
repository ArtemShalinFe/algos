package main

import (
	"reflect"
	"testing"
)

func Test_solution(t *testing.T) {
	type args struct {
		arr []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "1",
			args: args{
				arr: []string{
					"вышивание крестиком",
					"рисование мелками на парте",
					"настольный керлинг",
					"настольный керлинг",
					"кухня африканского племени ужасмай",
					"тяжелая атлетика",
					"таракановедение",
					"таракановедение",
				},
			},
			want: []string{
				"вышивание крестиком",
				"рисование мелками на парте",
				"настольный керлинг",
				"кухня африканского племени ужасмай",
				"тяжелая атлетика",
				"таракановедение",
			},
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
