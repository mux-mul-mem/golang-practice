package main

import (
	"fmt"
	"time"
	"math/rand"
)

const WIDTH int = 52
const HEIGHT int = 32
const INTERVAL int = 500
const INIT_RATE int = 10

func main() {
	generation := 0
	var board [HEIGHT][WIDTH]bool
	rand.Seed(time.Now().UnixNano())
	for h := 1; h < HEIGHT - 1; h++ {
		for w := 1; w < WIDTH - 1; w++ {
			board[h][w] = rand.Intn(INIT_RATE) == 0;
		}
	}
	for {
		generation += 1
		fmt.Println("\033[H\033[2J")
		for h := 0; h < HEIGHT; h++ {
			for w := 0; w < WIDTH; w++ {
				if board[h][w] {
					fmt.Print("#")
				} else {
					fmt.Print(" ")
				}
			}
			fmt.Print("\n")
		}
		var nextBoard [HEIGHT][WIDTH]bool
		for h := 1; h < HEIGHT - 1; h++ {
			for w := 1; w < WIDTH - 1; w++ {
				count := 0
				if board[h - 1][w - 1] {
					count += 1
				}
				if board[h - 1][w + 1] {
					count += 1
				}
				if board[h + 1][w - 1] {
					count += 1
				}
				if board[h + 1][w + 1] {
					count += 1
				}
				if board[h - 1][w] {
					count += 1
				}
				if board[h + 1][w] {
					count += 1
				}
				if board[h][w - 1] {
					count += 1
				}
				if board[h][w + 1] {
					count += 1
				}
				if board[h][w] && (count == 2 || count == 3) {
					nextBoard[h][w] = true
				} else if !board[h][w] && count == 3 {
					nextBoard[h][w] = true
				}
			}
		}
		board = nextBoard
		fmt.Printf("Generation %d.\n", generation)
		fmt.Println("Press Ctrl+C to Quit.")
		time.Sleep(time.Duration(INTERVAL) * time.Millisecond)
	}
}
