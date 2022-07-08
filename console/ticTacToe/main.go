package main

import (
	"os"
	"fmt"
	"time"
	"bufio"
	"strconv"
	"math/rand"
)

const INPUT_MIN_NUMBER int = 1
const INPUT_MAX_NUMBER int = 9

func main() {
	var scn = bufio.NewScanner(os.Stdin)
	var board [9]int
	fmt.Println("### Start ###")
	for {

		drawBoard(board)
		fmt.Println("Please input number (1 to 9).")
		var rawInput string

		if scn.Scan() {
			rawInput = scn.Text()
		}

		input, err := strconv.Atoi(rawInput)
		if err != nil {
			fmt.Println("### Error! ###")
			continue
		}

		if input < INPUT_MIN_NUMBER {
			fmt.Println("### Too small! ###")
			continue
		} else if input > INPUT_MAX_NUMBER {
			fmt.Println("### Too big! ###")
			continue
		} else if (board[input - 1]) > 0 {
			fmt.Println("### Already selected! ###")
			continue
		}

		board[input - 1] = 1
		if checkGameOver(board, 1) {
			drawBoard(board)
			fmt.Println("### You win! ###")
			break
		} else if countBoardSpace(board) < 1 {
			drawBoard(board)
			fmt.Println("### Draw! ###")
			break
		}

		rand.Seed(time.Now().UnixNano())
		randNumber := rand.Intn(countBoardSpace(board))

		for i := 0; i < INPUT_MAX_NUMBER; i++ {
			if (board[i]) > 0 {
				randNumber += 1
			}
			if randNumber == i {
				board[i] = 2
				break
			}
		}
		if checkGameOver(board, 2) {
			drawBoard(board)
			fmt.Println("### You lose! ###")
			break
		}
		fmt.Println("### Opponent selected %d. ###", randNumber + 1)
	}
}

func drawBoard(values [9]int) {
	var marks [9]string

	for i, v := range values {
		if v == 1 {
			marks[i] = "O"
		} else if v == 2 {
			marks[i] = "X"
		} else {
			marks[i] = strconv.Itoa(i + 1)
		}
	}
	fmt.Printf(
`=============
| %s | %s | %s |
=============
| %s | %s | %s |
=============
| %s | %s | %s |
=============
`,
		marks[0], marks[1], marks[2],
		marks[3], marks[4], marks[5],
		marks[6], marks[7], marks[8],
	)
}

func countBoardSpace(values [9]int) int {
	count := 0
	for _, v := range values {
		if v == 0 {
			count ++
		}
	}
	return count
}

func checkGameOver(values [9]int, value int) bool {
	return (values[0] == value && values[1] == value && values[2] == value) ||
		(values[3] == value && values[4] == value && values[5] == value) ||
		(values[6] == value && values[7] == value && values[8] == value) ||
		(values[0] == value && values[3] == value && values[6] == value) ||
		(values[1] == value && values[4] == value && values[7] == value) ||
		(values[2] == value && values[5] == value && values[8] == value) ||
		(values[0] == value && values[4] == value && values[8] == value) ||
		(values[2] == value && values[4] == value && values[6] == value) 
}
