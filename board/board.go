package board

import "fmt"

type Board [3][3]string

const (
	Empty = " "
	White = "●"
	Black = "○"
)

func Init() *Board {
	var eb Board
	for i := range eb {
		eb[i] = [3]string{Empty, Empty, Empty}
	}
	return &eb
}

//AddPiece...
func (b *Board) AddPiece(x, y int, p string) error {
	if b[x][y] != Empty {
		return fmt.Errorf("(%d,%d) is not empty.", x, y)
	}
	b[x][y] = p
	return nil
}

//   1 2 3
// 1 ●
// 2   ○
// 3     ●
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
