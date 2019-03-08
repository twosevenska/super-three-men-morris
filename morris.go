package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	empty = 0
	white = 1
	black = 2
)

func main() {
	fmt.Println("Welcome to the cousin of Tic-Tac-Toe - Kata! Have fun!")
	board := [3][3]int{}

	player := white
	scanner := bufio.NewScanner(os.Stdin)
	printBoard(board)
	for scanner.Scan() {
		input := scanner.Text()
		move := strings.Split(input, ",")
		x, _ := strconv.Atoi(move[0])
		y, _ := strconv.Atoi(move[1])
		board = addPiece(board, x, y, player)

		if player == white {
			player = black
		} else {
			player = white
		}
		printBoard(board)
	}

}

func addPiece(board [3][3]int, x, y, player int) [3][3]int {
	board[x][y] = player
	return board
}

//   1 2 3
// 1 ●
// 2   ○
// 3     ●
func printBoard(b [3][3]int) {
	fmt.Println("   0 1 2")
	fmt.Printf(" 0 %s %s %s\n", fp(b[0][0]), fp(b[0][1]), fp(b[0][2]))
	fmt.Printf(" 1 %s %s %s\n", fp(b[1][0]), fp(b[1][1]), fp(b[1][2]))
	fmt.Printf(" 2 %s %s %s\n", fp(b[2][0]), fp(b[2][1]), fp(b[2][2]))
}

func fp(i int) string {
	if i == white {
		return "●"
	}

	if i == black {
		return "○"
	}

	return " "
}

// valPos TODO: explain all val
func valPos(b [3][3]int, x, y int) bool {
	if b[x][y] == empty {
		return true
	}
	return false
}
