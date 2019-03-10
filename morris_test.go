package main

import (
	"fmt"
	"super-three-men-morris/board"
	"testing"
)

type testPG struct {
	input  string
	player string
	action string
	result bool
	state  board.Board
}

func TestPlayGame(t *testing.T) {
	game := board.Init()
	cases := []testPG{
		testPG{
			input:  "1,1",
			player: board.White,
			action: add,
			result: false,
			state: board.Board{
				{board.Empty, board.Empty, board.Empty},
				{board.Empty, board.White, board.Empty},
				{board.Empty, board.Empty, board.Empty},
			},
		},
		testPG{
			input:  "",
			player: board.Black,
			action: add,
			result: true,
			state: board.Board{
				{board.Empty, board.Empty, board.Empty},
				{board.Empty, board.White, board.Empty},
				{board.Empty, board.Empty, board.Empty},
			},
		},
		testPG{
			input:  "1,1",
			player: board.Black,
			action: add,
			result: true,
			state: board.Board{
				{board.Empty, board.Empty, board.Empty},
				{board.Empty, board.White, board.Empty},
				{board.Empty, board.Empty, board.Empty},
			},
		},
		testPG{
			input:  "1,2",
			player: board.Black,
			action: add,
			result: false,
			state: board.Board{
				{board.Empty, board.Empty, board.Empty},
				{board.Empty, board.White, board.Empty},
				{board.Empty, board.Black, board.Empty},
			},
		},
		testPG{
			input:  "0,0",
			player: board.White,
			action: add,
			result: false,
			state: board.Board{
				{board.White, board.Empty, board.Empty},
				{board.Empty, board.White, board.Empty},
				{board.Empty, board.Black, board.Empty},
			},
		},
		testPG{
			input:  "2,2",
			player: board.Black,
			action: add,
			result: false,
			state: board.Board{
				{board.White, board.Empty, board.Empty},
				{board.Empty, board.White, board.Empty},
				{board.Empty, board.Black, board.Black},
			},
		},
		testPG{
			input:  "0,2",
			player: board.White,
			action: add,
			result: false,
			state: board.Board{
				{board.White, board.Empty, board.Empty},
				{board.Empty, board.White, board.Empty},
				{board.White, board.Black, board.Black},
			},
		},
		testPG{
			input:  "0,1",
			player: board.Black,
			action: add,
			result: false,
			state: board.Board{
				{board.White, board.Empty, board.Empty},
				{board.Black, board.White, board.Empty},
				{board.White, board.Black, board.Black},
			},
		},
		// We can now perform moves and test them
		testPG{
			input:  "0,2 2,0",
			player: board.White,
			action: move,
			result: false,
			state: board.Board{
				{board.White, board.Empty, board.White},
				{board.Black, board.White, board.Empty},
				{board.Empty, board.Black, board.Black},
			},
		},
		testPG{
			input:  "0,2",
			player: board.Black,
			action: move,
			result: true,
			state: board.Board{
				{board.White, board.Empty, board.White},
				{board.Black, board.White, board.Empty},
				{board.Empty, board.Black, board.Black},
			},
		},
		testPG{
			input:  "0,2 ",
			player: board.Black,
			action: move,
			result: true,
			state: board.Board{
				{board.White, board.Empty, board.White},
				{board.Black, board.White, board.Empty},
				{board.Empty, board.Black, board.Black},
			},
		},
		testPG{
			input:  "0,2 asdfas",
			player: board.Black,
			action: move,
			result: true,
			state: board.Board{
				{board.White, board.Empty, board.White},
				{board.Black, board.White, board.Empty},
				{board.Empty, board.Black, board.Black},
			},
		},
		testPG{
			input:  "adsfads",
			player: board.Black,
			action: move,
			result: true,
			state: board.Board{
				{board.White, board.Empty, board.White},
				{board.Black, board.White, board.Empty},
				{board.Empty, board.Black, board.Black},
			},
		},
		testPG{
			input:  "adsfasd dafds",
			player: board.Black,
			action: move,
			result: true,
			state: board.Board{
				{board.White, board.Empty, board.White},
				{board.Black, board.White, board.Empty},
				{board.Empty, board.Black, board.Black},
			},
		},
		testPG{
			input:  "1,2 1,1",
			player: board.Black,
			action: move,
			result: true,
			state: board.Board{
				{board.White, board.Empty, board.White},
				{board.Black, board.White, board.Empty},
				{board.Empty, board.Black, board.Black},
			},
		},
		testPG{
			input:  "2,2 1,0",
			player: board.Black,
			action: move,
			result: true,
			state: board.Board{
				{board.White, board.Empty, board.White},
				{board.Black, board.White, board.Empty},
				{board.Empty, board.Black, board.Black},
			},
		},
		testPG{
			input:  "1,2 1,0",
			player: board.Black,
			action: move,
			result: false,
			state: board.Board{
				{board.White, board.Black, board.White},
				{board.Black, board.White, board.Empty},
				{board.Empty, board.Empty, board.Black},
			},
		},
		testPG{
			input:  "0,0 0,2",
			player: board.White,
			action: move,
			result: false,
			state: board.Board{
				{board.White, board.Black, board.White},
				{board.Black, board.White, board.Empty},
				{board.Empty, board.Empty, board.Black},
			},
		},
	}

	for i, v := range cases {
		r := playGame(game, v.player, v.action, v.input)
		if r != v.result {
			t.Errorf("Mismatch between expected \"%v\" and actual \"%v\" result for case %d", v.result, r, i)
		}
	}
}

type testEC struct {
	x   int
	y   int
	err *error
}

func (ec testEC) String() string {
	err := *ec.err
	if err == nil {
		return fmt.Sprintf("%d %d %v", ec.x, ec.y, nil)
	}
	return fmt.Sprintf("%d %d %v", ec.x, ec.y, err.Error())
}

func (ec testEC) isEqual(x, y int, err error) bool {
	if err != nil && ec.err == nil {
		return false
	}

	if err == nil && ec.err != nil {
		return false
	}

	if x != ec.x || y != ec.y {
		return false
	}
	return true
}

func TestExtractCoor(t *testing.T) {
	popError := fmt.Errorf("")
	errEC := testEC{
		x:   0,
		y:   0,
		err: &popError,
	}

	cases := map[string]testEC{
		"1,1": testEC{
			x:   1,
			y:   1,
			err: nil,
		},
		"0,0": testEC{
			x:   0,
			y:   0,
			err: nil,
		},
		"":    errEC,
		"4,1": errEC,
		"1,4": errEC,
	}

	for k, expected := range cases {
		x, y, err := extractCoor(k)
		r := testEC{
			x:   x,
			y:   y,
			err: &popError,
		}
		if !expected.isEqual(x, y, err) {
			t.Errorf("Expected {%+v} got {%+v}", expected, r)
		}
	}
}
