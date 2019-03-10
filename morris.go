package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"super-three-men-morris/board"
)

const rules = `
* Each player has 3 pieces, one set white and the other black
* The board is a 3x3 grid with 9 spaces
* Players take turns placing their pieces on empty intersections, starting with the white one
* Once all pieces are placed, player need to move one of their pieces per turn
* A piece may move to any vacant point on the same row/column/diagonal, not just an adjacent one
* To win a player must align their pieces on a line
`

const winTrophy = `
             ___________
            '._==_==_=_.'
            .-\:      /-.
           | (|:.     |) |
            '-|:.     |-'
              \::.    /
               '::. .'
                 ) (
               _.' '._
=====================================
           Player %s wins!
=====================================
`

const (
	zero = 0
	add  = "add"
	move = "move"
)

func main() {
	fmt.Println("Welcome to the SUPER Three men morris! Have fun!")
	fmt.Println(rules)

	game := board.Init()
	player := board.White
	piecesLeft := 3
	action := add

	scanner := bufio.NewScanner(os.Stdin)
	for {
		if piecesLeft == zero {
			action = move
		}

		fmt.Println(game)
		fmt.Printf("Player %s %s: ", player, action)

		scanner.Scan()
		input := scanner.Text()
		retry := playGame(game, player, action, input)
		if retry {
			continue
		}

		if game.CheckVictory() {
			fmt.Printf(winTrophy, player)
			os.Exit(0)
		}

		if player == board.White {
			player = board.Black
		} else {
			player = board.White
			piecesLeft = piecesLeft - 1
		}
		fmt.Print("\n")
	}
}

func playGame(game *board.Board, player, action, input string) bool {
	op := strings.Split(input, " ")
	if action == add {
		if len(op) < 1 {
			fmt.Println("\nNo coordinates received")
			return true
		}

		x, y, err := extractCoor(op[0])
		if err != nil {
			fmt.Println(err.Error())
			return true
		}

		err = game.AddPiece(x, y, player)
		if err != nil {
			fmt.Printf("Failed to add piece: %s", err.Error())
			return true
		}
	} else {
		if len(op) < 2 {
			fmt.Println("\nNo enough coordinates received")
			return true
		}

		ox, oy, err := extractCoor(op[0])
		if err != nil {
			fmt.Println(err.Error())
			return true
		}

		tx, ty, err := extractCoor(op[1])
		if err != nil {
			fmt.Println(err.Error())
			return true
		}

		err = game.MovePiece(ox, oy, tx, ty, player)
		if err != nil {
			fmt.Printf("Failed to move piece: %s", err.Error())
			return true
		}
	}
	return false
}

func extractCoor(str string) (int, int, error) {
	move := strings.Split(str, ",")
	x, err := strconv.Atoi(move[0])
	if err != nil {
		return 0, 0, fmt.Errorf("\nX coordinate in \"%s\" is invalid", move[0])
	}
	y, err := strconv.Atoi(move[1])
	if err != nil {
		return 0, 0, fmt.Errorf("\nY coordinate in \"%s\" is invalid", move[1])
	}

	if x < 0 || x > 3 {
		return 0, 0, fmt.Errorf("\nX coordinate invalid: %d out of bounds", x)
	}

	if y < 0 || y > 3 {
		return 0, 0, fmt.Errorf("\nY coordinate invalid: %d out of bounds", y)
	}

	return x, y, nil
}
