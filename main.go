package main

import (
	"fmt"
)

func main() {
	var board [10][10]int
	for i, _ := range board {
		for j, _ := range board {
			board[i][j] = 0
		}
	}
}

func PlayerPlacingShips(board [10][10]int){
	fmt.Println("Formatting your input, please do coordinates with no space, then a space and then the direction the ship should face(N, S, E, W)")
	fmt.Println("For example: E4 E")


}
func PrintBoard(board [10][10]int){
	if 
}