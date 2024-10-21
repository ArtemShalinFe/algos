package main

import (
	"reflect"
	"testing"
)

func TestApp_solution(t *testing.T) {
	type fields struct {
		res []string
	}
	type args struct {
		str string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		{
			name: "1",
			fields: fields{
				res: make([]string, 0),
			},
			args: args{
				str: "23",
			},
			want: []string{
				"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf",
			},
		},
		{
			name: "2",
			fields: fields{
				res: make([]string, 0),
			},
			args: args{
				str: "92",
			},
			want: []string{
				"wa", "wb", "wc", "xa", "xb", "xc", "ya", "yb", "yc", "za", "zb", "zc",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &App{
				res: tt.fields.res,
			}
			if got := a.solution(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("App.solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
