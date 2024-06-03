package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"

	"github.com/fatih/color"
	"github.com/inancgumus/screen"
)

// Code used for respresenting boards in code
// 0: empty space
// 1: hit
// 2: miss
// Battleships:
// 3: destroyer
// 4: crusier
// 5: submarine
// 6: battleship
// 7: carrier

func main() {
	var player1_board [10][10]int
	for i, collum := range player1_board {
		for j := range collum {
			player1_board[i][j] = 0
		}
	}

	var player2_board [10][10]int
	for i, collum := range player2_board {
		for j := range collum {
			player2_board[i][j] = 0
		}
	}
	fmt.Println("----PLAYER 1----")
	PlayerPlacingShips(player1_board)
	screen.Clear()
	screen.MoveTopLeft()
	fmt.Println("----PLAYER 2----")
	PlayerPlacingShips(player2_board)
	screen.Clear()
	screen.MoveTopLeft()
}

func PlayerPlacingShips(board [10][10]int) {
	fmt.Println("Formatting your input, please do coordinates with no space, then a space and then the direction the ship should face(N, S, E, W)")
	fmt.Println("For example: E4 E")
	fmt.Println("board:")
	PrintBoard(board)

	fmt.Print("Where do you want to place your destroyer (2)>> ")
	row, collum, direction := CollectUserShipInput()
	board = PlaceSingleShip(board, collum, row, direction, 2, 3)
	PrintBoard(board)

	fmt.Print("Where do you want to place your cruiser (3)>> ")
	row, collum, direction = CollectUserShipInput()
	fmt.Println("Row: ", row, "collum: ", collum, "direction: ", direction)
	board = PlaceSingleShip(board, collum, row, direction, 3, 4)
	PrintBoard(board)

	fmt.Print("Where do you want to place your submarine (3)>> ")
	row, collum, direction = CollectUserShipInput()
	fmt.Println("Row: ", row, "collum: ", collum, "direction: ", direction)
	board = PlaceSingleShip(board, collum, row, direction, 3, 5)
	PrintBoard(board)

	fmt.Print("Where do you want to place your battleship (4)>> ")
	row, collum, direction = CollectUserShipInput()
	fmt.Println("Row: ", row, "collum: ", collum, "direction: ", direction)
	board = PlaceSingleShip(board, collum, row, direction, 4, 6)
	PrintBoard(board)

	fmt.Print("Where do you want to place your carrier (5)>> ")
	row, collum, direction = CollectUserShipInput()
	fmt.Println("Row: ", row, "collum: ", collum, "direction: ", direction)
	board = PlaceSingleShip(board, collum, row, direction, 5, 7)
	PrintBoard(board)
}

func PlaceSingleShip(board [10][10]int, x int, y int, direction rune, length int, ship_type int) [10][10]int {
	fmt.Println("Row: ", x, "collum: ", y, "direction: ", direction)
	var error_color = color.New(color.FgRed)
	var dy int = 0
	var dx int = 0
	success := false
	y--
	for !success {
		switch direction {
		case 'S':
			dy = 1
		case 'E':
			dx = 1
		case 'N':
			dy = -1
		case 'W':
			dx = -1
		}
		success = true
		for i := range length {
			if y+(dy*i) >= 10 || x+(dx*i) >= 10 || y+(dy*i) < 0 || x+(dx*i) < 0 {
				//checks if out of bounds
				success = false
				error_color.Print("your ship goes out of bounds, please place again>> ")
				break
			} else if board[(y + (dy * i))][(x+(dx*i))] != 0 {
				//ensures no colisions with other ships
				success = false
				error_color.Print("Your ship collides with another, please place again>> ")
				break
			}
		}
		if !success {
			x, y, direction = CollectUserShipInput()
		}

	}

	for i := range length {
		board[y+(dy*i)][x+(dx*i)] = ship_type
	}
	return board
}

func CollectUserShipInput() (int, int, rune) {
	var error_color = color.New(color.FgRed)
	var success = false
	var success_row = false
	var success_collum = false
	var success_dir = false
	var collum_rune rune
	var row int
	var direction rune
	var collum int
	var allowed_dirs = [...]rune{'N', 'S', 'E', 'W'}
	for !success {
		reader := bufio.NewReader(os.Stdin)
		fmt.Scanf("%c%d %c", &collum_rune, &row, &direction)
		collum_rune = unicode.ToUpper(collum_rune)
		collum = int(collum_rune) - 65
		reader.ReadString('\n')
		if row >= 1 && row <= 10 {
			success_row = true
		}
		if collum >= 0 && collum <= 9 {
			success_collum = true
		}
		for _, value := range allowed_dirs {
			if value == direction {
				success_dir = true
				break
			}
		}
		if success_collum && success_row && success_dir {
			success = true
		}
		if !success {
			error_color.Print("invalid input, try again>> ")
		}
	}
	return row, collum, direction
}

func PrintBoard(board [10][10]int) {
	sea_color := color.New(color.FgCyan)
	miss_color := color.New(color.FgHiWhite)
	// print out letter row at top
	fmt.Print("  ")
	for i := range board {
		fmt.Print(" ")
		letter := 65 + i
		fmt.Printf("%c", letter)
	}
	fmt.Print("\n")

	// print out rest of board
	for x, collum := range board {
		// ensure all lines are correctly aligned
		if x != 9 {
			fmt.Print(" ")
		}
		// print row number
		fmt.Print(x + 1)
		for _, value := range collum {
			fmt.Print(" ")
			if value == 0 {
				sea_color.Print("~")
			} else {
				miss_color.Print(value)
			}
		}
		fmt.Print("\n")
	}
}
