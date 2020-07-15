package main

import "fmt"

func main() {
	board := fill(' ')

	//black pieces
	board[0][0] = 'r'
	board[0][7] = 'r'
	board[0][1] = 'n'
	board[0][6] = 'n'
	board[0][2] = 'b'
	board[0][5] = 'b'
	board[0][3] = 'q'
	board[0][4] = 'k'

	//White pieces
	board[7][0] = 'R'
	board[7][7] = 'R'
	board[7][1] = 'N'
	board[7][6] = 'N'
	board[7][2] = 'B'
	board[7][5] = 'B'
	board[7][3] = 'Q'
	board[7][4] = 'K'

	for column := range board[1] {
		board[1][column] = 'p'
		board[6][column] = 'P'
	}

	displayBoard(board)
}

func displayBoard(board [8][8]rune) {
	for row := range board {
		for column := range board[row] {
			fmt.Printf("%c ", board[row][column])
		}
		fmt.Println()
	}
}

func fill(symbol rune) (board [8][8]rune) {
	for row := range board {
		for column := range board[row] {
			board[row][column] = symbol
		}
	}
	return
}
