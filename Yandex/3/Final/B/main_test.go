package main

import (
	"reflect"
	"testing"
)

func Test_solution(t *testing.T) {
	type args struct {
		ps []Participant
	}
	tests := []struct {
		name string
		args args
		want []Participant
	}{
		{
			name: "1",
			args: args{
				ps: []Participant{
					{name: "alla", solved: 4, fine: 100},
					{name: "gena", solved: 6, fine: 1000},
					{name: "gosha", solved: 2, fine: 90},
					{name: "rita", solved: 2, fine: 90},
					{name: "timofey", solved: 4, fine: 80},
				},
			},
			want: []Participant{
				{name: "gena", solved: 6, fine: 1000},
				{name: "timofey", solved: 4, fine: 80},
				{name: "alla", solved: 4, fine: 100},
				{name: "gosha", solved: 2, fine: 90},
				{name: "rita", solved: 2, fine: 90},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.ps); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
