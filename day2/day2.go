package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	red int = iota + 12
	green
	blue
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

	allGames := strings.Split(string(read), "\r\n")
	sum := 0

	//NewGame:
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
				/*
					if nums[color] < atoi {
						continue NewGame
					}
				*/
			}
		}

		power := 1 //can't multiply by 0
		for _, value := range setValues {
			power *= value
		}
		sum += power
		/*
			atoi, err := strconv.Atoi(strings.Split(gameRound, " ")[1])
			if err != nil {
				panic(err)
			}
			sum += atoi
		*/
	}
	fmt.Println(sum)
}
