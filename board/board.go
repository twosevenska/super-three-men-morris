package board

import "fmt"

type Board [3][3]string

const (
	Empty = " "
	White = "●"
	Black = "○"
)

//rotatedBoard should be the same as the other one,
//difference being rows are now columns and vice-versa
var rotatedBoard Board

func Init() *Board {
	var b Board
	for i := range b {
		b[i] = [3]string{Empty, Empty, Empty}
		rotatedBoard[i] = [3]string{Empty, Empty, Empty}
	}
	return &b
}

//AddPiece...
func (b *Board) AddPiece(x, y int, p string) error {
	if b[x][y] != Empty {
		return fmt.Errorf("(%d,%d) is not empty.", x, y)
	}
	b[x][y] = p
	rotatedBoard[y][x] = p
	return nil
}

//   0 1 2
// 0 ●
// 1   ○
// 2     ●
func (b *Board) String() string {
	t := `
    0 1 2
  0 %s %s %s
  1 %s %s %s
  2 %s %s %s
  `
	return fmt.Sprintf(t,
		b[0][0], b[0][1], b[0][2],
		b[1][0], b[1][1], b[1][2],
		b[2][0], b[2][1], b[2][2],
	)
}

//CheckVictory inspects the board for the 3 types of victories
func (b *Board) CheckVictory() bool {
	for _, r := range b {
		if checkLineVictory(r) {
			return true
		}
	}

	for _, r := range rotatedBoard {
		if checkLineVictory(r) {
			return true
		}
	}

	if b[1][1] != Empty && b.checkCrossVictory() {
		return true
	}

	return false
}

// checkLineVictory is used to check both row and column victory but
// for both the board(row == row) and the rotatedBoard(row == column)
func checkLineVictory(r [3]string) bool {
	w := [3]string{White, White, White}
	b := [3]string{Black, Black, Black}
	if r == w || r == b {
		return true
	}
	return false
}

func (b *Board) checkCrossVictory() bool {
	corePiece := b[1][1]

	if b[0][0] == corePiece && b[2][2] == corePiece {
		return true
	}

	if b[0][2] == corePiece && b[2][0] == corePiece {
		return true
	}

	return false
}
