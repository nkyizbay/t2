package main

import (
	"bufio"
	"fmt"
	math "math/rand"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

func drawBoard(board []string) {
	fmt.Println(board[7] + "|" + board[8] + "|" + board[9])
	fmt.Println("-+-+-")
	fmt.Println(board[4] + "|" + board[5] + "|" + board[6])
	fmt.Println("-+-+-")
	fmt.Println(board[1] + "|" + board[2] + "|" + board[3])
}

func inputPlayerLetter() [2]string {
	letter := ""
	for !(letter == "X" || letter == "O") {
		fmt.Println("What you choose: X or O?")
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter X or O: ")
		// letter, _ = reader.ReadString('\n')
		psletter, _ := reader.ReadString('\n')
		letter = strings.Trim(psletter, "\n\r")
		letter = strings.ToTitle(letter)
	}

	if letter == "X" {
		return [2]string{"X", "O"}
	} else {
		return [2]string{"O", "X"}
	}
}

func whoGoesFirst() string {
	if true { // math.Intn(2) == 0
		return "Computer"
	} else {
		return "Human"
	}
}

func makeMove(board []string, letter string, move int) { // move int
	board[move] = letter
}

func isWinner(bo []string, le string) bool {
	return ((bo[7] == le && bo[8] == le && bo[9] == le) || // top
		(bo[4] == le && bo[5] == le && bo[6] == le) || // center
		(bo[1] == le && bo[2] == le && bo[3] == le) || // bottom
		(bo[7] == le && bo[4] == le && bo[1] == le) || // left side up
		(bo[8] == le && bo[5] == le && bo[2] == le) || // center up
		(bo[9] == le && bo[6] == le && bo[3] == le) || // right side up
		(bo[7] == le && bo[5] == le && bo[3] == le) || // diagonal
		(bo[9] == le && bo[5] == le && bo[1] == le)) // diagonal
}

func getBoardCopy(board []string) []string {
	var boardCopy []string
	for i := range board {
		boardCopy = append(boardCopy, board[i])
	}
	return boardCopy
}

func isSpaceFree(board []string, move int) bool {
	return board[move] == " "
}

func getPlayerMove(board []string) int {
	move := " "
	s := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	m, _ := strconv.Atoi(move)
	for !(slices.Contains(s, move)) || !(isSpaceFree(board, m)) {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Your next move? (1-9): ")
		psmove, _ := reader.ReadString('\n')
		move = strings.Trim(psmove, "\n\r")
		// move1, _ := strconv.Atoi(move)
		// return strconv.Atoi(move)
		// return move1
	}
	moveN, _ := strconv.Atoi(move)
	return moveN
}

func chooseRandomMoveFromList(board []string, movesList []int) int {
	var possibleMoves []int

	for _, x := range movesList {
		if isSpaceFree(board, x) {
			possibleMoves = append(possibleMoves, x)
		}
	}

	if len(possibleMoves) != 0 {
		id := math.Intn(len(possibleMoves))
		return possibleMoves[id]
	} else {
		return 15
	}
}

func getComputerMove(board []string, computerLetter string) int {
	var playerLetter string
	var boardCopy []string
	var move int
	if computerLetter == "X" {
		playerLetter = "0"
	} else {
		playerLetter = "X"
	}

	for i := 1; i < 10; i++ {
		boardCopy = getBoardCopy(board)
		if isSpaceFree(boardCopy, i) {
			makeMove(boardCopy, computerLetter, i)
			if isWinner(boardCopy, computerLetter) {
				return i
			}
		}
	}

	for i := 0; i < 10; i++ {
		boardCopy = getBoardCopy(board)
		if isSpaceFree(boardCopy, i) {
			makeMove(boardCopy, playerLetter, i)
			if isWinner(boardCopy, playerLetter) {
				return i
			}
		}
	}

	move = chooseRandomMoveFromList(board, []int{1, 3, 7, 9})

	if move != 15 {
		return move
	}

	if isSpaceFree(board, 5) {
		return 5
	}

	return chooseRandomMoveFromList(board, []int{2, 4, 6, 8})
}

func isBoardFull(board []string) bool {
	for i := 1; i < 10; i++ {
		if isSpaceFree(board, i) {
			return false
		}
	}
	return true
}

func main() {
	for {
		theBoard := []string{" ", " ", " ", " ", " ", " ", " ", " ", " ", " "}
		// playerLetter, computerLetter := inputPlayerLetter()
		x := inputPlayerLetter()
		playerLetter, computerLetter := x[0], x[1]
		turn := whoGoesFirst()
		gameIsPlaying := true

		for gameIsPlaying {
			if turn == "Human" {
				drawBoard(theBoard)
				move := getPlayerMove(theBoard)
				makeMove(theBoard, playerLetter, move)

				if isWinner(theBoard, playerLetter) {
					drawBoard(theBoard)
					fmt.Println("You won! Congratulations!")
					gameIsPlaying = false
				} else {
					if isBoardFull(theBoard) {
						drawBoard(theBoard)
						fmt.Println("Nobody won.")
						break
					} else {
						turn = "Computer"
					}
				}
			} else {
				move := getComputerMove(theBoard, computerLetter)
				makeMove(theBoard, computerLetter, move)

				if isWinner(theBoard, computerLetter) {
					drawBoard(theBoard)
					fmt.Println("Computer won.")
					gameIsPlaying = false
				} else {
					if isBoardFull(theBoard) {
						drawBoard(theBoard)
						fmt.Println("Nobody won.")
						break
					} else {
						turn = "Human"
					}

				}

			}
		}
	}
}
