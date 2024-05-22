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
	// fmt.Print("testing")
	board = PlaceSingleShip(board, 5, 5, 'E', 3)
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
	fmt.Print(x, y, direction)
	var dy int = 0
	var dx int = 0
	success := false
	for !success {
		x--
		y--
		board[y][x] = 1
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
		fmt.Printf("\ndy: %d\nDx: %d", dy, dx)
		fmt.Printf("\ny: %d\nx: %d", y, x)
		success = true
		for i := range length {
			fmt.Println("\n\n\n", i)
			fmt.Printf("\ndy: %d\nDx: %d", dy, dx)
			fmt.Printf("\ny: %d\nx: %d", y, x)
			fmt.Printf("\nchecking location: X: %d, and Y: %d", (x + (dx * i)), (y + (dy * i)))
			//checks if out of bounds
			if y+(dy*i) >= 10 || x+(dx*i) >= 10 || y+(dy*i) < 0 || x+(dx*i) < 0 {
				success = false
				color.Red("your ship goes out of bounds, please place again>> ")
			} else if board[(y + (dy * i))][(x+(dx*i))] == 1 {
				//ensures no colisions with other ships
				success = false
				color.Red("Your ship collides with another, please place again>> ")
			}
		}
		if !success {
			x, y, direction = CollectUserShipInput()
		}

	}

	for i := range length {
		board[y+(dy*i)][x+(dx*i)] = 1
	}
	return board
}

func CollectUserShipInput() (int, int, rune) {
	var collum_rune rune
	var row int
	var direction rune
	fmt.Scanf("%c%d %c", collum_rune, row, direction)
	collum := int(collum_rune) - 65
	return row, collum, direction
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
		//ensure all lines are correctly aligned
		if i != 9 {
			fmt.Print(" ")
		}
		// print row number
		fmt.Print(i + 1)
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
