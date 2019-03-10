package board

import (
	"testing"
)

type testCV struct {
	x      int
	y      int
	player string
	win    bool
}

func TestCheckVictoryRow(t *testing.T) {
	game := Init()
	plays := []testCV{
		testCV{
			x:      0,
			y:      0,
			player: White,
			win:    false,
		},
		testCV{
			x:      0,
			y:      1,
			player: White,
			win:    false,
		},
		testCV{
			x:      0,
			y:      2,
			player: White,
			win:    true,
		},
	}

	for i, v := range plays {
		game.AddPiece(v.x, v.y, v.player)
		r := game.CheckVictory()
		if r != v.win {
			t.Errorf("Mismatch between expected \"%v\" and actual \"%v\" victory state in step %d", v.win, r, i)
			t.Errorf("Board state: \n %s", game)
		}
	}
}

func TestCheckNonVictoryRow(t *testing.T) {
	game := Init()
	plays := []testCV{
		testCV{
			x:      0,
			y:      0,
			player: White,
			win:    false,
		},
		testCV{
			x:      0,
			y:      1,
			player: Black,
			win:    false,
		},
		testCV{
			x:      0,
			y:      2,
			player: White,
			win:    false,
		},
	}

	for i, v := range plays {
		game.AddPiece(v.x, v.y, v.player)
		r := game.CheckVictory()
		if r != v.win {
			t.Errorf("Mismatch between expected \"%v\" and actual \"%v\" victory state in step %d", v.win, r, i)
			t.Errorf("Board state: \n %s", game)
		}
	}
}

func TestCheckVictoryCol(t *testing.T) {
	game := Init()
	plays := []testCV{
		testCV{
			x:      0,
			y:      0,
			player: White,
			win:    false,
		},
		testCV{
			x:      1,
			y:      0,
			player: White,
			win:    false,
		},
		testCV{
			x:      2,
			y:      0,
			player: White,
			win:    true,
		},
	}

	for i, v := range plays {
		game.AddPiece(v.x, v.y, v.player)
		r := game.CheckVictory()
		if r != v.win {
			t.Errorf("Mismatch between expected \"%v\" and actual \"%v\" victory state in step %d", v.win, r, i)
			t.Errorf("Rotated board state: \n %s", rotatedBoard)
		}
	}
}

func TestCheckNonVictoryCol(t *testing.T) {
	game := Init()
	plays := []testCV{
		testCV{
			x:      0,
			y:      0,
			player: White,
			win:    false,
		},
		testCV{
			x:      1,
			y:      0,
			player: Black,
			win:    false,
		},
		testCV{
			x:      2,
			y:      0,
			player: White,
			win:    false,
		},
	}

	for i, v := range plays {
		game.AddPiece(v.x, v.y, v.player)
		r := game.CheckVictory()
		if r != v.win {
			t.Errorf("Mismatch between expected \"%v\" and actual \"%v\" victory state in step %d", v.win, r, i)
			t.Errorf("Rotated board state: \n %s", rotatedBoard)
		}
	}
}

func TestCheckVictoryCross(t *testing.T) {
	game := Init()
	plays := []testCV{
		testCV{
			x:      1,
			y:      1,
			player: White,
			win:    false,
		},
		testCV{
			x:      0,
			y:      0,
			player: White,
			win:    false,
		},
		testCV{
			x:      2,
			y:      2,
			player: White,
			win:    true,
		},
	}

	for i, v := range plays {
		game.AddPiece(v.x, v.y, v.player)
		r := game.CheckVictory()
		if r != v.win {
			t.Errorf("Mismatch between expected \"%v\" and actual \"%v\" victory state in step %d", v.win, r, i)
			t.Errorf("Board state: \n %s", game)
		}
	}

	game = Init()
	plays = []testCV{
		testCV{
			x:      1,
			y:      1,
			player: White,
			win:    false,
		},
		testCV{
			x:      0,
			y:      2,
			player: White,
			win:    false,
		},
		testCV{
			x:      2,
			y:      0,
			player: White,
			win:    true,
		},
	}

	for i, v := range plays {
		game.AddPiece(v.x, v.y, v.player)
		r := game.CheckVictory()
		if r != v.win {
			t.Errorf("Mismatch between expected \"%v\" and actual \"%v\" victory state in step %d", v.win, r, i)
			t.Errorf("Board state: \n %s", game)
		}
	}
}

func TestCheckNonVictoryCross(t *testing.T) {
	game := Init()
	plays := []testCV{
		testCV{
			x:      1,
			y:      1,
			player: Black,
			win:    false,
		},
		testCV{
			x:      0,
			y:      0,
			player: White,
			win:    false,
		},
		testCV{
			x:      2,
			y:      2,
			player: White,
			win:    false,
		},
	}

	for i, v := range plays {
		game.AddPiece(v.x, v.y, v.player)
		r := game.CheckVictory()
		if r != v.win {
			t.Errorf("Mismatch between expected \"%v\" and actual \"%v\" victory state in step %d", v.win, r, i)
			t.Errorf("Board state: \n %s", game)
		}
	}

	game = Init()
	plays = []testCV{
		testCV{
			x:      1,
			y:      1,
			player: Black,
			win:    false,
		},
		testCV{
			x:      0,
			y:      2,
			player: White,
			win:    false,
		},
		testCV{
			x:      2,
			y:      0,
			player: White,
			win:    false,
		},
	}

	for i, v := range plays {
		game.AddPiece(v.x, v.y, v.player)
		r := game.CheckVictory()
		if r != v.win {
			t.Errorf("Mismatch between expected \"%v\" and actual \"%v\" victory state in step %d", v.win, r, i)
			t.Errorf("Board state: \n %s", game)
		}
	}
}

type testIM struct {
	ox        int
	oy        int
	tx        int
	ty        int
	isInvalid bool
}

func TestInvalidMove(t *testing.T) {
	cases := []testIM{
		/* Case 0
		  0 1 2
		0 ●
		1   V
		2
		*/
		testIM{
			ox:        0,
			oy:        0,
			tx:        1,
			ty:        1,
			isInvalid: false,
		},
		/* Case 1
		  0 1 2
		0
		1   ●
		2  	V
		*/
		testIM{
			ox:        1,
			oy:        1,
			tx:        1,
			ty:        2,
			isInvalid: false,
		},
		/* Case 2
		  0 1 2
		0
		1   ●
		2  	  V
		*/
		testIM{
			ox:        1,
			oy:        1,
			tx:        2,
			ty:        2,
			isInvalid: false,
		},
		/* Case 3
		  0 1 2
		0
		1   ● V
		2
		*/
		testIM{
			ox:        1,
			oy:        1,
			tx:        2,
			ty:        1,
			isInvalid: false,
		},
		/* Case 4
		  0 1 2
		0 ● - V
		1
		2
		*/
		testIM{
			ox:        0,
			oy:        0,
			tx:        2,
			ty:        0,
			isInvalid: false,
		},
		/* Case 5
		  0 1 2
		0 ●
		1 |
		2 V
		*/
		testIM{
			ox:        0,
			oy:        0,
			tx:        0,
			ty:        2,
			isInvalid: false,
		},
		/* Case 6
		  0 1 2
		0 ●
		1   \
		2     V
		*/
		testIM{
			ox:        0,
			oy:        0,
			tx:        2,
			ty:        2,
			isInvalid: false,
		},
		/* Case 7
		  0 1 2
		0 ● -
		1   |
		2   V
		*/
		testIM{
			ox:        0,
			oy:        0,
			tx:        1,
			ty:        2,
			isInvalid: true,
		},
		/* Case 8
		  0 1 2
		0 ● - |
		1     V
		2
		*/
		testIM{
			ox:        0,
			oy:        0,
			tx:        2,
			ty:        1,
			isInvalid: true,
		},
		/* Case 9
		  0 1 2
		0   V
		1 ● |
		2
		*/
		testIM{
			ox:        0,
			oy:        1,
			tx:        1,
			ty:        0,
			isInvalid: false,
		},
		/* Case 10
		  0 1 2
		0
		1 ● - V
		2
		*/
		testIM{
			ox:        0,
			oy:        1,
			tx:        2,
			ty:        1,
			isInvalid: false,
		},
		/* Case 11
		  0 1 2
		0
		1 ●
		2 v
		*/
		testIM{
			ox:        0,
			oy:        1,
			tx:        0,
			ty:        2,
			isInvalid: false,
		},
		/* Case 12
		  0 1 2
		0
		1 ● - |
		2     x
		*/
		testIM{
			ox:        0,
			oy:        1,
			tx:        2,
			ty:        2,
			isInvalid: true,
		},
		/* Case 13
		  0 1 2
		0     x
		1 ● - |
		2
		*/
		testIM{
			ox:        0,
			oy:        1,
			tx:        2,
			ty:        0,
			isInvalid: true,
		},
	}

	for i, c := range cases {
		r := invalidMove(c.ox, c.oy, c.tx, c.ty)
		if r != c.isInvalid {
			t.Errorf("Expected \"%v\" move but got \"%v\" move instead in case %d", c.isInvalid, r, i)
		}
	}
}
