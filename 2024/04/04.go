package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func part1(arr [][]rune) int {
	lines := []string{}

	// get horizontal lines
	for _, row := range arr {
		lines = append(lines, string(row))
	}

	// get vertical lines
	for i := 0; i < len(arr[0]); i++ {
		col := []rune{}
		for j := 0; j < len(arr); j++ {
			col = append(col, arr[j][i])
		}
		lines = append(lines, string(col))
	}

	diag, countdig := make(map[int][]rune), make(map[int][]rune)

	for r := 0; r < len(arr); r++ {
		for c := 0; c < len(arr[0]); c++ {
			// move top left to bottom right
			kdig := r - c
			diag[kdig] = append(diag[kdig], arr[r][c])

			// move top right to bottom left
			kcdig := r + c
			countdig[kcdig] = append(countdig[kcdig], arr[r][c])
		}
	}

	for _, d := range diag {
		lines = append(lines, string(d))
	}
	for _, d := range countdig {
		lines = append(lines, string(d))
	}

	count := 0
	for _, line := range lines {
		count += strings.Count(line, "XMAS")
		count += strings.Count(line, "SAMX")
	}
	return count
}

func part2(arr [][]rune) int {
	count := 0
	rows, cols := len(arr), len(arr[0])

	for r := 1; r < rows-1; r++ {
		for c := 1; c < cols-1; c++ {
			if arr[r][c] == 'A' {
				diag1 := map[rune]bool{arr[r-1][c-1]: true, arr[r+1][c+1]: true}
				diag2 := map[rune]bool{arr[r-1][c+1]: true, arr[r+1][c-1]: true}

				if len(diag1) == 2 && diag1['M'] && diag1['S'] &&
					len(diag2) == 2 && diag2['M'] && diag2['S'] {
					count++
				}
			}
		}
	}
	return count
}
func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	arr := [][]rune{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		txt := scanner.Text()
		chars := []rune{}
		for _, v := range txt {
			chars = append(chars, v)
		}
		arr = append(arr, chars)
	}

	fmt.Println(part1(arr), part2(arr))
}
