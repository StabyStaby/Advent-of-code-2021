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
	fmt.Println(puzzle2(strings.Split(string(file), "\r\n")))
}

func puzzle1(data []string) int {
	bytes := []string{"0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0"}
	for i := 0; i < len(bytes); i++ {
		bytes[i] = CommonByte(data, i, true)
	}
	var gamma int32
	var epsilon int32
	for i := 0; i < len(bytes); i++ {
		if bytes[i] == "1" {
			gamma = gamma | (1 << (len(bytes) - 1 - i))
		} else {
			epsilon = epsilon | (1 << (len(bytes) - 1 - i))
		}
	}
	return int(gamma) * int(epsilon)
}
func puzzle2(data []string) int {
	selectedBytes := data
	for i := 0; len(selectedBytes) > 1; i++ {
		commonByte := CommonByte(selectedBytes, i, true)
		selectedBytes = getBytes(selectedBytes, i, commonByte)
	}
	tmp := selectedBytes[0]
	var oxygen int64
	for i := 0; i < 12; i++ {
		if string(tmp[i]) == "1" {
			oxygen = oxygen | (1 << (len(tmp) - i))
		}
	}
	oxygen, _ = strconv.ParseInt(tmp, 2, 64)
	fmt.Println(tmp)
	selectedBytes = data
	for i := 0; len(selectedBytes) > 1; i++ {
		commonByte := CommonByte(selectedBytes, i, false)
		selectedBytes = getBytes(selectedBytes, i, commonByte)
	}
	tmp = selectedBytes[0]
	var co2 int64 = 0
	for i := 0; i < 12; i++ {
		if string(tmp[i]) == "1" {
			co2 = co2 | (1 << (len(tmp) - i))
		}
	}
	fmt.Println()
	co2, _ = strconv.ParseInt(tmp, 2, 64)
	fmt.Println(tmp)
	fmt.Println(oxygen, co2)
	return int(oxygen) * int(co2)
}

func CommonByte(data []string, position int, mostCommon bool) string {
	count := 0
	for _, v := range data {
		if string(v[position]) == "1" {
			count++
		}
	}
	/*2 main cases and 2 sub cases each
	count returns the number of 1's
	if count > half most common is 1
	else the most common is 0
	*/
	if mostCommon {
		if count >= len(data)/2 {
			return "1"
		} else {
			return "0"
		}
	} else {
		if count < len(data)/2 {
			return "1"
		} else {
			return "0"
		}
	}
}
func getBytes(data []string, position int, bite string) []string {
	bytes := []string{}
	for _, v := range data {
		if string(v[position]) == bite {
			bytes = append(bytes, string(v))
		}
	}
	return bytes
}
