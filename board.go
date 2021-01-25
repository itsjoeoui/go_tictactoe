package main

import "fmt"

// Create board type
type board struct {
	data [10]string
}

// Displays the board
func (b board) displayBoard() {
	fmt.Println("-------------")
	fmt.Printf("| %v | %v | %v |\n", b.data[1], b.data[2], b.data[3])
	fmt.Println("-------------")
	fmt.Printf("| %v | %v | %v |\n", b.data[4], b.data[5], b.data[6])
	fmt.Println("-------------")
	fmt.Printf("| %v | %v | %v |\n", b.data[7], b.data[8], b.data[9])
	fmt.Println("-------------")
}

// Checks if there is a winner
func (b board) checkWin() bool {
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
func (b board) isFull() bool {
	var count int = 0
	for i := 0; i < 10; i++ {
		if b.data[i] == "X" || b.data[i] == "O" {
			count++
		}
	}
	return count == 9
}

// Takes the position and checks if the position is valid
func (b board) isValid(pos int) bool {
	return b.data[pos] != "X" && b.data[pos] != "O"
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

func (b board) assignPos(pos int, marker string) board {
	b.data[pos] = marker
	return b
}
