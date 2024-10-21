package main

import (
	"reflect"
	"testing"
)

func Test_solution(t *testing.T) {
	type args struct {
		ps []*Participant
	}
	tests := []struct {
		name string
		args args
		want []*Participant
	}{
		{
			name: "1",
			args: args{
				ps: []*Participant{
					{name: "alla", solved: 4, fine: 100},
					{name: "gena", solved: 6, fine: 1000},
					{name: "gosha", solved: 2, fine: 90},
					{name: "rita", solved: 2, fine: 90},
					{name: "timofey", solved: 4, fine: 80},
				},
			},
			want: []*Participant{
				{name: "gena", solved: 6, fine: 1000},
				{name: "timofey", solved: 4, fine: 80},
				{name: "alla", solved: 4, fine: 100},
				{name: "gosha", solved: 2, fine: 90},
				{name: "rita", solved: 2, fine: 90},
			},
		},
		{
			name: "2",
			args: args{
				ps: []*Participant{
					{name: "alla", solved: 0, fine: 10},
					{name: "gena", solved: 0, fine: 9},
					{name: "gosha", solved: 0, fine: 8},
					{name: "rita", solved: 0, fine: 7},
					{name: "timofey", solved: 0, fine: 0},
				},
			},
			want: []*Participant{
				{name: "timofey", solved: 0, fine: 0},
				{name: "rita", solved: 0, fine: 7},
				{name: "gosha", solved: 0, fine: 8},
				{name: "gena", solved: 0, fine: 9},
				{name: "alla", solved: 0, fine: 10},
			},
		},
		{
			name: "3",
			args: args{
				ps: []*Participant{
					{name: "alpyx", solved: 3, fine: 49},
					{name: "qly", solved: 62, fine: 28},
					{name: "cpbablljhrnlejcwq", solved: 45, fine: 32},
					{name: "yusoldfqnjlvwrcruts", solved: 81, fine: 39},
					{name: "kzmleudk", solved: 52, fine: 89},
					{name: "hpnpszayf", solved: 4, fine: 12},
					{name: "ojywnxuyqgmf", solved: 54, fine: 31},
					{name: "tvomvisbqrflkghnojt", solved: 49, fine: 79},
					{name: "yc", solved: 17, fine: 84},
					{name: "k", solved: 16, fine: 14},
					{name: "dgmqeqezjgzxfwhcn", solved: 68, fine: 21},
					{name: "zdbonelpgmbou", solved: 47, fine: 64},
					{name: "ehtxaytzgcsim", solved: 21, fine: 66},
					{name: "burbwymmeal", solved: 56, fine: 80},
					{name: "wqzynngqasdjj", solved: 8, fine: 15},
					{name: "auirlnkhxtmmytr", solved: 23, fine: 14},
					{name: "poypocw", solved: 0, fine: 72},
					{name: "uf", solved: 40, fine: 68},
					{name: "ezfetpflzoi", solved: 0, fine: 30},
					{name: "ygegvjzlfgumbo", solved: 80, fine: 33},
					{name: "rcwp", solved: 8, fine: 30},
					{name: "oaes", solved: 7, fine: 59},
					{name: "zgbktifgwvsgesleqclp", solved: 4, fine: 8},
					{name: "knzsujegjqmer", solved: 99, fine: 72},
				},
			},
			want: []*Participant{
				{name: "knzsujegjqmer", solved: 99, fine: 72},
				{name: "yusoldfqnjlvwrcruts", solved: 81, fine: 39},
				{name: "ygegvjzlfgumbo", solved: 80, fine: 33},
				{name: "dgmqeqezjgzxfwhcn", solved: 68, fine: 21},
				{name: "qly", solved: 62, fine: 28},
				{name: "burbwymmeal", solved: 56, fine: 80},
				{name: "ojywnxuyqgmf", solved: 54, fine: 31},
				{name: "kzmleudk", solved: 52, fine: 89},
				{name: "tvomvisbqrflkghnojt", solved: 49, fine: 79},
				{name: "zdbonelpgmbou", solved: 47, fine: 64},
				{name: "cpbablljhrnlejcwq", solved: 45, fine: 32},
				{name: "uf", solved: 40, fine: 68},
				{name: "auirlnkhxtmmytr", solved: 23, fine: 14},
				{name: "ehtxaytzgcsim", solved: 21, fine: 66},
				{name: "yc", solved: 17, fine: 84},
				{name: "k", solved: 16, fine: 14},
				{name: "wqzynngqasdjj", solved: 8, fine: 15},
				{name: "rcwp", solved: 8, fine: 30},
				{name: "oaes", solved: 7, fine: 59},
				{name: "zgbktifgwvsgesleqclp", solved: 4, fine: 8},
				{name: "hpnpszayf", solved: 4, fine: 12},
				{name: "alpyx", solved: 3, fine: 49},
				{name: "ezfetpflzoi", solved: 0, fine: 30},
				{name: "poypocw", solved: 0, fine: 72},
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
