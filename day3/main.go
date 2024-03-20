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
	fmt.Println(puzzle1(strings.Split(string(file), "\n")))
}

func puzzle1(data []string) int {
	bytes := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	for _, v := range data {
		for i := 0; i < len(v); i++ {
			number, err := strconv.Atoi(string(v[i]))
			if err != nil {
				log.Fatal(err)
			}
			if number == 1 {
				bytes[i]++
			}
		}
	}
	var gamma byte
	var epsilon byte
	for i := 0; i < len(bytes); i++ {
		// for i := len(bytes) - 1; i >= 0; i-- {
		if bytes[i] > len(data)/2 {
			gamma = gamma | (1 << i)
			epsilon = epsilon | (0 << i)
		} else {
			gamma = gamma | (0 << i)
			epsilon = epsilon | (1 << i)
		}
	}
	fmt.Println(bytes)
	fmt.Println(gamma)
	fmt.Println(epsilon)
	fmt.Print("epislon: ")
	for i := 11; i >= 0; i-- {
		bit := epsilon & (1 << i)
		if bit != 0 {
			fmt.Print("1")
		} else {
			fmt.Print("0")
		}
	}
	fmt.Println()
	fmt.Print("gamma: ")
	for i := 11; i >= 0; i-- {
		bit := gamma & (1 << i)
		if bit != 0 {
			fmt.Print("1")
		} else {
			fmt.Print("0")
		}
	}
	fmt.Println()
	return int(gamma) * int(epsilon)
}
