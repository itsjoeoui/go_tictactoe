package main

import (
	"fmt"
)

// Board class
type Board struct {
	data [10]string
}

// Displays the board
func (b Board) displayBoard() {
	fmt.Println("-------------")
	fmt.Printf("| %v | %v | %v |\n", b.data[1], b.data[2], b.data[3])
	fmt.Println("-------------")
	fmt.Printf("| %v | %v | %v |\n", b.data[4], b.data[5], b.data[6])
	fmt.Println("-------------")
	fmt.Printf("| %v | %v | %v |\n", b.data[7], b.data[8], b.data[9])
	fmt.Println("-------------")
}

// Checks if there is a winner
func (b Board) checkWin() bool {
	winArr := [8][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
		{1, 4, 7},
		{2, 5, 8},
		{3, 6, 9},
		{1, 5, 9},
		{3, 5, 7},
	}
	for i := 0; i < 8; i++ {
		if b.data[winArr[i][0]] == b.data[winArr[i][1]] &&
			b.data[winArr[i][0]] == b.data[winArr[i][2]] {
			return true
		}
	}
	return false
}

// Checks if the board is full
func (b Board) isFull() bool {
	var count int = 0
	for i := 0; i < 10; i++ {
		if b.data[i] == "X" || b.data[i] == "O" {
			count++
		}
	}
	return count == 9
}

// Takes the position and checks if the position is valid
func (b Board) isValid(pos int) bool {
	return b.data[pos] != "X" && b.data[pos] != "O"
}

// Returns a valid position input from the user
func (b Board) getPos() int {
	for true {
		fmt.Print("Pick a number from 1-9: ")
		var pos int
		fmt.Scan(&pos)
		if pos <= 9 && pos >= 1 && b.isValid(pos) {
			return pos
		}
	}
	return -1
}

func (b Board) assignPos(pos int, marker string) Board {
	b.data[pos] = marker
	return b
}

func main() {
	fmt.Println("Welcome to TicTacToe!")

	var board = Board{data: [10]string{" ", "1", "2", "3", "4", "5", "6", "7", "8", "9"}}

	var player1 string
	var player2 string
	fmt.Print("[Player1] - Please enter your name: ")
	fmt.Scanln(&player1)
	fmt.Print("[Player2] - Please enter your name: ")
	fmt.Scanln(&player2)

	// gameOn := true
	turn := 0

	for true {
		if board.isFull() {
			// Clear console
			fmt.Print("\033[H\033[2J")

			fmt.Println("You two tied!")

			board.displayBoard()
			break
		}

		var player string
		var marker string

		if turn%2 == 0 {
			player = player1
			marker = "X"
		} else {
			player = player2
			marker = "O"
		}

		// Clear console
		fmt.Print("\033[H\033[2J")

		fmt.Printf("It's %v's turn! [%v]\n", player, marker)
		board.displayBoard()
		board = board.assignPos(board.getPos(), marker)
		turn++
		if board.checkWin() {
			// Clear console
			fmt.Print("\033[H\033[2J")

			if turn%2 != 0 {
				fmt.Printf("Congratulations %v, you won!\n", player1)
			} else {
				fmt.Printf("Congratulations %v, you won!\n", player2)
			}

			board.displayBoard()
			break
		}
	}
}
