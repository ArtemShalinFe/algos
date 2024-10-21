package main

import (
	"reflect"
	"testing"
)

func Test_solution(t *testing.T) {
	type args struct {
		n   int
		k   int
		str string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// {
		// 	name: "1",
		// 	args: args{
		// 		n:   10,
		// 		k:   2,
		// 		str: "gggggooooogggggoooooogggggssshaa",
		// 	},
		// 	want: []int{0, 5},
		// },
		// {
		// 	name: "2",
		// 	args: args{
		// 		n:   3,
		// 		k:   4,
		// 		str: "allallallallalla",
		// 	},
		// 	want: []int{0, 1, 2},
		// },
		{
			name: "3",
			args: args{
				n:   4,
				k:   2,
				str: "keogaokkfawircucxvilarsdwlbubqabpnkpqbrcuckxsrmtaechlqbvsqhsoghmpoqqkshwcrismripmriphirknsodxflisozvcldtffjfmxzqjrmcjbojwgmxzqchugkshwwjcmkeqiuwsqwlbuotwhymhmsdrcucffhmliwlbuojwgsbsdxqmxzqjrmcqbuucldtsbsdnsodudlstcrejryudxuuchugiabkiacldtiaawfftkllbkbqabrejrwlbujbcrisxvilkxaohcffjrmcfflfymtvbrojwgnsodlqbvsrmtqqjhlczfkksoxvilarqaohjbpxzdtcuwnaiamobpthxvilchugjrmcarxffawichuwnajfymudxlikoruurnxjvwqbmxzqlczftvbralnsodmripaqljvunejrmcwvelxvilvunezhhcyudxbkwlbuqbmriprejrvwjfvqwnmripkxchvwwvelqqnsodidlmhvxvilxlikxqpnkpaqljlqbvffjbyhxfinqbbqabrcucrejrketichuwsqkxtvbrnsodcrisyhdgymevxvildgjrmccrisezlmhvqmyqwjcmtctczharllorlczfymvqwntkuwnaxfohwlbuuwnatichzmcrisjrmczhnsodbpthbpthliaemolftrvepxzdqbtcuwnatichuupnkpawtrveqknpvwrejrawkjrnvqwnotwhkshwmripqakjrnohsrmtsbsddgwjcmpxzdjbewqaevlqbvrejrmosoyudxkjrnlschugqknpuwnaqbqajffawixqrejrezuwnaorqbhmkeqiwveljrmckshwlixliktichbpthotwhinkkxftclczfkevwdgdgagaduevwevvqwnkshworezbqabotwhffkkbpthuulqbvchugtichmosoxflfdgzvbklfvqwnzvpxzdiddgchugagadtichmxzqll",
			},
			want: []int{12, 76, 108, 68, 24, 147, 118, 109, 110, 111, 112, 100, 178, 88, 126, 28, 214, 72, 16, 52, 53, 44, 17, 18, 8, 340, 314, 290, 426, 218, 434, 384, 32, 421, 422, 251, 37, 213, 371, 372, 142, 503, 359, 500, 134, 474, 552, 348, 334, 338, 339, 647, 648, 649, 515, 686, 150, 730, 199, 429, 718, 367, 137, 138, 673, 315, 792, 473, 899, 900, 949, 401},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.n, tt.args.k, tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
