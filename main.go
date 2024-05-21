package main

import (
	"fmt"

	"github.com/fatih/color"
)

// Code used for respresenting boards in code
// 0: empty space
// 1: ship
// 2: missed
// 3: hit ship

func main() {
	var board [10][10]int
	for i, collum := range board {
		for j := range collum {
			board[i][j] = 0
		}
	}
	PrintBoard(board)
}

func PlayerPlacingShips(board [10][10]int) {
	fmt.Println("Formatting your input, please do coordinates with no space, then a space and then the direction the ship should face(N, S, E, W)")
	fmt.Println("For example: E4 E")
	fmt.Println("board:")
	fmt.Print("Where to place destroyer (2) >> ")
	var collum rune
	var row int
	var direction rune

	fmt.Scanf("%c%d %c", collum, row, direction)
	collum_num := int(collum) - 65
	board = PlaceSingleShip(board, collum_num, row, direction, 2)

}

func PlaceSingleShip(board [10][10]int, x int, y int, direction rune, length int) [10][10]int {
	board[y][x] = 1
	var dy int = 0
	var dx int = 0
	switch direction {
	case 'N':
		dy = 1
	case 'E':
		dx = 1
	case 'S':
		dy = -1
	case 'W':
		dx = -1
	}
	for i := range length {
		board[y+(dy*i)][x+(dx*i)] = 1
	}
}

func PrintBoard(board [10][10]int) {
	sea_color := color.New(color.FgCyan)
	miss_color := color.New(color.FgHiWhite)
	// print out letter row at top
	fmt.Print(" ")
	for i := range board {
		fmt.Print(" ")
		letter := 65 + i
		fmt.Printf("%c", letter)
	}
	fmt.Print("\n")

	// print out rest of board
	for i, collum := range board {
		// print row number
		fmt.Print(i)
		for _, value := range collum {
			fmt.Print(" ")
			if value == 0 {
				sea_color.Print("~")
			} else if value == 1 {
				miss_color.Print("x")
			}
		}
		fmt.Print("\n")
	}
}
