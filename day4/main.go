package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
)

type BingoElement struct {
	number int
	marked bool
}
type BingoBoard struct{
	board [][]BingoElement
}

func main() {
	file, err := os.ReadFile("puzzle file.txt")
	if err != nil {
		log.Fatal(err)
	}
	
	puzzle1()
}
func puzzle1() {
	board := [5][]BingoElement{}
	element := BingoElement{9, false}
	board[0] = append(board[0], element)
	// board[0] = append(board[0], BingoElement{10,false})
	fmt.Println(board)

	matrix := [4][]int{{1}}

	matrix[0] = append(matrix[0], 2)
	matrix[1] = append(matrix[1], 2)
	matrix[2] = append(matrix[2], 2)
	matrix[3] = append(matrix[3], 2)
	fmt.Println(matrix)
}
var numbersRegex = regexp.MustCompile(`([0-9],)`)
func parseFile(file []byte) ([]int,BingoBoard,error){


}
