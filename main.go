package main

import "fmt"

func getYOrNInputAndConvert(prompt string) bool {
	for {
		fmt.Print(prompt)
		var yn string
		fmt.Scanln(&yn)
		if yn == "y" {
			return true
		} else if yn == "n" {
			return false
		}
	}
}

func main() {
	fmt.Println("Welcome to TicTacToe!")

	var myBoard = initBoard()

	var player1 string
	fmt.Print("[Player1] - Please enter your name: ")
	fmt.Scanln(&player1)

	fmt.Printf("Hey %v, ", player1)
	againstComp := getYOrNInputAndConvert("wanna play against a computer? [y/n] ")

	if againstComp {
		// TBD
		// isSmart := getYOrNInputAndConvert("Against a smart computer? [y/n] ")

	} else {
		var player2 string
		fmt.Print("[Player2] - Please enter your name: ")
		fmt.Scanln(&player2)

		var player string
		var marker string

		for {
			if myBoard.isFull() {
				fmt.Print("\033[H\033[2J") // Clear console
				fmt.Println("You two tied!")
				myBoard.displayBoard()
				break
			}

			if player == player1 {
				player = player2
				marker = "O"
			} else {
				player = player1
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
}
