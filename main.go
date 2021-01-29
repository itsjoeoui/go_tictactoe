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

	var player string
	var marker string

	for {
		if myBoard.isFull() {
			fmt.Print("\033[H\033[2J") // Clear console
			fmt.Println("You two tied!")
			myBoard.displayBoard()
			break
		}

		if player == player1.name {
			player = player2.name
			marker = "O"
		} else {
			player = player1.name
			marker = "X"
		}

		fmt.Print("\033[H\033[2J") // Clear console
		fmt.Printf("It's %v's turn! [%v]\n", player, marker)
		myBoard.displayBoard()
		myBoard.assignPos(myBoard.getPos(), marker)

		if myBoard.checkWin() {
			fmt.Print("\033[H\033[2J") // Clear console
			fmt.Printf("Congratulations %v, you won!\n", player)
			myBoard.displayBoard()
			break
		}
	}
}
