package main

import (
	"fmt"
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
	fmt.Println(puzzle1(strings.Split(string(file), "\r\n")))
}

func puzzle1(data []string) int {
	bytes := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	for _, v := range data {
		for i := 0; i < len(v); i++ {
			// number, err := strconv.Atoi(string(v[i]))
			number, err := strconv.ParseInt(string(v[i]), 2, 32)
			if err != nil {
				log.Fatal(err)
			}
			if number == 1 {
				bytes[i]++
			}
		}
	}
	var gamma int32
	var epsilon int32
	for i := 0; i < len(bytes); i++ {
		if bytes[i] > len(data)/2 {
			gamma = gamma | (1 << (len(bytes) - 1 - i))
		} else {
			epsilon = epsilon | (1 << (len(bytes) - 1 - i))
		}
	}
	return int(gamma) * int(epsilon)
}
func puzzle2(data []string) int {
	selectedBytes := []string{}

	return 0
}
