package main

import "fmt"

// Create board type
type board struct {
	data [9]string
}

func initBoard() board {
	return board{data: [9]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}}
}

// Displays the board
func (b board) displayBoard() {
	fmt.Println("-------------")
	fmt.Printf("| %v | %v | %v |\n", b.data[0], b.data[1], b.data[2])
	fmt.Println("-------------")
	fmt.Printf("| %v | %v | %v |\n", b.data[3], b.data[4], b.data[5])
	fmt.Println("-------------")
	fmt.Printf("| %v | %v | %v |\n", b.data[6], b.data[7], b.data[8])
	fmt.Println("-------------")
}

// Checks if there is a winner
func (b board) checkWin() bool {
	winArr := [8][3]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
		{0, 3, 6},
		{1, 4, 7},
		{2, 5, 8},
		{0, 4, 8},
		{2, 4, 6},
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
func (b board) isFull() bool {
	var count int = 0
	for i := 0; i < 9; i++ {
		if b.data[i] == "X" || b.data[i] == "O" {
			count++
		}
	}
	return count == 9
}

// Takes the position and checks if the position is valid
func (b board) isValid(pos int) bool {
	return b.data[pos-1] != "X" && b.data[pos-1] != "O"
}

// Returns a valid position input from the user
func (b board) getPos() int {
	for {
		fmt.Print("Pick a number from 1-9: ")
		var pos int
		fmt.Scan(&pos)
		if pos <= 9 && pos >= 1 && b.isValid(pos) {
			return pos
		}
	}
}

func (b *board) assignPos(pos int, marker string) {
	(*b).data[pos-1] = marker
}
