package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/fatih/color"
	"github.com/inancgumus/screen"
)

type Player struct {
	number int
	board  [10][10]int
	radar  [10][10]int
	hits   int
	misses int
	sunk   int
}

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
	var player1 = Player{number: 1, hits: 0, misses: 0, sunk: 0}
	InitBoard(&player1.board)
	InitBoard(&player1.radar)

	var player2 = Player{number: 1, hits: 0, misses: 0, sunk: 0}
	InitBoard(&player2.board)
	InitBoard(&player2.radar)

	PrintPlayerTurn(&player1)
	fmt.Println("----PLAYER 1----")
	player1.board = PlayerPlacingShips(player1.board)
	screen.Clear()
	screen.MoveTopLeft()
	fmt.Println("----PLAYER 2----")
	player2.board = PlayerPlacingShips(player2.board)
	screen.Clear()
	screen.MoveTopLeft()
	var playing = true
	for playing {
		var collum, row int
		PrintPlayerTurn(&player1)
		fmt.Print("\n\nWhere do you want to place your ship >> ")
		row, collum = CollectUserAttackInput()
		AttackBoard(&player1, &player2, row, collum)
	}
}

func InitBoard(board *[10][10]int) {
	for i, collum := range board {
		for j := range collum {
			board[i][j] = 0
		}
	}
}

func PlayerPlacingShips(board [10][10]int) [10][10]int {
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
	return board
}

func AttackBoard(attacker *Player, victim *Player, row int, collum int) {
	if victim.board[collum][row] >= 3 {
		attacker.radar[collum][row] = 1
		victim.board[collum][row] = 1
		attacker.hits++
		fmt.Println("HIT!!!!!!")
	}
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

func CollectUserAttackInput() (int, int) {
	var error_color = color.New(color.FgRed)
	var success = false
	var success_row = false
	var success_collum = false
	var collum_rune rune
	var row int
	var collum int
	for !success {
		reader := bufio.NewReader(os.Stdin)
		fmt.Scanf("%c%d", &collum_rune, &row)
		collum_rune = unicode.ToUpper(collum_rune)
		collum = int(collum_rune) - 65
		reader.ReadString('\n')
		if row >= 1 && row <= 10 {
			success_row = true
		}
		if collum >= 0 && collum <= 9 {
			success_collum = true
		}
		if success_collum && success_row {
			success = true
		}
		if !success {
			error_color.Print("invalid input, try again>> ")
		}
	}
	return row, collum
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
	sea_color := color.New(color.FgCyan, color.BgBlue)
	ship_color := color.New(color.FgBlack, color.BgHiWhite)
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
		space_color := color.New()
		for _, value := range collum {
			space_color.Print(" ")
			if value == 0 {
				sea_color.Print("~")
				space_color = sea_color
			} else {
				ship_color.Print("O")
				space_color = ship_color
			}
		}
		fmt.Print("\n")
	}
}

func PrintPlayerTurn(player *Player) {
	stats := fmt.Sprintf("PLAYER: %d\nMisses: %d\nHits: %d\nSunk %d", player.number, player.misses, player.hits, player.sunk)
	borderStats := CreateBorder(stats, *color.New())
	borderStatsRows := strings.Split(borderStats, "\n")
	//print Radar + stats
	radar_color := color.New(color.FgHiGreen, color.BgBlack)
	// radar_miss := color.New(color.FgHiBlue)
	radar_hit := color.New(color.FgHiRed)
	fmt.Print("  ")
	for i := range player.board {
		fmt.Print(" ")
		letter := 65 + i
		fmt.Printf("%c", letter)
	}
	fmt.Print("\n")

	// print out rest of board
	for x, collum := range player.board {
		// ensure all lines are correctly aligned
		if x != 9 {
			fmt.Print(" ")
		}
		// print row number
		fmt.Print(x + 1)
		space_color := color.New(color.BgBlack)
		for _, value := range collum {
			space_color.Print(" ")
			if value == 0 {
				radar_color.Print("-")

			} else {
				radar_hit.Print("X")
			}
		}
		if x < len(borderStatsRows) {
			fmt.Print("\t\t", borderStatsRows[x])
		}
		fmt.Print("\n")
	}
	PrintBoard(player.board)
}

func CreateBorder(text string, text_color color.Color) string {
	var lines = strings.Split(text, "\n")
	var length int = 0
	for _, line := range lines {
		if len(line) > length {
			length = len(line)
		}
	}
	var result string = "┌"
	for i := 0; i < length; i++ {
		result += "─"
	}
	result += "┐\n"
	for _, line := range lines {
		result += "│"
		for i := 0; i < length; i++ {
			if i < len(line) {
				result += string(line[i])
			} else {
				result += " "
			}
		}
		result += "│\n"
	}
	result += "└"
	for i := 0; i < length; i++ {
		result += "─"
	}
	result += "┘\n"
	// text_color.Printf(result)
	return result
}
