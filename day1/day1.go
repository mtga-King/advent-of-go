package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var numbers = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func main() {
	wd := filepath.Join("day1", "input.txt")
	read, err := os.ReadFile(wd)
	if err != nil {
		panic(err)
	}

	calibrations := strings.Split(string(read), "\r\n")

	sum := 0
	for _, c := range calibrations {
		str := ""
		calibration := []rune(c)
		size := len(calibration)

		for f := 0; f < size; f++ {
			forward := string(calibration[f])
			if value, _ := strconv.Atoi(forward); value != 0 {
				str += forward
				str += reverse(calibration, size)
				break
			}

			if f+5 <= size {
				check := string(calibration[f : f+5])
				value, ok := numbers[check]
				if ok {
					str += value
					str += reverse(calibration, size)
					break
				}
			}

			if f+4 <= size {
				check := string(calibration[f : f+4])
				value, ok := numbers[check]
				if ok {
					str += value
					str += reverse(calibration, size)
					break
				}
			}

			if f+3 <= size {
				check := string(calibration[f : f+3])
				value, ok := numbers[check]
				if ok {
					str += value
					str += reverse(calibration, size)
					break
				}
			}
		}

		atoi, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}

		sum += atoi
	}
	fmt.Println(sum)
}

func reverse(calibration []rune, size int) string {
	for f := size - 1; f >= 0; f-- {
		reverse := string(calibration[f])
		if value, _ := strconv.Atoi(reverse); value != 0 {
			return reverse
		}

		if f-4 >= 0 {
			check := string(calibration[f-4 : f+1])
			if value, ok := numbers[check]; ok {
				return value
			}
		}

		if f-3 >= 0 {
			check := string(calibration[f-3 : f+1])
			if value, ok := numbers[check]; ok {
				return value
			}
		}

		if f-2 >= 0 {
			check := string(calibration[f-2 : f+1])
			if value, ok := numbers[check]; ok {
				return value
			}
		}
	}
	return ""
}
