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
		n int
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
				n: 3,
			},
			want: []string{
				"((()))",
				"(()())",
				"(())()",
				"()(())",
				"()()()",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &App{
				res: tt.fields.res,
			}
			if got := a.solution(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("App.solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
