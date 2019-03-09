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
