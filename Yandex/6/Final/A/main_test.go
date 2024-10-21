package main

import "testing"

func Test_solution(t *testing.T) {
	type args struct {
		mx [][]*Edge
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				mx: [][]*Edge{
					{},
					{{Weight: 5, NodeID: 2}, {Weight: 6, NodeID: 3}},
					{{Weight: 5, NodeID: 1}, {Weight: 8, NodeID: 4}},
					{{Weight: 6, NodeID: 1}, {Weight: 3, NodeID: 4}},
					{{Weight: 3, NodeID: 3}, {Weight: 8, NodeID: 2}},
				},
			},
			want: "19",
		},
		{
			name: "2",
			args: args{
				mx: [][]*Edge{
					{},
					{{Weight: 1, NodeID: 2}, {Weight: 2, NodeID: 2}},
					{{Weight: 1, NodeID: 1}, {Weight: 2, NodeID: 1}, {Weight: 1, NodeID: 3}},
					{{Weight: 1, NodeID: 2}},
				},
			},
			want: "3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.mx); got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
