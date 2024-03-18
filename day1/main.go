package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.ReadFile("puzzle fiile.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(puzzle1(strings.Split(string(file), "\n")))
	log.Println(puzzle2(strings.Split(string(file), "\n")))
}
func puzzle1(depths []string) int {

	previousDepth, _ := strconv.Atoi(depths[0])
	var count int = 0
	for _, depth := range depths {
		i, err := strconv.Atoi(depth)
		if err != nil {
			log.Fatal(err)
		}
		if i > previousDepth {
			count++
		}
		previousDepth = i
	}
	log.Println(count)
	return count
}
func puzzle2(depths []string) int {
	var previousSum int
	count := 0
	for i := 0; i < len(depths)-2; i++ {
		curentSum, _ := strconv.Atoi(depths[i])
		if i+1 < len(depths) {
			x, _ := strconv.Atoi(depths[i+1])
			curentSum = curentSum + x
		}
		if i+2 < len(depths) {
			x, _ := strconv.Atoi(depths[i+2])
			curentSum = curentSum + x
		}
		if curentSum > previousSum && previousSum != 0 {
			count++
		}
		previousSum = curentSum
	}
	return count
}
