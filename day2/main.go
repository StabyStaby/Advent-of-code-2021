package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("puzzle file.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(puzzle1(strings.Split(string(file), "\n")))
	log.Println(puzzle2(strings.Split(string(file), "\n")))
	// add a test case if file does not exist

}
func puzzle1(instructions []string) int {
	depth := 0
	horizontal := 0

	for _, i := range instructions {
		instructuion, value := getIntinstructuion(strings.Split(i, " "))
		parseInstruction(instructuion, value, &depth, &horizontal)
	}
	return depth * horizontal
}
func puzzle2(instructions []string) int {
	depth := 0
	horizontal := 0
	aim := 0
	for _, i := range instructions {
		instructuion, value := getIntinstructuion(strings.Split(i, " "))
		parseInstruction2(instructuion, value, &depth, &horizontal, &aim)
	}
	return depth * horizontal

}
func parseInstruction2(instruction string, value int, depth *int, horizontal *int, aim *int) {
	if instruction == "forward" {
		*horizontal = *horizontal + value
		*depth = *depth + (*aim * value)
	}
	if instruction == "down" {
		*aim = *aim + value
	}
	if instruction == "up" {
		*aim = *aim - value
	}
}
func parseInstruction(instruction string, value int, depth *int, horizontal *int) {
	if instruction == "forward" {
		*horizontal = *horizontal + value
	}
	if instruction == "down" {
		*depth = *depth + value
	}
	if instruction == "up" {
		*depth = *depth - value
	}
}
func getIntinstructuion(instruction []string) (string, int) {
	value, err := strconv.Atoi(instruction[1])
	if err != nil {
		log.Fatal(err)
	}
	return instruction[0], value
}
