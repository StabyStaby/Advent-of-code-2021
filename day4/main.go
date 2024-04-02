package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"
)

type BingoElement struct {
	number int
	marked bool
}
type BingoBoard struct {
	board [5][]BingoElement
}

func main() {
	now := time.Now()
	file, err := os.Open("puzzle file.txt")
	if err != nil {
		log.Fatal(err)
	}
	filleScanner := bufio.NewScanner(file)
	filleScanner.Split(bufio.ScanLines)

	numbers, bingo, err := parseFile(*filleScanner)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(bingo[3].board)
	fmt.Println(numbers[1])

	checkForWinner(bingo[3])

	file.Close()
	// puzzle1()

	fmt.Println(time.Now().UnixNano() - now.UnixNano())
}
func puzzle1() {

}

var numbersRegex = regexp.MustCompile(`[0-9]*|\S,?`)

func parseFile(scanner bufio.Scanner) ([]int, []BingoBoard, error) {
	numbers := []int{}
	newBoard := false
	tmpBoard := BingoBoard{}
	returnBoard := []BingoBoard{}
	for scanner.Scan() {
		if scanner.Text() == "" {
			newBoard = true
			if len(tmpBoard.board) != 0 {
				returnBoard = append(returnBoard, tmpBoard)
			}
		} else {
			numbersString := numbersRegex.FindAllString(scanner.Text(), -1)
			if len(numbersString) > 5 {
				for index, number := range numbersString {
					i, err := strconv.Atoi(number)
					if err != nil {
						log.Println(index)
						return nil, nil, err
					}
					numbers = append(numbers, i)
				}
			} else {
				if newBoard {
					tmpBoard = BingoBoard{}
					newBoard = false
				}
				for i, number := range numbersString {
					tmpNumber, err := strconv.Atoi(number)
					if err != nil {
						log.Println(i)
						return nil, nil, err
					}
					tmpBoard.board[i] = append(tmpBoard.board[i], BingoElement{tmpNumber, false})
				}
			}
		}
	}
	return numbers, returnBoard, nil
}

func checkForWinner(board BingoBoard) bool {
	// need to check column and row for
	var consecutiveColumn int
	var consecutiveRow int
	for i, element := range board.board {
		fmt.Printf("element: %v\n", element)
		if element[i].marked {
			consecutiveColumn++
		}
		for _, subelement := range board.board[i] {
			fmt.Printf("subelement: %v\n", subelement)
			if subelement.marked {
				consecutiveRow++
			}
		}
	}
	fmt.Printf("consecutiveColumn: %v\n", consecutiveColumn)
	fmt.Printf("consecutiveRow: %v\n", consecutiveRow)
	return false
}
