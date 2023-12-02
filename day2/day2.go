package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var nums = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func main() {
	wd := filepath.Join("day2", "input.txt")
	read, err := os.ReadFile(wd)
	if err != nil {
		panic(err)
	}

	part1(read)
	part2(read)
}

func part1(read []uint8) {
	allGames := strings.Split(string(read), "\r\n")
	sum := 0

NewGame:
	for _, g := range allGames {
		games := strings.Split(g, ": ")
		gameRound, sets := games[0], strings.Split(games[1], "; ")

		for _, s := range sets {
			set := strings.Split(s, ", ")

			for _, v := range set {
				values := strings.Split(v, " ")
				value, color := values[0], values[1]

				atoi, err := strconv.Atoi(value)
				if err != nil {
					panic(err)
				}

				if nums[color] < atoi {
					continue NewGame
				}

			}
		}

		atoi, err := strconv.Atoi(strings.Split(gameRound, " ")[1])
		if err != nil {
			panic(err)
		}
		sum += atoi
	}
	fmt.Println(sum)
}

func part2(read []uint8) {
	allGames := strings.Split(string(read), "\r\n")
	sum := 0

	for _, g := range allGames {
		games := strings.Split(g, ": ")
		_, sets := games[0], strings.Split(games[1], "; ") //gameRound

		setValues := make(map[string]int)
		for _, s := range sets {
			set := strings.Split(s, ", ")

			for _, v := range set {
				values := strings.Split(v, " ")
				value, color := values[0], values[1]

				atoi, err := strconv.Atoi(value)
				if err != nil {
					panic(err)
				}

				v, _ := setValues[color]
				if v < atoi {
					setValues[color] = atoi
				}
			}
		}

		power := 1 //can't multiply by 0
		for _, value := range setValues {
			power *= value
		}
		sum += power
	}
	fmt.Println(sum)
}
