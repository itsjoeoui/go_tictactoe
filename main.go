package main

import "fmt"

func main() {
	fmt.Println("Welcome to TicTacToe!")

	var myBoard = initBoard()

	var player1 humanPlayer
	fmt.Print("[Player1] - Please enter your name: ")
	player1.assignName()

	var player2 humanPlayer
	fmt.Print("[Player2] - Please enter your name: ")
	player2.assignName()

	// gameOn := true
	turn := 0

	for {
		if myBoard.isFull() {
			// Clear console
			fmt.Print("\033[H\033[2J")
			fmt.Println("You two tied!")
			myBoard.displayBoard()
			break
		}

		var player string
		var marker string

		if turn%2 == 0 {
			player = player1.name
			marker = "X"
		} else {
			player = player2.name
			marker = "O"
		}

		// Clear console
		fmt.Print("\033[H\033[2J")

		fmt.Printf("It's %v's turn! [%v]\n", player, marker)

		myBoard.displayBoard()
		myBoard.assignPos(myBoard.getPos(), marker)

		turn++
		if myBoard.checkWin() {
			// Clear console
			fmt.Print("\033[H\033[2J")
			fmt.Printf("Congratulations %v, you won!\n", player)
			myBoard.displayBoard()
			break
		}
	}
}
